package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/walkergriggs/hellosb/restapi/operations/service_instances"
	"github.com/walkergriggs/hellosb/state"
)

type updateServiceInstance struct {
	store *state.StateStore
}

func NewUpdateServiceInstanceHandler(store *state.StateStore) service_instances.ServiceInstanceUpdateHandler {
	return &updateServiceInstance{
		store: store,
	}
}

func (impl *updateServiceInstance) Handle(params service_instances.ServiceInstanceUpdateParams, principal interface{}) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		if params.Body == nil {
			rw.WriteHeader(400)
			return
		}

		err := impl.store.UpdateServiceInstance(params.InstanceID, params.Body)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		rw.WriteHeader(200)
	})
}
