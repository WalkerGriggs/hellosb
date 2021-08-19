package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/walkergriggs/hellosb/api/operations/service_instances"
	"github.com/walkergriggs/hellosb/state"
)

type getServiceInstance struct {
	store *state.StateStore
}

func NewGetServiceInstanceHandler(store *state.StateStore) service_instances.ServiceInstanceGetHandler {
	return &getServiceInstance{
		store: store,
	}
}

func (impl *getServiceInstance) Handle(params service_instances.ServiceInstanceGetParams, principal interface{}) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		instance, err := impl.store.GetServiceInstance(params.InstanceID)
		if err != nil {
			http.Error(rw, "Failed to get service instance: "+err.Error(), http.StatusInternalServerError)
			return
		}

		b, err := instance.MarshalBinary()
		if err != nil {
			http.Error(rw, "Failed to marshal response: "+err.Error(), http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(b)
	})
}
