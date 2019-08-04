// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/delgus/swagger-rest/internal/models"
)

// NewPutBooksParams creates a new PutBooksParams object
// no default values defined in spec.
func NewPutBooksParams() PutBooksParams {

	return PutBooksParams{}
}

// PutBooksParams contains all the bound params for the put books operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutBooks
type PutBooksParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The book to create.
	  In: body
	*/
	Book *models.Book
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutBooksParams() beforehand.
func (o *PutBooksParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Book
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("book", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Book = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
