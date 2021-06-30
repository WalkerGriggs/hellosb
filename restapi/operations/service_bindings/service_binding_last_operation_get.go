// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ServiceBindingLastOperationGetHandlerFunc turns a function with the right signature into a service binding last operation get handler
type ServiceBindingLastOperationGetHandlerFunc func(ServiceBindingLastOperationGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceBindingLastOperationGetHandlerFunc) Handle(params ServiceBindingLastOperationGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ServiceBindingLastOperationGetHandler interface for that can handle valid service binding last operation get params
type ServiceBindingLastOperationGetHandler interface {
	Handle(ServiceBindingLastOperationGetParams, interface{}) middleware.Responder
}

// NewServiceBindingLastOperationGet creates a new http.Handler for the service binding last operation get operation
func NewServiceBindingLastOperationGet(ctx *middleware.Context, handler ServiceBindingLastOperationGetHandler) *ServiceBindingLastOperationGet {
	return &ServiceBindingLastOperationGet{Context: ctx, Handler: handler}
}

/* ServiceBindingLastOperationGet swagger:route GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation ServiceBindings serviceBindingLastOperationGet

last requested operation state for service binding

*/
type ServiceBindingLastOperationGet struct {
	Context *middleware.Context
	Handler ServiceBindingLastOperationGetHandler
}

func (o *ServiceBindingLastOperationGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewServiceBindingLastOperationGetParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
