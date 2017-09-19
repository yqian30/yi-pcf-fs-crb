package crb_web

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"crb/models"
)

/*CreateCopyCreated Successfully created a copy of the data and an entry in the Metadata-DB

swagger:response createCopyCreated
*/
type CreateCopyCreated struct {

	/*Copy URL
	  In: Body
	*/
	Payload *models.CopyInstance `json:"body,omitempty"`
}

// NewCreateCopyCreated creates CreateCopyCreated with default headers values
func NewCreateCopyCreated() *CreateCopyCreated {
	return &CreateCopyCreated{}
}

// WithPayload adds the payload to the create copy created response
func (o *CreateCopyCreated) WithPayload(payload *models.CopyInstance) *CreateCopyCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create copy created response
func (o *CreateCopyCreated) SetPayload(payload *models.CopyInstance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCopyCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateCopyForbidden A copy instance with the same copyId exists. Either use different copyId or delete existing copyId instance before retrying

swagger:response createCopyForbidden
*/
type CreateCopyForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateCopyForbidden creates CreateCopyForbidden with default headers values
func NewCreateCopyForbidden() *CreateCopyForbidden {
	return &CreateCopyForbidden{}
}

// WithPayload adds the payload to the create copy forbidden response
func (o *CreateCopyForbidden) WithPayload(payload *models.Error) *CreateCopyForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create copy forbidden response
func (o *CreateCopyForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCopyForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateCopyInternalServerError Internal error. Data could not be copied to copy-repo or copy instance metadata could not be created. If copy instance metadata can't be created, copy instance will be deleted

swagger:response createCopyInternalServerError
*/
type CreateCopyInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateCopyInternalServerError creates CreateCopyInternalServerError with default headers values
func NewCreateCopyInternalServerError() *CreateCopyInternalServerError {
	return &CreateCopyInternalServerError{}
}

// WithPayload adds the payload to the create copy internal server error response
func (o *CreateCopyInternalServerError) WithPayload(payload *models.Error) *CreateCopyInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create copy internal server error response
func (o *CreateCopyInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCopyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateCopyDefault Unexpected error

swagger:response createCopyDefault
*/
type CreateCopyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateCopyDefault creates CreateCopyDefault with default headers values
func NewCreateCopyDefault(code int) *CreateCopyDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateCopyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create copy default response
func (o *CreateCopyDefault) WithStatusCode(code int) *CreateCopyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create copy default response
func (o *CreateCopyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create copy default response
func (o *CreateCopyDefault) WithPayload(payload *models.Error) *CreateCopyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create copy default response
func (o *CreateCopyDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCopyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
