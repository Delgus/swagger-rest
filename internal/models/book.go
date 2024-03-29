// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Book book
// swagger:model Book
type Book struct {

	// The author name of the book
	// Required: true
	Author *string `json:"author"`

	// ID of the book
	// Required: true
	// Minimum: 1
	ID *int64 `json:"id"`

	// The book title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this book
func (m *Book) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Book) validateAuthor(formats strfmt.Registry) error {

	if err := validate.Required("author", "body", m.Author); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Book) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Book) UnmarshalBinary(b []byte) error {
	var res Book
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
