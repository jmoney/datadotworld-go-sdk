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

// DatasetHydrationDto dataset hydration dto
//
// swagger:model DatasetHydrationDto
type DatasetHydrationDto struct {

	// abstract
	Abstract string `json:"abstract,omitempty"`

	// agentid
	Agentid string `json:"agentid,omitempty"`

	// created by
	CreatedBy *AgentHydrationDto `json:"createdBy,omitempty"`

	// datasetid
	Datasetid string `json:"datasetid,omitempty"`

	// headline
	Headline string `json:"headline,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project
	Project *bool `json:"project,omitempty"`

	// referent
	Referent string `json:"referent,omitempty"`

	// relationship
	Relationship *ResourceRelationshipDto `json:"relationship,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// user updated
	// Format: date-time
	UserUpdated strfmt.DateTime `json:"userUpdated,omitempty"`
}

// Validate validates this dataset hydration dto
func (m *DatasetHydrationDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasetHydrationDto) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *DatasetHydrationDto) validateRelationship(formats strfmt.Registry) error {
	if swag.IsZero(m.Relationship) { // not required
		return nil
	}

	if m.Relationship != nil {
		if err := m.Relationship.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationship")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relationship")
			}
			return err
		}
	}

	return nil
}

func (m *DatasetHydrationDto) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DatasetHydrationDto) validateUserUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.UserUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("userUpdated", "body", "date-time", m.UserUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dataset hydration dto based on the context it is used
func (m *DatasetHydrationDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelationship(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasetHydrationDto) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {
		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *DatasetHydrationDto) contextValidateRelationship(ctx context.Context, formats strfmt.Registry) error {

	if m.Relationship != nil {
		if err := m.Relationship.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationship")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relationship")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatasetHydrationDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasetHydrationDto) UnmarshalBinary(b []byte) error {
	var res DatasetHydrationDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
