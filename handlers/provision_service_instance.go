package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"k8s.io/utils/pointer"

	"github.com/walkergriggs/hellosb/api/operations/service_instances"
	"github.com/walkergriggs/hellosb/models"
	"github.com/walkergriggs/hellosb/state"
)

type provisionServiceInstance struct {
	store *state.StateStore
}

func NewProvisionServiceInstanceHandler(store *state.StateStore) service_instances.ServiceInstanceProvisionHandler {
	return &provisionServiceInstance{
		store: store,
	}
}

func (impl *provisionServiceInstance) Handle(params service_instances.ServiceInstanceProvisionParams, principal interface{}) middleware.Responder {
	request := params.Body

	response := &models.ServiceInstanceProvisionResponse{}

	acceptsIncomplete := pointer.BoolDeref(params.AcceptsIncomplete, false)

	instance := &models.ServiceInstanceResource{
		MaintenanceInfo: request.MaintenanceInfo,
		PlanID:          pointer.StringDeref(request.PlanID, "planid"),
		ServiceID:       pointer.StringDeref(request.ServiceID, "serviceid"),
		Parameters:      request.Parameters,
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		b, err := response.MarshalBinary()

	HAS_ERR:
		if err != nil {
			http.Error(rw, "Failed to provision service instance: "+err.Error(), http.StatusInternalServerError)
			return
		}

		existing, err := impl.store.GetServiceInstance(params.InstanceID)
		if err != nil {
			goto HAS_ERR
		}

		if existing != nil {
			if reflect.DeepEqual(existing.Parameters, instance.Parameters) {
				if isInstanceProvisioning(existing) && acceptsIncomplete {
					// MUST be returned if the Service Instance provisioning is in
					// progress. The operation string MUST match that returned for the
					// original request. This triggers the Platform to poll the Last
					// Operation for Service Instances endpoint for operation status. Note
					// that a re-sent PUT request MUST return a 202 Accepted, not a 200
					// OK, if the Service Instance is not yet fully provisioned.
					rw.WriteHeader(http.StatusAccepted)
				} else {
					// SHOULD be returned if the Service Instance already exists, is
					// fully provisioned, and the requested parameters are identical to
					// the existing Service Instance. This response is only valid in
					// synchronous operations.
					rw.WriteHeader(http.StatusOK)
				}
				rw.Write(b)
			} else {

				// MUST be returned if a Service Instance with the same id already
				// exists or is being provisioned but with different attributes.
				rw.WriteHeader(http.StatusConflict)
			}
			return
		} else if acceptsIncomplete {
			// MUST be returned if the Service Instance provisioning is in
			// progress. The operation string MUST match that returned for the
			// original request. This triggers the Platform to poll the Last
			// Operation for Service Instances endpoint for operation status. Note
			// that a re-sent PUT request MUST return a 202 Accepted, not a 200
			// OK, if the Service Instance is not yet fully provisioned.
			rw.WriteHeader(http.StatusAccepted)
			return
		} else {
			attrs, err := ParseSystemAttributes(instance.Parameters)
			if err != nil {
				goto HAS_ERR
			}

			instance.Metadata = &models.ServiceInstanceMetadata{
				Attributes: attrs,
			}

			err = impl.store.InsertServiceInstance(params.InstanceID, instance)
			if err != nil {
				goto HAS_ERR
			}

			// MUST be returned if the Service Instance was provisioned as a result of
			// this request.
			rw.WriteHeader(http.StatusCreated)
			rw.Write(b)
		}
	})
}

// isInstanceProvisioning is used to determine if a service is _currently_
// provisioning. A false response does not indicate that the instance has failed
// to provision or has already provisioned; it's entirely time-based.
func isInstanceProvisioning(instance *models.ServiceInstanceResource) bool {
	attrs := instance.Metadata.Attributes.(map[string]interface{})

	createdAt := attrs["created_at"].(time.Time)
	duration, ok := attrs["provision_duration"].(time.Duration)
	if !ok {
		return false // a service without a provision duration is assumed to be synchronous
	}

	return createdAt.Add(duration).After(time.Now())
}

// ParseSystemAttributes is used to parrse any hellosb-specific parameters.
// These attributes are used to codify behavior that is shared among all service
// instances, regardless of plans, maintenances etc. Examples include a
// "created_at" timestamp or "provision_duration".
func ParseSystemAttributes(parameters interface{}) (interface{}, error) {
	attrs := map[string]interface{}{
		"created_at": time.Now(),
	}

	if parameters != nil {
		params, ok := parameters.(map[string]interface{}) // making assumptions about params shape
		if !ok {
			return nil, fmt.Errorf("Failed to parse parameters. Is it a map?")
		}

		if val, ok := params["opts_provision_duration"]; ok {
			str, ok := val.(string)
			if !ok {
				return nil, fmt.Errorf("Failed to parse opts_provision_duration. Is it a string?")
			}

			duration, err := time.ParseDuration(str)
			if err != nil {
				return nil, err
			}

			attrs["provision_duration"] = duration
		}
	}

	return attrs, nil
}
