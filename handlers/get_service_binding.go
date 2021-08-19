package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/walkergriggs/hellosb/api/operations/service_bindings"
	"github.com/walkergriggs/hellosb/state"
)

type getServiceBindingImpl struct {
	store *state.StateStore
}

func NewGetServiceBindingHandler(store *state.StateStore) service_bindings.ServiceBindingGetHandler {
	return &getServiceBindingImpl{
		store: store,
	}
}

func (impl *getServiceBindingImpl) Handle(params service_bindings.ServiceBindingGetParams, principal interface{}) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		binding, err := impl.store.GetServiceBinding(params.BindingID)
		if err != nil {
			http.Error(rw, "Failed to get service binding: "+err.Error(), http.StatusInternalServerError)
			return
		}

		b, err := binding.MarshalBinary()
		if err != nil {
			http.Error(rw, "Failed to marshal response: "+err.Error(), http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(b)
	})
}
