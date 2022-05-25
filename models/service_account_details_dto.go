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

// ServiceAccountDetailsDto service account details dto
//
// swagger:model ServiceAccountDetailsDto
type ServiceAccountDetailsDto struct {

	// active
	Active *bool `json:"active,omitempty"`

	// desired username
	DesiredUsername string `json:"desiredUsername,omitempty"`

	// expiry date
	// Format: date-time
	ExpiryDate strfmt.DateTime `json:"expiryDate,omitempty"`
}

// Validate validates this service account details dto
func (m *ServiceAccountDetailsDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiryDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceAccountDetailsDto) validateExpiryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expiryDate", "body", "date-time", m.ExpiryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service account details dto based on context it is used
func (m *ServiceAccountDetailsDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceAccountDetailsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceAccountDetailsDto) UnmarshalBinary(b []byte) error {
	var res ServiceAccountDetailsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
