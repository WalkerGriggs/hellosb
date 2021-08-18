package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/walkergriggs/hellosb/api/operations/service_instances"
	"github.com/walkergriggs/hellosb/state"
)

type deprovisionServiceInstance struct {
	store *state.StateStore
}

func NewDeprovisionServiceInstanceHandler(store *state.StateStore) service_instances.ServiceInstanceDeprovisionHandler {
	return &deprovisionServiceInstance{
		store: store,
	}
}

func (impl *deprovisionServiceInstance) Handle(params service_instances.ServiceInstanceDeprovisionParams, principal interface{}) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		err := impl.store.DeleteServiceInstance(params.InstanceID)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		rw.WriteHeader(200)
	})
}
