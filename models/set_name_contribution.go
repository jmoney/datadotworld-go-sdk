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

// SetNameContribution set name contribution
//
// swagger:model SetNameContribution
type SetNameContribution struct {
	contributionHydratablesField []*ContributionHydratable

	// entity type
	// Required: true
	EntityType *string `json:"entityType"`

	// name
	Name string `json:"name,omitempty"`

	// override type
	// Enum: [SET ADD REMOVE]
	OverrideType string `json:"overrideType,omitempty"`

	// target
	// Required: true
	Target *string `json:"target"`
}

// ContributionHydratables gets the contribution hydratables of this subtype
func (m *SetNameContribution) ContributionHydratables() []*ContributionHydratable {
	return m.contributionHydratablesField
}

// SetContributionHydratables sets the contribution hydratables of this subtype
func (m *SetNameContribution) SetContributionHydratables(val []*ContributionHydratable) {
	m.contributionHydratablesField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SetNameContribution) UnmarshalJSON(raw []byte) error {
	var data struct {

		// entity type
		// Required: true
		EntityType *string `json:"entityType"`

		// name
		Name string `json:"name,omitempty"`

		// override type
		// Enum: [SET ADD REMOVE]
		OverrideType string `json:"overrideType,omitempty"`

		// target
		// Required: true
		Target *string `json:"target"`
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

	var result SetNameContribution

	result.contributionHydratablesField = base.ContributionHydratables

	result.EntityType = data.EntityType
	result.Name = data.Name
	result.OverrideType = data.OverrideType
	result.Target = data.Target

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SetNameContribution) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// entity type
		// Required: true
		EntityType *string `json:"entityType"`

		// name
		Name string `json:"name,omitempty"`

		// override type
		// Enum: [SET ADD REMOVE]
		OverrideType string `json:"overrideType,omitempty"`

		// target
		// Required: true
		Target *string `json:"target"`
	}{

		EntityType: m.EntityType,

		Name: m.Name,

		OverrideType: m.OverrideType,

		Target: m.Target,
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

// Validate validates this set name contribution
func (m *SetNameContribution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContributionHydratables(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverrideType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetNameContribution) validateContributionHydratables(formats strfmt.Registry) error {

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

func (m *SetNameContribution) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entityType", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

var setNameContributionTypeOverrideTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SET","ADD","REMOVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		setNameContributionTypeOverrideTypePropEnum = append(setNameContributionTypeOverrideTypePropEnum, v)
	}
}

// property enum
func (m *SetNameContribution) validateOverrideTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, setNameContributionTypeOverrideTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SetNameContribution) validateOverrideType(formats strfmt.Registry) error {

	if swag.IsZero(m.OverrideType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOverrideTypeEnum("overrideType", "body", m.OverrideType); err != nil {
		return err
	}

	return nil
}

func (m *SetNameContribution) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this set name contribution based on the context it is used
func (m *SetNameContribution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContributionHydratables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetNameContribution) contextValidateContributionHydratables(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SetNameContribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetNameContribution) UnmarshalBinary(b []byte) error {
	var res SetNameContribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
