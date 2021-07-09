package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"k8s.io/utils/pointer"

	"github.com/walkergriggs/hellosb/models"
	"github.com/walkergriggs/hellosb/restapi/operations/service_instances"
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

	instance := &models.ServiceInstanceResource{
		MaintenanceInfo: request.MaintenanceInfo,
		PlanID:          pointer.StringDeref(request.PlanID, "planid"),
		ServiceID:       pointer.StringDeref(request.ServiceID, "serviceid"),
		Parameters:      request.Parameters,
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		err := impl.store.InsertServiceInstance(params.InstanceID, instance)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		rw.WriteHeader(200)
	})
}
