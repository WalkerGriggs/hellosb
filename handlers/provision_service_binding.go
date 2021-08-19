package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/walkergriggs/hellosb/api/operations/service_bindings"
	"github.com/walkergriggs/hellosb/models"
	"github.com/walkergriggs/hellosb/state"
)

type provisionServiceBinding struct {
	store *state.StateStore
}

func NewProvisionServiceBindingHandler(store *state.StateStore) service_bindings.ServiceBindingBindingHandler {
	return &provisionServiceBinding{
		store: store,
	}
}

func (impl *provisionServiceBinding) Handle(params service_bindings.ServiceBindingBindingParams, principal interface{}) middleware.Responder {
	request := params.Body

	binding := &models.ServiceBindingResource{
		Parameters: request.Parameters,
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		err := impl.store.InsertServiceBinding(params.BindingID, binding)
		if err != nil {
			http.Error(rw, "Failed to bind service binding: "+err.Error(), http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
	})
}
