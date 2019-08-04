// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/delgus/swagger-rest/internal/models"
)

// GetBooksOKCode is the HTTP code returned for type GetBooksOK
const GetBooksOKCode int = 200

/*GetBooksOK OK

swagger:response getBooksOK
*/
type GetBooksOK struct {

	/*
	  In: Body
	*/
	Payload models.ArrayOfBooks `json:"body,omitempty"`
}

// NewGetBooksOK creates GetBooksOK with default headers values
func NewGetBooksOK() *GetBooksOK {

	return &GetBooksOK{}
}

// WithPayload adds the payload to the get books o k response
func (o *GetBooksOK) WithPayload(payload models.ArrayOfBooks) *GetBooksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get books o k response
func (o *GetBooksOK) SetPayload(payload models.ArrayOfBooks) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBooksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.ArrayOfBooks{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBooksNotFoundCode is the HTTP code returned for type GetBooksNotFound
const GetBooksNotFoundCode int = 404

/*GetBooksNotFound Not Found

swagger:response getBooksNotFound
*/
type GetBooksNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBooksNotFound creates GetBooksNotFound with default headers values
func NewGetBooksNotFound() *GetBooksNotFound {

	return &GetBooksNotFound{}
}

// WithPayload adds the payload to the get books not found response
func (o *GetBooksNotFound) WithPayload(payload *models.Error) *GetBooksNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get books not found response
func (o *GetBooksNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBooksNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
