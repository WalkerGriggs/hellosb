// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/walkergriggs/hellosb/models"
)

// ServiceBindingLastOperationGetOKCode is the HTTP code returned for type ServiceBindingLastOperationGetOK
const ServiceBindingLastOperationGetOKCode int = 200

/*ServiceBindingLastOperationGetOK OK

swagger:response serviceBindingLastOperationGetOK
*/
type ServiceBindingLastOperationGetOK struct {
	/*Indicates when to retry the request

	 */
	RetryAfter string `json:"RetryAfter"`

	/*
	  In: Body
	*/
	Payload *models.LastOperationResource `json:"body,omitempty"`
}

// NewServiceBindingLastOperationGetOK creates ServiceBindingLastOperationGetOK with default headers values
func NewServiceBindingLastOperationGetOK() *ServiceBindingLastOperationGetOK {

	return &ServiceBindingLastOperationGetOK{}
}

// WithRetryAfter adds the retryAfter to the service binding last operation get o k response
func (o *ServiceBindingLastOperationGetOK) WithRetryAfter(retryAfter string) *ServiceBindingLastOperationGetOK {
	o.RetryAfter = retryAfter
	return o
}

// SetRetryAfter sets the retryAfter to the service binding last operation get o k response
func (o *ServiceBindingLastOperationGetOK) SetRetryAfter(retryAfter string) {
	o.RetryAfter = retryAfter
}

// WithPayload adds the payload to the service binding last operation get o k response
func (o *ServiceBindingLastOperationGetOK) WithPayload(payload *models.LastOperationResource) *ServiceBindingLastOperationGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding last operation get o k response
func (o *ServiceBindingLastOperationGetOK) SetPayload(payload *models.LastOperationResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingLastOperationGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header RetryAfter

	retryAfter := o.RetryAfter
	if retryAfter != "" {
		rw.Header().Set("RetryAfter", retryAfter)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingLastOperationGetBadRequestCode is the HTTP code returned for type ServiceBindingLastOperationGetBadRequest
const ServiceBindingLastOperationGetBadRequestCode int = 400

/*ServiceBindingLastOperationGetBadRequest Bad Request

swagger:response serviceBindingLastOperationGetBadRequest
*/
type ServiceBindingLastOperationGetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingLastOperationGetBadRequest creates ServiceBindingLastOperationGetBadRequest with default headers values
func NewServiceBindingLastOperationGetBadRequest() *ServiceBindingLastOperationGetBadRequest {

	return &ServiceBindingLastOperationGetBadRequest{}
}

// WithPayload adds the payload to the service binding last operation get bad request response
func (o *ServiceBindingLastOperationGetBadRequest) WithPayload(payload *models.Error) *ServiceBindingLastOperationGetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding last operation get bad request response
func (o *ServiceBindingLastOperationGetBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingLastOperationGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingLastOperationGetUnauthorizedCode is the HTTP code returned for type ServiceBindingLastOperationGetUnauthorized
const ServiceBindingLastOperationGetUnauthorizedCode int = 401

/*ServiceBindingLastOperationGetUnauthorized Unauthorized

swagger:response serviceBindingLastOperationGetUnauthorized
*/
type ServiceBindingLastOperationGetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingLastOperationGetUnauthorized creates ServiceBindingLastOperationGetUnauthorized with default headers values
func NewServiceBindingLastOperationGetUnauthorized() *ServiceBindingLastOperationGetUnauthorized {

	return &ServiceBindingLastOperationGetUnauthorized{}
}

// WithPayload adds the payload to the service binding last operation get unauthorized response
func (o *ServiceBindingLastOperationGetUnauthorized) WithPayload(payload *models.Error) *ServiceBindingLastOperationGetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding last operation get unauthorized response
func (o *ServiceBindingLastOperationGetUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingLastOperationGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingLastOperationGetNotFoundCode is the HTTP code returned for type ServiceBindingLastOperationGetNotFound
const ServiceBindingLastOperationGetNotFoundCode int = 404

/*ServiceBindingLastOperationGetNotFound Not Found

swagger:response serviceBindingLastOperationGetNotFound
*/
type ServiceBindingLastOperationGetNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingLastOperationGetNotFound creates ServiceBindingLastOperationGetNotFound with default headers values
func NewServiceBindingLastOperationGetNotFound() *ServiceBindingLastOperationGetNotFound {

	return &ServiceBindingLastOperationGetNotFound{}
}

// WithPayload adds the payload to the service binding last operation get not found response
func (o *ServiceBindingLastOperationGetNotFound) WithPayload(payload *models.Error) *ServiceBindingLastOperationGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding last operation get not found response
func (o *ServiceBindingLastOperationGetNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingLastOperationGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingLastOperationGetGoneCode is the HTTP code returned for type ServiceBindingLastOperationGetGone
const ServiceBindingLastOperationGetGoneCode int = 410

/*ServiceBindingLastOperationGetGone Gone

swagger:response serviceBindingLastOperationGetGone
*/
type ServiceBindingLastOperationGetGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingLastOperationGetGone creates ServiceBindingLastOperationGetGone with default headers values
func NewServiceBindingLastOperationGetGone() *ServiceBindingLastOperationGetGone {

	return &ServiceBindingLastOperationGetGone{}
}

// WithPayload adds the payload to the service binding last operation get gone response
func (o *ServiceBindingLastOperationGetGone) WithPayload(payload *models.Error) *ServiceBindingLastOperationGetGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding last operation get gone response
func (o *ServiceBindingLastOperationGetGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingLastOperationGetGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingLastOperationGetPreconditionFailedCode is the HTTP code returned for type ServiceBindingLastOperationGetPreconditionFailed
const ServiceBindingLastOperationGetPreconditionFailedCode int = 412

/*ServiceBindingLastOperationGetPreconditionFailed Precondition Failed

swagger:response serviceBindingLastOperationGetPreconditionFailed
*/
type ServiceBindingLastOperationGetPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingLastOperationGetPreconditionFailed creates ServiceBindingLastOperationGetPreconditionFailed with default headers values
func NewServiceBindingLastOperationGetPreconditionFailed() *ServiceBindingLastOperationGetPreconditionFailed {

	return &ServiceBindingLastOperationGetPreconditionFailed{}
}

// WithPayload adds the payload to the service binding last operation get precondition failed response
func (o *ServiceBindingLastOperationGetPreconditionFailed) WithPayload(payload *models.Error) *ServiceBindingLastOperationGetPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding last operation get precondition failed response
func (o *ServiceBindingLastOperationGetPreconditionFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingLastOperationGetPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ServiceBindingLastOperationGetDefault Unexpected

swagger:response serviceBindingLastOperationGetDefault
*/
type ServiceBindingLastOperationGetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingLastOperationGetDefault creates ServiceBindingLastOperationGetDefault with default headers values
func NewServiceBindingLastOperationGetDefault(code int) *ServiceBindingLastOperationGetDefault {
	if code <= 0 {
		code = 500
	}

	return &ServiceBindingLastOperationGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the service binding last operation get default response
func (o *ServiceBindingLastOperationGetDefault) WithStatusCode(code int) *ServiceBindingLastOperationGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the service binding last operation get default response
func (o *ServiceBindingLastOperationGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the service binding last operation get default response
func (o *ServiceBindingLastOperationGetDefault) WithPayload(payload *models.Error) *ServiceBindingLastOperationGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding last operation get default response
func (o *ServiceBindingLastOperationGetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingLastOperationGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
