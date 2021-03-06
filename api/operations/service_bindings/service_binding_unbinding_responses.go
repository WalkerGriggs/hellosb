// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/walkergriggs/hellosb/models"
)

// ServiceBindingUnbindingOKCode is the HTTP code returned for type ServiceBindingUnbindingOK
const ServiceBindingUnbindingOKCode int = 200

/*ServiceBindingUnbindingOK OK

swagger:response serviceBindingUnbindingOK
*/
type ServiceBindingUnbindingOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewServiceBindingUnbindingOK creates ServiceBindingUnbindingOK with default headers values
func NewServiceBindingUnbindingOK() *ServiceBindingUnbindingOK {

	return &ServiceBindingUnbindingOK{}
}

// WithPayload adds the payload to the service binding unbinding o k response
func (o *ServiceBindingUnbindingOK) WithPayload(payload interface{}) *ServiceBindingUnbindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding o k response
func (o *ServiceBindingUnbindingOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ServiceBindingUnbindingAcceptedCode is the HTTP code returned for type ServiceBindingUnbindingAccepted
const ServiceBindingUnbindingAcceptedCode int = 202

/*ServiceBindingUnbindingAccepted Accepted

swagger:response serviceBindingUnbindingAccepted
*/
type ServiceBindingUnbindingAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.AsyncOperation `json:"body,omitempty"`
}

// NewServiceBindingUnbindingAccepted creates ServiceBindingUnbindingAccepted with default headers values
func NewServiceBindingUnbindingAccepted() *ServiceBindingUnbindingAccepted {

	return &ServiceBindingUnbindingAccepted{}
}

// WithPayload adds the payload to the service binding unbinding accepted response
func (o *ServiceBindingUnbindingAccepted) WithPayload(payload *models.AsyncOperation) *ServiceBindingUnbindingAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding accepted response
func (o *ServiceBindingUnbindingAccepted) SetPayload(payload *models.AsyncOperation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingUnbindingBadRequestCode is the HTTP code returned for type ServiceBindingUnbindingBadRequest
const ServiceBindingUnbindingBadRequestCode int = 400

/*ServiceBindingUnbindingBadRequest Bad Request

swagger:response serviceBindingUnbindingBadRequest
*/
type ServiceBindingUnbindingBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingUnbindingBadRequest creates ServiceBindingUnbindingBadRequest with default headers values
func NewServiceBindingUnbindingBadRequest() *ServiceBindingUnbindingBadRequest {

	return &ServiceBindingUnbindingBadRequest{}
}

// WithPayload adds the payload to the service binding unbinding bad request response
func (o *ServiceBindingUnbindingBadRequest) WithPayload(payload *models.Error) *ServiceBindingUnbindingBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding bad request response
func (o *ServiceBindingUnbindingBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingUnbindingUnauthorizedCode is the HTTP code returned for type ServiceBindingUnbindingUnauthorized
const ServiceBindingUnbindingUnauthorizedCode int = 401

/*ServiceBindingUnbindingUnauthorized Unauthorized

swagger:response serviceBindingUnbindingUnauthorized
*/
type ServiceBindingUnbindingUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingUnbindingUnauthorized creates ServiceBindingUnbindingUnauthorized with default headers values
func NewServiceBindingUnbindingUnauthorized() *ServiceBindingUnbindingUnauthorized {

	return &ServiceBindingUnbindingUnauthorized{}
}

// WithPayload adds the payload to the service binding unbinding unauthorized response
func (o *ServiceBindingUnbindingUnauthorized) WithPayload(payload *models.Error) *ServiceBindingUnbindingUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding unauthorized response
func (o *ServiceBindingUnbindingUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingUnbindingGoneCode is the HTTP code returned for type ServiceBindingUnbindingGone
const ServiceBindingUnbindingGoneCode int = 410

/*ServiceBindingUnbindingGone Gone

swagger:response serviceBindingUnbindingGone
*/
type ServiceBindingUnbindingGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingUnbindingGone creates ServiceBindingUnbindingGone with default headers values
func NewServiceBindingUnbindingGone() *ServiceBindingUnbindingGone {

	return &ServiceBindingUnbindingGone{}
}

// WithPayload adds the payload to the service binding unbinding gone response
func (o *ServiceBindingUnbindingGone) WithPayload(payload *models.Error) *ServiceBindingUnbindingGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding gone response
func (o *ServiceBindingUnbindingGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingUnbindingPreconditionFailedCode is the HTTP code returned for type ServiceBindingUnbindingPreconditionFailed
const ServiceBindingUnbindingPreconditionFailedCode int = 412

/*ServiceBindingUnbindingPreconditionFailed Precondition Failed

swagger:response serviceBindingUnbindingPreconditionFailed
*/
type ServiceBindingUnbindingPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingUnbindingPreconditionFailed creates ServiceBindingUnbindingPreconditionFailed with default headers values
func NewServiceBindingUnbindingPreconditionFailed() *ServiceBindingUnbindingPreconditionFailed {

	return &ServiceBindingUnbindingPreconditionFailed{}
}

// WithPayload adds the payload to the service binding unbinding precondition failed response
func (o *ServiceBindingUnbindingPreconditionFailed) WithPayload(payload *models.Error) *ServiceBindingUnbindingPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding precondition failed response
func (o *ServiceBindingUnbindingPreconditionFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingUnbindingUnprocessableEntityCode is the HTTP code returned for type ServiceBindingUnbindingUnprocessableEntity
const ServiceBindingUnbindingUnprocessableEntityCode int = 422

/*ServiceBindingUnbindingUnprocessableEntity Unprocessable Entity

swagger:response serviceBindingUnbindingUnprocessableEntity
*/
type ServiceBindingUnbindingUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingUnbindingUnprocessableEntity creates ServiceBindingUnbindingUnprocessableEntity with default headers values
func NewServiceBindingUnbindingUnprocessableEntity() *ServiceBindingUnbindingUnprocessableEntity {

	return &ServiceBindingUnbindingUnprocessableEntity{}
}

// WithPayload adds the payload to the service binding unbinding unprocessable entity response
func (o *ServiceBindingUnbindingUnprocessableEntity) WithPayload(payload *models.Error) *ServiceBindingUnbindingUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding unprocessable entity response
func (o *ServiceBindingUnbindingUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ServiceBindingUnbindingDefault Unexpected

swagger:response serviceBindingUnbindingDefault
*/
type ServiceBindingUnbindingDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingUnbindingDefault creates ServiceBindingUnbindingDefault with default headers values
func NewServiceBindingUnbindingDefault(code int) *ServiceBindingUnbindingDefault {
	if code <= 0 {
		code = 500
	}

	return &ServiceBindingUnbindingDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the service binding unbinding default response
func (o *ServiceBindingUnbindingDefault) WithStatusCode(code int) *ServiceBindingUnbindingDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the service binding unbinding default response
func (o *ServiceBindingUnbindingDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the service binding unbinding default response
func (o *ServiceBindingUnbindingDefault) WithPayload(payload *models.Error) *ServiceBindingUnbindingDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding default response
func (o *ServiceBindingUnbindingDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
