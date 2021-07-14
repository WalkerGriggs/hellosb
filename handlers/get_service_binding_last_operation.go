package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"k8s.io/utils/pointer"

	"github.com/walkergriggs/hellosb/models"
	"github.com/walkergriggs/hellosb/restapi/operations/service_bindings"
	"github.com/walkergriggs/hellosb/state"
)

type getServiceBindingLastOperation struct {
	store *state.StateStore
}

func NewGetServiceBindingLastOperationHandler(store *state.StateStore) service_bindings.ServiceBindingLastOperationGetHandler {
	return &getServiceBindingLastOperation{
		store: store,
	}
}

func (impl *getServiceBindingLastOperation) Handle(params service_bindings.ServiceBindingLastOperationGetParams, principal interface{}) middleware.Responder {
	operation := &models.LastOperationResource{
		State:            pointer.String("succeeded"),
		Description:      "Good to go",
		InstanceUsable:   true,
		UpdateRepeatable: false,
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		binding, err := impl.store.GetServiceBinding(params.BindingID)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		if binding == nil {
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
