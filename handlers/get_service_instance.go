package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/walkergriggs/hellosb/restapi/operations/service_instances"
)

type getServiceInstanceImpl struct{}

func NewGetServiceInstanceHandler() service_instances.ServiceInstanceGetHandler {
	return &getServiceInstanceImpl{}
}

func (impl *getServiceInstanceImpl) Handle(params service_instances.ServiceInstanceGetParams, principal interface{}) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		file, err := ioutil.ReadFile("./mocks/service_instance.json")
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		rw.WriteHeader(200)
		rw.Write([]byte(file))
	})
}
