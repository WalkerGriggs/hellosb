// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/walkergriggs/hellosb/models"
)

// ServiceInstanceDeprovisionOKCode is the HTTP code returned for type ServiceInstanceDeprovisionOK
const ServiceInstanceDeprovisionOKCode int = 200

/*ServiceInstanceDeprovisionOK OK

swagger:response serviceInstanceDeprovisionOK
*/
type ServiceInstanceDeprovisionOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewServiceInstanceDeprovisionOK creates ServiceInstanceDeprovisionOK with default headers values
func NewServiceInstanceDeprovisionOK() *ServiceInstanceDeprovisionOK {

	return &ServiceInstanceDeprovisionOK{}
}

// WithPayload adds the payload to the service instance deprovision o k response
func (o *ServiceInstanceDeprovisionOK) WithPayload(payload interface{}) *ServiceInstanceDeprovisionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance deprovision o k response
func (o *ServiceInstanceDeprovisionOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceDeprovisionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ServiceInstanceDeprovisionAcceptedCode is the HTTP code returned for type ServiceInstanceDeprovisionAccepted
const ServiceInstanceDeprovisionAcceptedCode int = 202

/*ServiceInstanceDeprovisionAccepted Accepted

swagger:response serviceInstanceDeprovisionAccepted
*/
type ServiceInstanceDeprovisionAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.AsyncOperation `json:"body,omitempty"`
}

// NewServiceInstanceDeprovisionAccepted creates ServiceInstanceDeprovisionAccepted with default headers values
func NewServiceInstanceDeprovisionAccepted() *ServiceInstanceDeprovisionAccepted {

	return &ServiceInstanceDeprovisionAccepted{}
}

