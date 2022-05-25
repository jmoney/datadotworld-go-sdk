// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateExternalResourceContribution create external resource contribution
//
// swagger:model CreateExternalResourceContribution
type CreateExternalResourceContribution struct {
	contributionHydratablesField []*ContributionHydratable

	// entity iri
	EntityIri string `json:"entityIri,omitempty"`

	// entity type
	// Required: true
	EntityType *string `json:"entityType"`

	// external Url
	// Required: true
	ExternalURL *string `json:"externalUrl"`

	// name
	Name string `json:"name,omitempty"`
}

// ContributionHydratables gets the contribution hydratables of this subtype
func (m *CreateExternalResourceContribution) ContributionHydratables() []*ContributionHydratable {
	return m.contributionHydratablesField
}

// SetContributionHydratables sets the contribution hydratables of this subtype
func (m *CreateExternalResourceContribution) SetContributionHydratables(val []*ContributionHydratable) {
	m.contributionHydratablesField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CreateExternalResourceContribution) UnmarshalJSON(raw []byte) error {
	var data struct {

		// entity iri
		EntityIri string `json:"entityIri,omitempty"`

		// entity type
		// Required: true
		EntityType *string `json:"entityType"`

		// external Url
		// Required: true
		ExternalURL *string `json:"externalUrl"`

		// name
		Name string `json:"name,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ContributionHydratables []*ContributionHydratable `json:"contributionHydratables"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result CreateExternalResourceContribution

	result.contributionHydratablesField = base.ContributionHydratables

	result.EntityIri = data.EntityIri
	result.EntityType = data.EntityType
	result.ExternalURL = data.ExternalURL
	result.Name = data.Name

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CreateExternalResourceContribution) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// entity iri
		EntityIri string `json:"entityIri,omitempty"`

		// entity type
		// Required: true
		EntityType *string `json:"entityType"`

		// external Url
		// Required: true
		ExternalURL *string `json:"externalUrl"`

		// name
		Name string `json:"name,omitempty"`
	}{

		EntityIri: m.EntityIri,

		EntityType: m.EntityType,

		ExternalURL: m.ExternalURL,

		Name: m.Name,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ContributionHydratables []*ContributionHydratable `json:"contributionHydratables"`
	}{

		ContributionHydratables: m.ContributionHydratables(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this create external resource contribution
func (m *CreateExternalResourceContribution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContributionHydratables(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateExternalResourceContribution) validateContributionHydratables(formats strfmt.Registry) error {

	if swag.IsZero(m.ContributionHydratables()) { // not required
		return nil
	}

	for i := 0; i < len(m.ContributionHydratables()); i++ {
		if swag.IsZero(m.contributionHydratablesField[i]) { // not required
			continue
		}

		if m.contributionHydratablesField[i] != nil {
			if err := m.contributionHydratablesField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contributionHydratables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contributionHydratables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateExternalResourceContribution) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entityType", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *CreateExternalResourceContribution) validateExternalURL(formats strfmt.Registry) error {

	if err := validate.Required("externalUrl", "body", m.ExternalURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create external resource contribution based on the context it is used
func (m *CreateExternalResourceContribution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContributionHydratables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateExternalResourceContribution) contextValidateContributionHydratables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ContributionHydratables()); i++ {

		if m.contributionHydratablesField[i] != nil {
			if err := m.contributionHydratablesField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contributionHydratables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contributionHydratables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateExternalResourceContribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateExternalResourceContribution) UnmarshalBinary(b []byte) error {
	var res CreateExternalResourceContribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}