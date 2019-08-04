// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteBooksBookIDParams creates a new DeleteBooksBookIDParams object
// no default values defined in spec.
func NewDeleteBooksBookIDParams() DeleteBooksBookIDParams {

	return DeleteBooksBookIDParams{}
}

// DeleteBooksBookIDParams contains all the bound params for the delete books book ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteBooksBookID
type DeleteBooksBookIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Numeric ID of the book to get
	  Required: true
	  Minimum: 1
	  In: path
	*/
	BookID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteBooksBookIDParams() beforehand.
func (o *DeleteBooksBookIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBookID, rhkBookID, _ := route.Params.GetOK("bookID")
	if err := o.bindBookID(rBookID, rhkBookID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBookID binds and validates parameter BookID from path.
func (o *DeleteBooksBookIDParams) bindBookID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("bookID", "path", "int64", raw)
	}
	o.BookID = value

	if err := o.validateBookID(formats); err != nil {
		return err
	}

	return nil
}

// validateBookID carries on validations for parameter BookID
func (o *DeleteBooksBookIDParams) validateBookID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("bookID", "path", int64(o.BookID), 1, false); err != nil {
		return err
	}

	return nil
}
