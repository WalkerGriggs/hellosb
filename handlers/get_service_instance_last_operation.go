package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"k8s.io/utils/pointer"

	"github.com/walkergriggs/hellosb/api/operations/service_instances"
	"github.com/walkergriggs/hellosb/models"
	"github.com/walkergriggs/hellosb/state"
)

type getServiceInstanceLastOperation struct {
	store *state.StateStore
}

func NewGetServiceInstanceLastOperationHandler(store *state.StateStore) service_instances.ServiceInstanceLastOperationGetHandler {
	return &getServiceInstanceLastOperation{
		store: store,
	}
}

func (impl *getServiceInstanceLastOperation) Handle(params service_instances.ServiceInstanceLastOperationGetParams, principal interface{}) middleware.Responder {
	operation := &models.LastOperationResource{
		State:            pointer.String("succeeded"),
		Description:      "Good to go",
		InstanceUsable:   true,
		UpdateRepeatable: false,
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		instance, err := impl.store.GetServiceInstance(params.InstanceID)
		if err != nil {
			http.Error(rw, "Failed to get service instance: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if instance == nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		b, err := operation.MarshalBinary()
		if err != nil {
			http.Error(rw, "Failed to marshal response: "+err.Error(), http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(b)
	})
}
