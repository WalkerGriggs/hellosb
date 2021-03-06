// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/walkergriggs/hellosb/models"
)

// NewServiceBindingBindingParams creates a new ServiceBindingBindingParams object
//
// There are no default values defined in the spec.
func NewServiceBindingBindingParams() ServiceBindingBindingParams {

	return ServiceBindingBindingParams{}
}

// ServiceBindingBindingParams contains all the bound params for the service binding binding operation
// typically these are obtained from a http.Request
//
// swagger:parameters serviceBinding.binding
type ServiceBindingBindingParams struct {

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
	/*parameters for the requested service binding
	  Required: true
	  In: body
	*/
	Body *models.ServiceBindingRequest
	/*instance id of instance to provision
	  Required: true
	  In: path
	*/
	InstanceID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewServiceBindingBindingParams() beforehand.
func (o *ServiceBindingBindingParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ServiceBindingRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}

	rInstanceID, rhkInstanceID, _ := route.Params.GetOK("instance_id")
	if err := o.bindInstanceID(rInstanceID, rhkInstanceID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXBrokerAPIOriginatingIdentity binds and validates parameter XBrokerAPIOriginatingIdentity from header.
func (o *ServiceBindingBindingParams) bindXBrokerAPIOriginatingIdentity(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ServiceBindingBindingParams) bindXBrokerAPIRequestIdentity(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ServiceBindingBindingParams) bindXBrokerAPIVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ServiceBindingBindingParams) bindAcceptsIncomplete(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ServiceBindingBindingParams) bindBindingID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ServiceBindingBindingParams) bindInstanceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.InstanceID = raw

	return nil
}
