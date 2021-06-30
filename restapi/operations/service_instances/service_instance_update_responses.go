// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/walkergriggs/hellosb/models"
)

// ServiceInstanceUpdateOKCode is the HTTP code returned for type ServiceInstanceUpdateOK
const ServiceInstanceUpdateOKCode int = 200

/*ServiceInstanceUpdateOK OK

swagger:response serviceInstanceUpdateOK
*/
type ServiceInstanceUpdateOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewServiceInstanceUpdateOK creates ServiceInstanceUpdateOK with default headers values
func NewServiceInstanceUpdateOK() *ServiceInstanceUpdateOK {

	return &ServiceInstanceUpdateOK{}
}

// WithPayload adds the payload to the service instance update o k response
func (o *ServiceInstanceUpdateOK) WithPayload(payload interface{}) *ServiceInstanceUpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance update o k response
func (o *ServiceInstanceUpdateOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ServiceInstanceUpdateAcceptedCode is the HTTP code returned for type ServiceInstanceUpdateAccepted
const ServiceInstanceUpdateAcceptedCode int = 202

/*ServiceInstanceUpdateAccepted Accepted

swagger:response serviceInstanceUpdateAccepted
*/
type ServiceInstanceUpdateAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ServiceInstanceAsyncOperation `json:"body,omitempty"`
}

// NewServiceInstanceUpdateAccepted creates ServiceInstanceUpdateAccepted with default headers values
func NewServiceInstanceUpdateAccepted() *ServiceInstanceUpdateAccepted {

	return &ServiceInstanceUpdateAccepted{}
}

// WithPayload adds the payload to the service instance update accepted response
func (o *ServiceInstanceUpdateAccepted) WithPayload(payload *models.ServiceInstanceAsyncOperation) *ServiceInstanceUpdateAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance update accepted response
func (o *ServiceInstanceUpdateAccepted) SetPayload(payload *models.ServiceInstanceAsyncOperation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceUpdateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceUpdateBadRequestCode is the HTTP code returned for type ServiceInstanceUpdateBadRequest
const ServiceInstanceUpdateBadRequestCode int = 400

/*ServiceInstanceUpdateBadRequest Bad Request

swagger:response serviceInstanceUpdateBadRequest
*/
type ServiceInstanceUpdateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceUpdateBadRequest creates ServiceInstanceUpdateBadRequest with default headers values
func NewServiceInstanceUpdateBadRequest() *ServiceInstanceUpdateBadRequest {

	return &ServiceInstanceUpdateBadRequest{}
}

// WithPayload adds the payload to the service instance update bad request response
func (o *ServiceInstanceUpdateBadRequest) WithPayload(payload *models.Error) *ServiceInstanceUpdateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance update bad request response
func (o *ServiceInstanceUpdateBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceUpdateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceUpdateUnauthorizedCode is the HTTP code returned for type ServiceInstanceUpdateUnauthorized
const ServiceInstanceUpdateUnauthorizedCode int = 401

/*ServiceInstanceUpdateUnauthorized Unauthorized

swagger:response serviceInstanceUpdateUnauthorized
*/
type ServiceInstanceUpdateUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceUpdateUnauthorized creates ServiceInstanceUpdateUnauthorized with default headers values
func NewServiceInstanceUpdateUnauthorized() *ServiceInstanceUpdateUnauthorized {

	return &ServiceInstanceUpdateUnauthorized{}
}

// WithPayload adds the payload to the service instance update unauthorized response
func (o *ServiceInstanceUpdateUnauthorized) WithPayload(payload *models.Error) *ServiceInstanceUpdateUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance update unauthorized response
func (o *ServiceInstanceUpdateUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceUpdatePreconditionFailedCode is the HTTP code returned for type ServiceInstanceUpdatePreconditionFailed
const ServiceInstanceUpdatePreconditionFailedCode int = 412

/*ServiceInstanceUpdatePreconditionFailed Precondition Failed

swagger:response serviceInstanceUpdatePreconditionFailed
*/
type ServiceInstanceUpdatePreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceUpdatePreconditionFailed creates ServiceInstanceUpdatePreconditionFailed with default headers values
func NewServiceInstanceUpdatePreconditionFailed() *ServiceInstanceUpdatePreconditionFailed {

	return &ServiceInstanceUpdatePreconditionFailed{}
}

// WithPayload adds the payload to the service instance update precondition failed response
func (o *ServiceInstanceUpdatePreconditionFailed) WithPayload(payload *models.Error) *ServiceInstanceUpdatePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance update precondition failed response
func (o *ServiceInstanceUpdatePreconditionFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceUpdatePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceUpdateUnprocessableEntityCode is the HTTP code returned for type ServiceInstanceUpdateUnprocessableEntity
const ServiceInstanceUpdateUnprocessableEntityCode int = 422

/*ServiceInstanceUpdateUnprocessableEntity Unprocessable Entity

swagger:response serviceInstanceUpdateUnprocessableEntity
*/
type ServiceInstanceUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceUpdateUnprocessableEntity creates ServiceInstanceUpdateUnprocessableEntity with default headers values
func NewServiceInstanceUpdateUnprocessableEntity() *ServiceInstanceUpdateUnprocessableEntity {

	return &ServiceInstanceUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the service instance update unprocessable entity response
func (o *ServiceInstanceUpdateUnprocessableEntity) WithPayload(payload *models.Error) *ServiceInstanceUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance update unprocessable entity response
func (o *ServiceInstanceUpdateUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ServiceInstanceUpdateDefault Unexpected

swagger:response serviceInstanceUpdateDefault
*/
type ServiceInstanceUpdateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceUpdateDefault creates ServiceInstanceUpdateDefault with default headers values
func NewServiceInstanceUpdateDefault(code int) *ServiceInstanceUpdateDefault {
	if code <= 0 {
		code = 500
	}

	return &ServiceInstanceUpdateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the service instance update default response
func (o *ServiceInstanceUpdateDefault) WithStatusCode(code int) *ServiceInstanceUpdateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the service instance update default response
func (o *ServiceInstanceUpdateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the service instance update default response
func (o *ServiceInstanceUpdateDefault) WithPayload(payload *models.Error) *ServiceInstanceUpdateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance update default response
func (o *ServiceInstanceUpdateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceUpdateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
