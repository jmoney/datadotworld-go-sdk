// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateDatasetResponse create dataset response
//
// swagger:model CreateDatasetResponse
type CreateDatasetResponse struct {

	// message
	// Max Length: 256
	// Min Length: 0
	Message *string `json:"message,omitempty"`

	// uri
	// Required: true
	// Format: uri
	URI *strfmt.URI `json:"uri"`
}

// Validate validates this create dataset response
func (m *CreateDatasetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDatasetResponse) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if err := validate.MinLength("message", "body", *m.Message, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("message", "body", *m.Message, 256); err != nil {
		return err
	}

	return nil
}

func (m *CreateDatasetResponse) validateURI(formats strfmt.Registry) error {

	if err := validate.Required("uri", "body", m.URI); err != nil {
		return err
	}

	if err := validate.FormatOf("uri", "body", "uri", m.URI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create dataset response based on context it is used
func (m *CreateDatasetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateDatasetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDatasetResponse) UnmarshalBinary(b []byte) error {
	var res CreateDatasetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
