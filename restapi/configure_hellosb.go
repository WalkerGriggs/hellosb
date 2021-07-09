// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/walkergriggs/hellosb/handlers"
	"github.com/walkergriggs/hellosb/restapi/operations"
	"github.com/walkergriggs/hellosb/restapi/operations/service_bindings"
	"github.com/walkergriggs/hellosb/restapi/operations/service_instances"
	"github.com/walkergriggs/hellosb/state"
)

//go:generate swagger generate server --target ./hellosb --name Hellosb --spec ~/Documents/servicebroker/swagger.yaml --principal interface{}

func configureFlags(api *operations.HellosbAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HellosbAPI) http.Handler {
	store, err := state.NewStateStore()
	if err != nil {
		panic(err)
	}

	// Expected interface func(string, ...interface{})
	// api.Logger = log.Printf

	// api.UseRedoc()

	api.UseSwaggerUI()

	api.ServeError = errors.ServeError

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the Authorization header is set with the Basic scheme
	api.BasicAuthAuth = func(user string, pass string) (interface{}, error) {
		return "TODO", nil
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	api.CatalogCatalogGetHandler = handlers.NewGetCatalogHandler()

	api.ServiceInstancesServiceInstanceProvisionHandler = handlers.NewProvisionServiceInstanceHandler(store)

	api.ServiceInstancesServiceInstanceGetHandler = handlers.NewGetServiceInstanceHandler(store)

	api.ServiceBindingsServiceBindingGetHandler = handlers.NewGetServiceBindingHandler()

	if api.ServiceBindingsServiceBindingBindingHandler == nil {
		api.ServiceBindingsServiceBindingBindingHandler = service_bindings.ServiceBindingBindingHandlerFunc(func(params service_bindings.ServiceBindingBindingParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation service_bindings.ServiceBindingBinding has not yet been implemented")
		})
	}

	if api.ServiceBindingsServiceBindingLastOperationGetHandler == nil {
		api.ServiceBindingsServiceBindingLastOperationGetHandler = service_bindings.ServiceBindingLastOperationGetHandlerFunc(func(params service_bindings.ServiceBindingLastOperationGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation service_bindings.ServiceBindingLastOperationGet has not yet been implemented")
		})
	}

	if api.ServiceBindingsServiceBindingUnbindingHandler == nil {
		api.ServiceBindingsServiceBindingUnbindingHandler = service_bindings.ServiceBindingUnbindingHandlerFunc(func(params service_bindings.ServiceBindingUnbindingParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation service_bindings.ServiceBindingUnbinding has not yet been implemented")
		})
	}

	if api.ServiceInstancesServiceInstanceDeprovisionHandler == nil {
		api.ServiceInstancesServiceInstanceDeprovisionHandler = service_instances.ServiceInstanceDeprovisionHandlerFunc(func(params service_instances.ServiceInstanceDeprovisionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation service_instances.ServiceInstanceDeprovision has not yet been implemented")
		})
	}

	if api.ServiceInstancesServiceInstanceLastOperationGetHandler == nil {
		api.ServiceInstancesServiceInstanceLastOperationGetHandler = service_instances.ServiceInstanceLastOperationGetHandlerFunc(func(params service_instances.ServiceInstanceLastOperationGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation service_instances.ServiceInstanceLastOperationGet has not yet been implemented")
		})
	}

	if api.ServiceInstancesServiceInstanceUpdateHandler == nil {
		api.ServiceInstancesServiceInstanceUpdateHandler = service_instances.ServiceInstanceUpdateHandlerFunc(func(params service_instances.ServiceInstanceUpdateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation service_instances.ServiceInstanceUpdate has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
