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

// FileSummaryResponse file summary response
//
// swagger:model FileSummaryResponse
type FileSummaryResponse struct {

	// created
	// Required: true
	Created *string `json:"created"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// description
	// Max Length: 120
	// Min Length: 1
	Description string `json:"description,omitempty"`

	// labels
	Labels []string `json:"labels"`

	// name
	// Required: true
	// Max Length: 128
	// Min Length: 1
	// Pattern: ^[^/]+$
	Name *string `json:"name"`

	// size in bytes
	SizeInBytes int64 `json:"sizeInBytes,omitempty"`

	// source
	Source *FileSourceSummaryResponse `json:"source,omitempty"`

	// updated
	// Required: true
	Updated *string `json:"updated"`

	// updated by
	UpdatedBy string `json:"updatedBy,omitempty"`
}

// Validate validates this file summary response
func (m *FileSummaryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileSummaryResponse) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *FileSummaryResponse) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", m.Description, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", m.Description, 120); err != nil {
		return err
	}

	return nil
}

func (m *FileSummaryResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 128); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[^/]+$`); err != nil {
		return err
	}

	return nil
}

func (m *FileSummaryResponse) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *FileSummaryResponse) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("updated", "body", m.Updated); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this file summary response based on the context it is used
func (m *FileSummaryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileSummaryResponse) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileSummaryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileSummaryResponse) UnmarshalBinary(b []byte) error {
	var res FileSummaryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
