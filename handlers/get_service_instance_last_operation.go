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
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		if instance == nil {
			rw.WriteHeader(404)
			return
		}

		b, err := operation.MarshalBinary()
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		rw.WriteHeader(200)
		rw.Write(b)
	})
}
