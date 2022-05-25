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

// LinkedDatasetCreateOrUpdateRequest linked dataset create or update request
//
// swagger:model LinkedDatasetCreateOrUpdateRequest
type LinkedDatasetCreateOrUpdateRequest struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// owner
	// Required: true
	Owner *string `json:"owner"`
}

// Validate validates this linked dataset create or update request
func (m *LinkedDatasetCreateOrUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkedDatasetCreateOrUpdateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *LinkedDatasetCreateOrUpdateRequest) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this linked dataset create or update request based on context it is used
func (m *LinkedDatasetCreateOrUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinkedDatasetCreateOrUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkedDatasetCreateOrUpdateRequest) UnmarshalBinary(b []byte) error {
	var res LinkedDatasetCreateOrUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
