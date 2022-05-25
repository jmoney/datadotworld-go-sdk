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

// WebCredentials web credentials
//
// swagger:model WebCredentials
type WebCredentials struct {

	// password
	// Max Length: 1024
	// Min Length: 0
	Password *string `json:"password,omitempty"`

	// user
	// Required: true
	// Max Length: 1024
	// Min Length: 0
	User *string `json:"user"`
}

// Validate validates this web credentials
func (m *WebCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebCredentials) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("password", "body", *m.Password, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", *m.Password, 1024); err != nil {
		return err
	}

	return nil
}

func (m *WebCredentials) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if err := validate.MinLength("user", "body", *m.User, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("user", "body", *m.User, 1024); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this web credentials based on context it is used
func (m *WebCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebCredentials) UnmarshalBinary(b []byte) error {
	var res WebCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}