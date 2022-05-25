// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CatalogHydrationDto catalog hydration dto
//
// swagger:model CatalogHydrationDto
type CatalogHydrationDto struct {

	// agentid
	Agentid string `json:"agentid,omitempty"`

	// catalogid
	Catalogid string `json:"catalogid,omitempty"`

	// created by
	CreatedBy *AgentHydrationDto `json:"createdBy,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// entry type hierarchy
	EntryTypeHierarchy []*EntryType `json:"entryTypeHierarchy"`

	// entry type label
	EntryTypeLabel string `json:"entryTypeLabel,omitempty"`

	// sources
	Sources []*SourceID `json:"sources"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this catalog hydration dto
func (m *CatalogHydrationDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntryTypeHierarchy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSources(formats); err != nil {
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

func (m *CatalogHydrationDto) validateCreatedBy(formats strfmt.Registry) error {
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

func (m *CatalogHydrationDto) validateEntryTypeHierarchy(formats strfmt.Registry) error {
	if swag.IsZero(m.EntryTypeHierarchy) { // not required
		return nil
	}

	for i := 0; i < len(m.EntryTypeHierarchy); i++ {
		if swag.IsZero(m.EntryTypeHierarchy[i]) { // not required
			continue
		}

		if m.EntryTypeHierarchy[i] != nil {
			if err := m.EntryTypeHierarchy[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entryTypeHierarchy" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entryTypeHierarchy" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogHydrationDto) validateSources(formats strfmt.Registry) error {
	if swag.IsZero(m.Sources) { // not required
		return nil
	}

	for i := 0; i < len(m.Sources); i++ {
		if swag.IsZero(m.Sources[i]) { // not required
			continue
		}

		if m.Sources[i] != nil {
			if err := m.Sources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogHydrationDto) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this catalog hydration dto based on the context it is used
func (m *CatalogHydrationDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntryTypeHierarchy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogHydrationDto) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CatalogHydrationDto) contextValidateEntryTypeHierarchy(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntryTypeHierarchy); i++ {

		if m.EntryTypeHierarchy[i] != nil {
			if err := m.EntryTypeHierarchy[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entryTypeHierarchy" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entryTypeHierarchy" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogHydrationDto) contextValidateSources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sources); i++ {

		if m.Sources[i] != nil {
			if err := m.Sources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogHydrationDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogHydrationDto) UnmarshalBinary(b []byte) error {
	var res CatalogHydrationDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
