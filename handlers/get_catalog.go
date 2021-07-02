package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/walkergriggs/hellosb/restapi/operations/catalog"
)

type getCatalogImpl struct{}

func NewGetCatalogHandler() catalog.CatalogGetHandler {
	return &getCatalogImpl{}
}

func (impl *getCatalogImpl) Handle(params catalog.CatalogGetParams, principal interface{}) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		file, err := ioutil.ReadFile("./mocks/catalog.json")
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		rw.WriteHeader(200)
		rw.Write([]byte(file))
	})
}