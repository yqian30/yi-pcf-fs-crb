package crb_web

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"crb/models"
)

/*GetRepositoryInfoOK Successful operation

swagger:response getRepositoryInfoOK
*/
type GetRepositoryInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.RepositoryInfo `json:"body,omitempty"`
}

// NewGetRepositoryInfoOK creates GetRepositoryInfoOK with default headers values
func NewGetRepositoryInfoOK() *GetRepositoryInfoOK {
	return &GetRepositoryInfoOK{}
}

// WithPayload adds the payload to the get repository info o k response
func (o *GetRepositoryInfoOK) WithPayload(payload *models.RepositoryInfo) *GetRepositoryInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repository info o k response
func (o *GetRepositoryInfoOK) SetPayload(payload *models.RepositoryInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRepositoryInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetRepositoryInfoNotFound No repositories content in Metadata-DB

swagger:response getRepositoryInfoNotFound
*/
type GetRepositoryInfoNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRepositoryInfoNotFound creates GetRepositoryInfoNotFound with default headers values
func NewGetRepositoryInfoNotFound() *GetRepositoryInfoNotFound {
	return &GetRepositoryInfoNotFound{}
}

// WithPayload adds the payload to the get repository info not found response
func (o *GetRepositoryInfoNotFound) WithPayload(payload *models.Error) *GetRepositoryInfoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repository info not found response
func (o *GetRepositoryInfoNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRepositoryInfoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetRepositoryInfoInternalServerError Internal error -- most likely Metadata-DB not accessible

swagger:response getRepositoryInfoInternalServerError
*/
type GetRepositoryInfoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRepositoryInfoInternalServerError creates GetRepositoryInfoInternalServerError with default headers values
func NewGetRepositoryInfoInternalServerError() *GetRepositoryInfoInternalServerError {
	return &GetRepositoryInfoInternalServerError{}
}

// WithPayload adds the payload to the get repository info internal server error response
func (o *GetRepositoryInfoInternalServerError) WithPayload(payload *models.Error) *GetRepositoryInfoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repository info internal server error response
func (o *GetRepositoryInfoInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRepositoryInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetRepositoryInfoDefault Unexpected error

swagger:response getRepositoryInfoDefault
*/
type GetRepositoryInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRepositoryInfoDefault creates GetRepositoryInfoDefault with default headers values
func NewGetRepositoryInfoDefault(code int) *GetRepositoryInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRepositoryInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get repository info default response
func (o *GetRepositoryInfoDefault) WithStatusCode(code int) *GetRepositoryInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get repository info default response
func (o *GetRepositoryInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get repository info default response
func (o *GetRepositoryInfoDefault) WithPayload(payload *models.Error) *GetRepositoryInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repository info default response
func (o *GetRepositoryInfoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRepositoryInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
