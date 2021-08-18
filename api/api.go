package api

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"

	"github.com/walkergriggs/hellosb/api/operations/catalog"
	"github.com/walkergriggs/hellosb/api/operations/service_bindings"
	"github.com/walkergriggs/hellosb/api/operations/service_instances"
	"github.com/walkergriggs/hellosb/state"
)

type API struct {
	config   *Config
	store    *state.StateStore
	spec     *loads.Document
	context  *middleware.Context
	formats  strfmt.Registry
	handlers map[hKey]http.Handler

	JSONConsumer runtime.Consumer
	JSONProducer runtime.Producer

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// BasicAuthAuth registers a function that takes username and password and returns a principal
	// it performs authentication with basic auth
	BasicAuthFunc func(string, string) (interface{}, error)

	// Catalog Handlers
	CatalogGetHandler catalog.CatalogGetHandler

	// ServiceBinding Handlers
	ServiceBindingBindingHandler          service_bindings.ServiceBindingBindingHandler
	ServiceBindingGetHandler              service_bindings.ServiceBindingGetHandler
	ServiceBindingLastOperationGetHandler service_bindings.ServiceBindingLastOperationGetHandler
	ServiceBindingUnbindingHandler        service_bindings.ServiceBindingUnbindingHandler

	// ServiceInstance Handlers
	ServiceInstanceDeprovisionHandler      service_instances.ServiceInstanceDeprovisionHandler
	ServiceInstanceGetHandler              service_instances.ServiceInstanceGetHandler
	ServiceInstanceLastOperationGetHandler service_instances.ServiceInstanceLastOperationGetHandler
	ServiceInstanceProvisionHandler        service_instances.ServiceInstanceProvisionHandler
	ServiceInstanceUpdateHandler           service_instances.ServiceInstanceUpdateHandler
}

func New(spec *loads.Document) *API {
	api := &API{
		BasicAuthenticator: security.BasicAuth,
		BasicAuthFunc: func(user string, pass string) (interface{}, error) {
			return "TODO", nil // TODO
		},

		JSONConsumer: runtime.JSONConsumer(),
		JSONProducer: runtime.JSONProducer(),

		spec:       spec,
		formats:    strfmt.Default,
		ServeError: errors.ServeError,
		handlers:   make(map[hKey]http.Handler),
	}

	return api
}

func (api *API) Context() *middleware.Context {
	if api.context == nil {
		api.context = middleware.NewRoutableContext(api.spec, api, nil)
	}

	return api.context
}

// Implements RoutableAPI interface -- https://pkg.go.dev/github.com/go-openapi/runtime/middleware#RoutableAPI
func (api *API) HandlerFor(method, path string) (http.Handler, bool) {
	if api.handlers == nil {
		return nil, false
	}

	h, ok := api.handlers[hKey{method, path}]
	return h, ok
}

// Implements RoutableAPI interface -- https://pkg.go.dev/github.com/go-openapi/runtime/middleware#RoutableAPI
func (api *API) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return api.ServeError
}

// Implements RoutableAPI interface -- https://pkg.go.dev/github.com/go-openapi/runtime/middleware#RoutableAPI
func (api *API) ConsumersFor(types []string) map[string]runtime.Consumer {
	return map[string]runtime.Consumer{
		"application/json": api.JSONConsumer,
	}
}

// Implements RoutableAPI interface -- https://pkg.go.dev/github.com/go-openapi/runtime/middleware#RoutableAPI
func (api *API) ProducersFor(types []string) map[string]runtime.Producer {
	return map[string]runtime.Producer{
		"application/json": api.JSONProducer,
	}
}

// Implements RoutableAPI interface -- https://pkg.go.dev/github.com/go-openapi/runtime/middleware#RoutableAPI
func (api *API) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return map[string]runtime.Authenticator{
		"basicAuth": api.BasicAuthenticator(api.BasicAuthFunc),
	}
}

// Implements RoutableAPI interface -- https://pkg.go.dev/github.com/go-openapi/runtime/middleware#RoutableAPI
func (api *API) Authorizer() runtime.Authorizer {
	return security.Authorized()
}

// Implements RoutableAPI interface -- https://pkg.go.dev/github.com/go-openapi/runtime/middleware#RoutableAPI
func (api *API) Formats() strfmt.Registry {
	return api.formats
}

// Implements RoutableAPI interface -- https://pkg.go.dev/github.com/go-openapi/runtime/middleware#RoutableAPI
func (api *API) DefaultProduces() string {
	return "application/json"
}

// Implements RoutableAPI interface -- https://pkg.go.dev/github.com/go-openapi/runtime/middleware#RoutableAPI
func (api *API) DefaultConsumes() string {
	return "application/json"
}

func (api *API) Serve(builder middleware.Builder) http.Handler {
	api.Init()
	return api.context.APIHandler(builder)
}

func (api *API) Init() {
	api.Context()

	api.handlers = map[hKey]http.Handler{
		{"GET", "/v2/catalog"}: catalog.NewCatalogGet(api.context, api.CatalogGetHandler),
		{"GET", "/v2/service_instances/{instance_id}/service_bindings/{binding_id}"}:                service_bindings.NewServiceBindingGet(api.context, api.ServiceBindingGetHandler),
		{"GET", "/v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation"}: service_bindings.NewServiceBindingLastOperationGet(api.context, api.ServiceBindingLastOperationGetHandler),
		{"GET", "/v2/service_instances/{instance_id}"}:                                              service_instances.NewServiceInstanceGet(api.context, api.ServiceInstanceGetHandler),
		{"GET", "/v2/service_instances/{instance_id}/last_operation"}:                               service_instances.NewServiceInstanceLastOperationGet(api.context, api.ServiceInstanceLastOperationGetHandler),
		{"PUT", "/v2/service_instances/{instance_id}/service_bindings/{binding_id}"}:                service_bindings.NewServiceBindingBinding(api.context, api.ServiceBindingBindingHandler),
		{"PUT", "/v2/service_instances/{instance_id}"}:                                              service_instances.NewServiceInstanceProvision(api.context, api.ServiceInstanceProvisionHandler),
		{"DELETE", "/v2/service_instances/{instance_id}/service_bindings/{binding_id}"}:             service_bindings.NewServiceBindingUnbinding(api.context, api.ServiceBindingUnbindingHandler),
		{"DELETE", "/v2/service_instances/{instance_id}"}:                                           service_instances.NewServiceInstanceDeprovision(api.context, api.ServiceInstanceDeprovisionHandler),
		{"PATCH", "/v2/service_instances/{instance_id}"}:                                            service_instances.NewServiceInstanceUpdate(api.context, api.ServiceInstanceUpdateHandler),
	}
}

type hKey struct {
	method string
	path   string
}