// WithPayload adds the payload to the service instance deprovision accepted response
func (o *ServiceInstanceDeprovisionAccepted) WithPayload(payload *models.AsyncOperation) *ServiceInstanceDeprovisionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance deprovision accepted response
func (o *ServiceInstanceDeprovisionAccepted) SetPayload(payload *models.AsyncOperation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceDeprovisionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceDeprovisionBadRequestCode is the HTTP code returned for type ServiceInstanceDeprovisionBadRequest
const ServiceInstanceDeprovisionBadRequestCode int = 400

/*ServiceInstanceDeprovisionBadRequest Bad Request

swagger:response serviceInstanceDeprovisionBadRequest
*/
type ServiceInstanceDeprovisionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceDeprovisionBadRequest creates ServiceInstanceDeprovisionBadRequest with default headers values
func NewServiceInstanceDeprovisionBadRequest() *ServiceInstanceDeprovisionBadRequest {

	return &ServiceInstanceDeprovisionBadRequest{}
}

// WithPayload adds the payload to the service instance deprovision bad request response
func (o *ServiceInstanceDeprovisionBadRequest) WithPayload(payload *models.Error) *ServiceInstanceDeprovisionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance deprovision bad request response
func (o *ServiceInstanceDeprovisionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceDeprovisionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceDeprovisionUnauthorizedCode is the HTTP code returned for type ServiceInstanceDeprovisionUnauthorized
const ServiceInstanceDeprovisionUnauthorizedCode int = 401

/*ServiceInstanceDeprovisionUnauthorized Unauthorized

swagger:response serviceInstanceDeprovisionUnauthorized
*/
type ServiceInstanceDeprovisionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceDeprovisionUnauthorized creates ServiceInstanceDeprovisionUnauthorized with default headers values
func NewServiceInstanceDeprovisionUnauthorized() *ServiceInstanceDeprovisionUnauthorized {

	return &ServiceInstanceDeprovisionUnauthorized{}
}

// WithPayload adds the payload to the service instance deprovision unauthorized response
func (o *ServiceInstanceDeprovisionUnauthorized) WithPayload(payload *models.Error) *ServiceInstanceDeprovisionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance deprovision unauthorized response
func (o *ServiceInstanceDeprovisionUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceDeprovisionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceDeprovisionGoneCode is the HTTP code returned for type ServiceInstanceDeprovisionGone
const ServiceInstanceDeprovisionGoneCode int = 410

/*ServiceInstanceDeprovisionGone Gone

swagger:response serviceInstanceDeprovisionGone
*/
type ServiceInstanceDeprovisionGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceDeprovisionGone creates ServiceInstanceDeprovisionGone with default headers values
func NewServiceInstanceDeprovisionGone() *ServiceInstanceDeprovisionGone {

	return &ServiceInstanceDeprovisionGone{}
}

// WithPayload adds the payload to the service instance deprovision gone response
func (o *ServiceInstanceDeprovisionGone) WithPayload(payload *models.Error) *ServiceInstanceDeprovisionGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance deprovision gone response
func (o *ServiceInstanceDeprovisionGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceDeprovisionGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceDeprovisionPreconditionFailedCode is the HTTP code returned for type ServiceInstanceDeprovisionPreconditionFailed
const ServiceInstanceDeprovisionPreconditionFailedCode int = 412

/*ServiceInstanceDeprovisionPreconditionFailed Precondition Failed

swagger:response serviceInstanceDeprovisionPreconditionFailed
*/
type ServiceInstanceDeprovisionPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceDeprovisionPreconditionFailed creates ServiceInstanceDeprovisionPreconditionFailed with default headers values
func NewServiceInstanceDeprovisionPreconditionFailed() *ServiceInstanceDeprovisionPreconditionFailed {

	return &ServiceInstanceDeprovisionPreconditionFailed{}
}

// WithPayload adds the payload to the service instance deprovision precondition failed response
func (o *ServiceInstanceDeprovisionPreconditionFailed) WithPayload(payload *models.Error) *ServiceInstanceDeprovisionPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance deprovision precondition failed response
func (o *ServiceInstanceDeprovisionPreconditionFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceDeprovisionPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceDeprovisionUnprocessableEntityCode is the HTTP code returned for type ServiceInstanceDeprovisionUnprocessableEntity
const ServiceInstanceDeprovisionUnprocessableEntityCode int = 422

/*ServiceInstanceDeprovisionUnprocessableEntity Unprocessable Entity

swagger:response serviceInstanceDeprovisionUnprocessableEntity
*/
type ServiceInstanceDeprovisionUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceDeprovisionUnprocessableEntity creates ServiceInstanceDeprovisionUnprocessableEntity with default headers values
func NewServiceInstanceDeprovisionUnprocessableEntity() *ServiceInstanceDeprovisionUnprocessableEntity {

	return &ServiceInstanceDeprovisionUnprocessableEntity{}
}

// WithPayload adds the payload to the service instance deprovision unprocessable entity response
func (o *ServiceInstanceDeprovisionUnprocessableEntity) WithPayload(payload *models.Error) *ServiceInstanceDeprovisionUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance deprovision unprocessable entity response
func (o *ServiceInstanceDeprovisionUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceDeprovisionUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ServiceInstanceDeprovisionDefault Unexpected

swagger:response serviceInstanceDeprovisionDefault
*/
type ServiceInstanceDeprovisionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceDeprovisionDefault creates ServiceInstanceDeprovisionDefault with default headers values
func NewServiceInstanceDeprovisionDefault(code int) *ServiceInstanceDeprovisionDefault {
	if code <= 0 {
		code = 500
	}

	return &ServiceInstanceDeprovisionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the service instance deprovision default response
func (o *ServiceInstanceDeprovisionDefault) WithStatusCode(code int) *ServiceInstanceDeprovisionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the service instance deprovision default response
func (o *ServiceInstanceDeprovisionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the service instance deprovision default response
func (o *ServiceInstanceDeprovisionDefault) WithPayload(payload *models.Error) *ServiceInstanceDeprovisionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance deprovision default response
func (o *ServiceInstanceDeprovisionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceDeprovisionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
