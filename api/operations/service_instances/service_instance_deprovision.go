// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ServiceInstanceDeprovisionHandlerFunc turns a function with the right signature into a service instance deprovision handler
type ServiceInstanceDeprovisionHandlerFunc func(ServiceInstanceDeprovisionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceInstanceDeprovisionHandlerFunc) Handle(params ServiceInstanceDeprovisionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ServiceInstanceDeprovisionHandler interface for that can handle valid service instance deprovision params
type ServiceInstanceDeprovisionHandler interface {
	Handle(ServiceInstanceDeprovisionParams, interface{}) middleware.Responder
}

// NewServiceInstanceDeprovision creates a new http.Handler for the service instance deprovision operation
func NewServiceInstanceDeprovision(ctx *middleware.Context, handler ServiceInstanceDeprovisionHandler) *ServiceInstanceDeprovision {
	return &ServiceInstanceDeprovision{Context: ctx, Handler: handler}
}

/* ServiceInstanceDeprovision swagger:route DELETE /v2/service_instances/{instance_id} ServiceInstances serviceInstanceDeprovision

deprovision a service instance

*/
type ServiceInstanceDeprovision struct {
	Context *middleware.Context
	Handler ServiceInstanceDeprovisionHandler
}

func (o *ServiceInstanceDeprovision) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewServiceInstanceDeprovisionParams()
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
