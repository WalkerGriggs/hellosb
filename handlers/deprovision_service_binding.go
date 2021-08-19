package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/walkergriggs/hellosb/api/operations/service_bindings"
	"github.com/walkergriggs/hellosb/state"
)

type deprovisionServiceBinding struct {
	store *state.StateStore
}

func NewDeprovisionServiceBindingHandler(store *state.StateStore) service_bindings.ServiceBindingUnbindingHandler {
	return &deprovisionServiceBinding{
		store: store,
	}
}

func (impl *deprovisionServiceBinding) Handle(params service_bindings.ServiceBindingUnbindingParams, principal interface{}) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		err := impl.store.DeleteServiceBinding(params.BindingID)
		if err != nil {
			http.Error(rw, "Failed to deprovision service binding: "+err.Error(), http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
	})
}
