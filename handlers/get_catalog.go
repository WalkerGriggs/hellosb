package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/walkergriggs/hellosb/api/operations/catalog"
	"github.com/walkergriggs/hellosb/models"
)

type getCatalogImpl struct {
	catalog *models.Catalog
}

func NewGetCatalogHandler(catalog *models.Catalog) catalog.CatalogGetHandler {
	return &getCatalogImpl{catalog}
}

func (impl *getCatalogImpl) Handle(params catalog.CatalogGetParams, principal interface{}) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		b, err := impl.catalog.MarshalBinary()
		if err != nil {
			http.Error(rw, "Failed to marshal response: "+err.Error(), http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(b)
	})
}
