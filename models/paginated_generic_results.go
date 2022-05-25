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

// PaginatedGenericResults paginated generic results
//
// swagger:model PaginatedGenericResults
type PaginatedGenericResults struct {

	// count
	// Minimum: 0
	Count *int32 `json:"count,omitempty"`

	// hydrations
	Hydrations *SearchHydrations `json:"hydrations,omitempty"`

	// next
	Next string `json:"next,omitempty"`

	// records
	Records []interface{} `json:"records"`
}

// Validate validates this paginated generic results
func (m *PaginatedGenericResults) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHydrations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedGenericResults) validateCount(formats strfmt.Registry) error {
	if swag.IsZero(m.Count) { // not required
		return nil
	}

	if err := validate.MinimumInt("count", "body", int64(*m.Count), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PaginatedGenericResults) validateHydrations(formats strfmt.Registry) error {
	if swag.IsZero(m.Hydrations) { // not required
		return nil
	}

	if m.Hydrations != nil {
		if err := m.Hydrations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hydrations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hydrations")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this paginated generic results based on the context it is used
func (m *PaginatedGenericResults) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHydrations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedGenericResults) contextValidateHydrations(ctx context.Context, formats strfmt.Registry) error {

	if m.Hydrations != nil {
		if err := m.Hydrations.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hydrations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hydrations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaginatedGenericResults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginatedGenericResults) UnmarshalBinary(b []byte) error {
	var res PaginatedGenericResults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
