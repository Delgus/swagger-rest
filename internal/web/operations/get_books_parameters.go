// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBooksParams creates a new GetBooksParams object
// no default values defined in spec.
func NewGetBooksParams() GetBooksParams {

	return GetBooksParams{}
}

// GetBooksParams contains all the bound params for the get books operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetBooks
type GetBooksParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Author for search book
	  In: query
	*/
	Author *string
	/*Title for search book
	  In: query
	*/
	Title *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBooksParams() beforehand.
func (o *GetBooksParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAuthor, qhkAuthor, _ := qs.GetOK("author")
	if err := o.bindAuthor(qAuthor, qhkAuthor, route.Formats); err != nil {
		res = append(res, err)
	}

	qTitle, qhkTitle, _ := qs.GetOK("title")
	if err := o.bindTitle(qTitle, qhkTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthor binds and validates parameter Author from query.
func (o *GetBooksParams) bindAuthor(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Author = &raw

	return nil
}

// bindTitle binds and validates parameter Title from query.
func (o *GetBooksParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Title = &raw

	return nil
}