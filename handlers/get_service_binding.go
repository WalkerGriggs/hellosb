package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/walkergriggs/hellosb/restapi/operations/service_bindings"
)

type getServiceBindingImpl struct{}

func NewGetServiceBindingHandler() service_bindings.ServiceBindingGetHandler {
	return &getServiceBindingImpl{}
}

func (impl *getServiceBindingImpl) Handle(params service_bindings.ServiceBindingGetParams, principal interface{}) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		file, err := ioutil.ReadFile("./mocks/service_binding.json")
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		rw.WriteHeader(200)
		rw.Write([]byte(file))
	})
}
