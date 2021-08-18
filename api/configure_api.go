package api

import (
	"net/http"

	"github.com/walkergriggs/hellosb/handlers"
	"github.com/walkergriggs/hellosb/state"
)

type Config struct {
	CatalogPath string
}

func (api *API) Configure(config *Config) http.Handler {
	if api.config == nil {
		api.config = config
	}

	if api.store == nil {
		store, err := state.NewStateStore()
		if err != nil {
			panic(err)
		}
		api.store = store
	}

	// CatalogGet Handler
	api.CatalogGetHandler = handlers.NewGetCatalogHandler(api.config.CatalogPath)

	// ServiceInstance Provision Handler
	api.ServiceInstanceProvisionHandler = handlers.NewProvisionServiceInstanceHandler(api.store)

	// ServiceInstance Get Handler
	api.ServiceInstanceGetHandler = handlers.NewGetServiceInstanceHandler(api.store)

	// ServiceInstance Update Handler
	api.ServiceInstanceUpdateHandler = handlers.NewUpdateServiceInstanceHandler(api.store)

	// ServiceInstance Deprovision Handler
	api.ServiceInstanceDeprovisionHandler = handlers.NewDeprovisionServiceInstanceHandler(api.store)

	// ServiceInstance LastOperation Handler
	api.ServiceInstanceLastOperationGetHandler = handlers.NewGetServiceInstanceLastOperationHandler(api.store)

	// ServiceBinding Provision Handler
	api.ServiceBindingBindingHandler = handlers.NewProvisionServiceBindingHandler(api.store)

	// ServiceBinding Get Handler
	api.ServiceBindingGetHandler = handlers.NewGetServiceBindingHandler(api.store)

	// ServiceBinding Deprovision Handler
	api.ServiceBindingUnbindingHandler = handlers.NewDeprovisionServiceBindingHandler(api.store)

	// ServiceBinding LastOperation Handler
	api.ServiceBindingLastOperationGetHandler = handlers.NewGetServiceBindingLastOperationHandler(api.store)

	return api.Serve(nil)
}
