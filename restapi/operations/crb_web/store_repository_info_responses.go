package crb_web

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"crb/models"
)

/*StoreRepositoryInfoCreated Successfully created a filesystem repo info in Metadata-DB

swagger:response storeRepositoryInfoCreated
*/
type StoreRepositoryInfoCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RepositoryInstance `json:"body,omitempty"`
}

// NewStoreRepositoryInfoCreated creates StoreRepositoryInfoCreated with default headers values
func NewStoreRepositoryInfoCreated() *StoreRepositoryInfoCreated {
	return &StoreRepositoryInfoCreated{}
}

// WithPayload adds the payload to the store repository info created response
func (o *StoreRepositoryInfoCreated) WithPayload(payload *models.RepositoryInstance) *StoreRepositoryInfoCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the store repository info created response
func (o *StoreRepositoryInfoCreated) SetPayload(payload *models.RepositoryInstance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StoreRepositoryInfoCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*StoreRepositoryInfoUnprocessableEntity Unprocessable Entity - Most likely either the addr or user fields are empty

swagger:response storeRepositoryInfoUnprocessableEntity
*/
type StoreRepositoryInfoUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewStoreRepositoryInfoUnprocessableEntity creates StoreRepositoryInfoUnprocessableEntity with default headers values
func NewStoreRepositoryInfoUnprocessableEntity() *StoreRepositoryInfoUnprocessableEntity {
	return &StoreRepositoryInfoUnprocessableEntity{}
}

// WithPayload adds the payload to the store repository info unprocessable entity response
func (o *StoreRepositoryInfoUnprocessableEntity) WithPayload(payload *models.Error) *StoreRepositoryInfoUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the store repository info unprocessable entity response
func (o *StoreRepositoryInfoUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StoreRepositoryInfoUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*StoreRepositoryInfoInternalServerError Internal error - Most likely Metadata-DB not accessible

swagger:response storeRepositoryInfoInternalServerError
*/
type StoreRepositoryInfoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewStoreRepositoryInfoInternalServerError creates StoreRepositoryInfoInternalServerError with default headers values
func NewStoreRepositoryInfoInternalServerError() *StoreRepositoryInfoInternalServerError {
	return &StoreRepositoryInfoInternalServerError{}
}

// WithPayload adds the payload to the store repository info internal server error response
func (o *StoreRepositoryInfoInternalServerError) WithPayload(payload *models.Error) *StoreRepositoryInfoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the store repository info internal server error response
func (o *StoreRepositoryInfoInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StoreRepositoryInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*StoreRepositoryInfoDefault unexpected error

swagger:response storeRepositoryInfoDefault
*/
type StoreRepositoryInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewStoreRepositoryInfoDefault creates StoreRepositoryInfoDefault with default headers values
func NewStoreRepositoryInfoDefault(code int) *StoreRepositoryInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &StoreRepositoryInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the store repository info default response
func (o *StoreRepositoryInfoDefault) WithStatusCode(code int) *StoreRepositoryInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the store repository info default response
func (o *StoreRepositoryInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the store repository info default response
func (o *StoreRepositoryInfoDefault) WithPayload(payload *models.Error) *StoreRepositoryInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the store repository info default response
func (o *StoreRepositoryInfoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StoreRepositoryInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
