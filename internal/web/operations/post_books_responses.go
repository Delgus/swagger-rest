// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostBooksCreatedCode is the HTTP code returned for type PostBooksCreated
const PostBooksCreatedCode int = 201

/*PostBooksCreated Created

swagger:response postBooksCreated
*/
type PostBooksCreated struct {
}

// NewPostBooksCreated creates PostBooksCreated with default headers values
func NewPostBooksCreated() *PostBooksCreated {

	return &PostBooksCreated{}
}

// WriteResponse to the client
func (o *PostBooksCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostBooksInternalServerErrorCode is the HTTP code returned for type PostBooksInternalServerError
const PostBooksInternalServerErrorCode int = 500

/*PostBooksInternalServerError Internal Server Error

swagger:response postBooksInternalServerError
*/
type PostBooksInternalServerError struct {
}

// NewPostBooksInternalServerError creates PostBooksInternalServerError with default headers values
func NewPostBooksInternalServerError() *PostBooksInternalServerError {

	return &PostBooksInternalServerError{}
}

// WriteResponse to the client
func (o *PostBooksInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
