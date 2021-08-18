package api

import (
	"net/http"

	"github.com/walkergriggs/hellosb/handlers"
	"github.com/walkergriggs/hellosb/state"
)

func (api *API) Configure() http.Handler {
	store, err := state.NewStateStore()
	if err != nil {
		panic(err)
	}

	// CatalogGet Handler
	api.CatalogGetHandler = handlers.NewGetCatalogHandler("/Users/w.griggs/Documents/catalog.json")

	// ServiceInstance Provision Handler
	api.ServiceInstanceProvisionHandler = handlers.NewProvisionServiceInstanceHandler(store)

	// ServiceInstance Get Handler
	api.ServiceInstanceGetHandler = handlers.NewGetServiceInstanceHandler(store)

	// ServiceInstance Update Handler
	api.ServiceInstanceUpdateHandler = handlers.NewUpdateServiceInstanceHandler(store)

	// ServiceInstance Deprovision Handler
	api.ServiceInstanceDeprovisionHandler = handlers.NewDeprovisionServiceInstanceHandler(store)

	// ServiceInstance LastOperation Handler
	api.ServiceInstanceLastOperationGetHandler = handlers.NewGetServiceInstanceLastOperationHandler(store)

	// ServiceBinding Provision Handler
	api.ServiceBindingBindingHandler = handlers.NewProvisionServiceBindingHandler(store)

	// ServiceBinding Get Handler
	api.ServiceBindingGetHandler = handlers.NewGetServiceBindingHandler(store)

	// ServiceBinding Deprovision Handler
	api.ServiceBindingUnbindingHandler = handlers.NewDeprovisionServiceBindingHandler(store)

	// ServiceBinding LastOperation Handler
	api.ServiceBindingLastOperationGetHandler = handlers.NewGetServiceBindingLastOperationHandler(store)

	return api.Serve(nil)
}
