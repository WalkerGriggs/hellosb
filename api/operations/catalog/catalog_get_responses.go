// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/walkergriggs/hellosb/models"
)

// CatalogGetOKCode is the HTTP code returned for type CatalogGetOK
const CatalogGetOKCode int = 200

/*CatalogGetOK catalog response

swagger:response catalogGetOK
*/
type CatalogGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Catalog `json:"body,omitempty"`
}

// NewCatalogGetOK creates CatalogGetOK with default headers values
func NewCatalogGetOK() *CatalogGetOK {

	return &CatalogGetOK{}
}

// WithPayload adds the payload to the catalog get o k response
func (o *CatalogGetOK) WithPayload(payload *models.Catalog) *CatalogGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the catalog get o k response
func (o *CatalogGetOK) SetPayload(payload *models.Catalog) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CatalogGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CatalogGetUnauthorizedCode is the HTTP code returned for type CatalogGetUnauthorized
const CatalogGetUnauthorizedCode int = 401

/*CatalogGetUnauthorized Unauthorized

swagger:response catalogGetUnauthorized
*/
type CatalogGetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCatalogGetUnauthorized creates CatalogGetUnauthorized with default headers values
func NewCatalogGetUnauthorized() *CatalogGetUnauthorized {

	return &CatalogGetUnauthorized{}
}

// WithPayload adds the payload to the catalog get unauthorized response
func (o *CatalogGetUnauthorized) WithPayload(payload *models.Error) *CatalogGetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the catalog get unauthorized response
func (o *CatalogGetUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CatalogGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CatalogGetPreconditionFailedCode is the HTTP code returned for type CatalogGetPreconditionFailed
const CatalogGetPreconditionFailedCode int = 412

/*CatalogGetPreconditionFailed Precondition Failed

swagger:response catalogGetPreconditionFailed
*/
type CatalogGetPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCatalogGetPreconditionFailed creates CatalogGetPreconditionFailed with default headers values
func NewCatalogGetPreconditionFailed() *CatalogGetPreconditionFailed {

	return &CatalogGetPreconditionFailed{}
}

// WithPayload adds the payload to the catalog get precondition failed response
func (o *CatalogGetPreconditionFailed) WithPayload(payload *models.Error) *CatalogGetPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the catalog get precondition failed response
func (o *CatalogGetPreconditionFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CatalogGetPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CatalogGetDefault Unexpected

swagger:response catalogGetDefault
*/
type CatalogGetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCatalogGetDefault creates CatalogGetDefault with default headers values
func NewCatalogGetDefault(code int) *CatalogGetDefault {
	if code <= 0 {
		code = 500
	}

	return &CatalogGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the catalog get default response
func (o *CatalogGetDefault) WithStatusCode(code int) *CatalogGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the catalog get default response
func (o *CatalogGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the catalog get default response
func (o *CatalogGetDefault) WithPayload(payload *models.Error) *CatalogGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the catalog get default response
func (o *CatalogGetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CatalogGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}