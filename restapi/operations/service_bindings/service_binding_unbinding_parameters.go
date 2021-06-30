// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewServiceBindingUnbindingParams creates a new ServiceBindingUnbindingParams object
//
// There are no default values defined in the spec.
func NewServiceBindingUnbindingParams() ServiceBindingUnbindingParams {

	return ServiceBindingUnbindingParams{}
}

// ServiceBindingUnbindingParams contains all the bound params for the service binding unbinding operation
// typically these are obtained from a http.Request
//
// swagger:parameters serviceBinding.unbinding
type ServiceBindingUnbindingParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*identity of the user that initiated the request from the Platform
	  In: header
	*/
	XBrokerAPIOriginatingIdentity *string
	/*idenity of the request from the Platform
	  In: header
	*/
	XBrokerAPIRequestIdentity *string
	/*version number of the Service Broker API that the Platform will use
	  Required: true
	  In: header
	*/
	XBrokerAPIVersion string
	/*asynchronous operations supported
	  In: query
	*/
	AcceptsIncomplete *bool
	/*binding id of binding to create
	  Required: true
	  In: path
	*/
	BindingID string
	/*instance id of instance to provision
	  Required: true
	  In: path
	*/
	InstanceID string
	/*id of the plan associated with the instance being deleted
	  Required: true
	  In: query
	*/
	PlanID string
	/*id of the service associated with the instance being deleted
	  Required: true
	  In: query
	*/
	ServiceID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewServiceBindingUnbindingParams() beforehand.
func (o *ServiceBindingUnbindingParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXBrokerAPIOriginatingIdentity(r.Header[http.CanonicalHeaderKey("X-Broker-API-Originating-Identity")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXBrokerAPIRequestIdentity(r.Header[http.CanonicalHeaderKey("X-Broker-API-Request-Identity")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXBrokerAPIVersion(r.Header[http.CanonicalHeaderKey("X-Broker-API-Version")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qAcceptsIncomplete, qhkAcceptsIncomplete, _ := qs.GetOK("accepts_incomplete")
	if err := o.bindAcceptsIncomplete(qAcceptsIncomplete, qhkAcceptsIncomplete, route.Formats); err != nil {
		res = append(res, err)
	}

	rBindingID, rhkBindingID, _ := route.Params.GetOK("binding_id")
	if err := o.bindBindingID(rBindingID, rhkBindingID, route.Formats); err != nil {
		res = append(res, err)
	}

	rInstanceID, rhkInstanceID, _ := route.Params.GetOK("instance_id")
	if err := o.bindInstanceID(rInstanceID, rhkInstanceID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPlanID, qhkPlanID, _ := qs.GetOK("plan_id")
	if err := o.bindPlanID(qPlanID, qhkPlanID, route.Formats); err != nil {
		res = append(res, err)
	}

	qServiceID, qhkServiceID, _ := qs.GetOK("service_id")
	if err := o.bindServiceID(qServiceID, qhkServiceID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXBrokerAPIOriginatingIdentity binds and validates parameter XBrokerAPIOriginatingIdentity from header.
func (o *ServiceBindingUnbindingParams) bindXBrokerAPIOriginatingIdentity(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.XBrokerAPIOriginatingIdentity = &raw

	return nil
}

// bindXBrokerAPIRequestIdentity binds and validates parameter XBrokerAPIRequestIdentity from header.
func (o *ServiceBindingUnbindingParams) bindXBrokerAPIRequestIdentity(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.XBrokerAPIRequestIdentity = &raw

	return nil
}

// bindXBrokerAPIVersion binds and validates parameter XBrokerAPIVersion from header.
func (o *ServiceBindingUnbindingParams) bindXBrokerAPIVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Broker-API-Version", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Broker-API-Version", "header", raw); err != nil {
		return err
	}
	o.XBrokerAPIVersion = raw

	return nil
}

// bindAcceptsIncomplete binds and validates parameter AcceptsIncomplete from query.
func (o *ServiceBindingUnbindingParams) bindAcceptsIncomplete(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("accepts_incomplete", "query", "bool", raw)
	}
	o.AcceptsIncomplete = &value

	return nil
}

// bindBindingID binds and validates parameter BindingID from path.
func (o *ServiceBindingUnbindingParams) bindBindingID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.BindingID = raw

	return nil
}

// bindInstanceID binds and validates parameter InstanceID from path.
func (o *ServiceBindingUnbindingParams) bindInstanceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.InstanceID = raw

	return nil
}

// bindPlanID binds and validates parameter PlanID from query.
func (o *ServiceBindingUnbindingParams) bindPlanID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("plan_id", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("plan_id", "query", raw); err != nil {
		return err
	}
	o.PlanID = raw

	return nil
}

// bindServiceID binds and validates parameter ServiceID from query.
func (o *ServiceBindingUnbindingParams) bindServiceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("service_id", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("service_id", "query", raw); err != nil {
		return err
	}
	o.ServiceID = raw

	return nil
}
