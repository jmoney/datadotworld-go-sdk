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

// AddUsesDataFromContribution add uses data from contribution
//
// swagger:model AddUsesDataFromContribution
type AddUsesDataFromContribution struct {
	contributionHydratablesField []*ContributionHydratable

	// analysis iri
	// Required: true
	AnalysisIri *string `json:"analysisIri"`

	// override type
	// Enum: [SET ADD REMOVE]
	OverrideType string `json:"overrideType,omitempty"`

	// table iri
	// Required: true
	TableIri *string `json:"tableIri"`
}

// ContributionHydratables gets the contribution hydratables of this subtype
func (m *AddUsesDataFromContribution) ContributionHydratables() []*ContributionHydratable {
	return m.contributionHydratablesField
}

// SetContributionHydratables sets the contribution hydratables of this subtype
func (m *AddUsesDataFromContribution) SetContributionHydratables(val []*ContributionHydratable) {
	m.contributionHydratablesField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AddUsesDataFromContribution) UnmarshalJSON(raw []byte) error {
	var data struct {

		// analysis iri
		// Required: true
		AnalysisIri *string `json:"analysisIri"`

		// override type
		// Enum: [SET ADD REMOVE]
		OverrideType string `json:"overrideType,omitempty"`

		// table iri
		// Required: true
		TableIri *string `json:"tableIri"`
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

	var result AddUsesDataFromContribution

	result.contributionHydratablesField = base.ContributionHydratables

	result.AnalysisIri = data.AnalysisIri
	result.OverrideType = data.OverrideType
	result.TableIri = data.TableIri

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AddUsesDataFromContribution) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// analysis iri
		// Required: true
		AnalysisIri *string `json:"analysisIri"`

		// override type
		// Enum: [SET ADD REMOVE]
		OverrideType string `json:"overrideType,omitempty"`

		// table iri
		// Required: true
		TableIri *string `json:"tableIri"`
	}{

		AnalysisIri: m.AnalysisIri,

		OverrideType: m.OverrideType,

		TableIri: m.TableIri,
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

// Validate validates this add uses data from contribution
func (m *AddUsesDataFromContribution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContributionHydratables(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnalysisIri(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverrideType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTableIri(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddUsesDataFromContribution) validateContributionHydratables(formats strfmt.Registry) error {

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

func (m *AddUsesDataFromContribution) validateAnalysisIri(formats strfmt.Registry) error {

	if err := validate.Required("analysisIri", "body", m.AnalysisIri); err != nil {
		return err
	}

	return nil
}

var addUsesDataFromContributionTypeOverrideTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SET","ADD","REMOVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addUsesDataFromContributionTypeOverrideTypePropEnum = append(addUsesDataFromContributionTypeOverrideTypePropEnum, v)
	}
}

// property enum
func (m *AddUsesDataFromContribution) validateOverrideTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addUsesDataFromContributionTypeOverrideTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AddUsesDataFromContribution) validateOverrideType(formats strfmt.Registry) error {

	if swag.IsZero(m.OverrideType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOverrideTypeEnum("overrideType", "body", m.OverrideType); err != nil {
		return err
	}

	return nil
}

func (m *AddUsesDataFromContribution) validateTableIri(formats strfmt.Registry) error {

	if err := validate.Required("tableIri", "body", m.TableIri); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this add uses data from contribution based on the context it is used
func (m *AddUsesDataFromContribution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContributionHydratables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddUsesDataFromContribution) contextValidateContributionHydratables(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AddUsesDataFromContribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddUsesDataFromContribution) UnmarshalBinary(b []byte) error {
	var res AddUsesDataFromContribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
