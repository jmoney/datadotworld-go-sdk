// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProjectPutRequest project put request
//
// swagger:model ProjectPutRequest
type ProjectPutRequest struct {

	// license
	// Enum: [CC BY-NC CC BY-NC-ND CC BY-NC-SA CC BY-ND CC BY-ND 3.0 CC-0 CC-BY CC-BY 3.0 CC-BY 3.0 AU CC-BY 3.0 IGO CC-BY 3.0 NZ CC-BY-IGO CC-BY-NC 3.0 CC-BY-NC 3.0 NZ CC-BY-NC-ND-NZ-3.0 CC-BY-NC-SA 3.0 CC-BY-NC-SA 3.0 NZ CC-BY-SA CC-BY-SA 3.0 CC-BY-SA 3.0 NZ CDLA-Permissive-1.0 CDLA-Sharing-1.0 Italian-ODL MIT License ODC-BY ODC-ODbL OGL OGL-Canada OGL-Nova Scotia OGL-UK OSODL Other PDDL Public Domain]
	License string `json:"license,omitempty"`

	// linked datasets
	LinkedDatasets []*LinkedDatasetCreateOrUpdateRequest `json:"linkedDatasets"`

	// objective
	// Max Length: 120
	// Min Length: 0
	Objective *string `json:"objective,omitempty"`

	// properties
	Properties map[string]JSONNode `json:"properties,omitempty"`

	// summary
	// Max Length: 25000
	// Min Length: 0
	Summary *string `json:"summary,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// title
	// Required: true
	// Max Length: 60
	// Min Length: 1
	Title *string `json:"title"`

	// visibility
	// Required: true
	// Enum: [OPEN PRIVATE]
	Visibility *string `json:"visibility"`
}

// Validate validates this project put request
func (m *ProjectPutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLicense(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkedDatasets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjective(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var projectPutRequestTypeLicensePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CC BY-NC","CC BY-NC-ND","CC BY-NC-SA","CC BY-ND","CC BY-ND 3.0","CC-0","CC-BY","CC-BY 3.0","CC-BY 3.0 AU","CC-BY 3.0 IGO","CC-BY 3.0 NZ","CC-BY-IGO","CC-BY-NC 3.0","CC-BY-NC 3.0 NZ","CC-BY-NC-ND-NZ-3.0","CC-BY-NC-SA 3.0","CC-BY-NC-SA 3.0 NZ","CC-BY-SA","CC-BY-SA 3.0","CC-BY-SA 3.0 NZ","CDLA-Permissive-1.0","CDLA-Sharing-1.0","Italian-ODL","MIT License","ODC-BY","ODC-ODbL","OGL","OGL-Canada","OGL-Nova Scotia","OGL-UK","OSODL","Other","PDDL","Public Domain"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectPutRequestTypeLicensePropEnum = append(projectPutRequestTypeLicensePropEnum, v)
	}
}

const (

	// ProjectPutRequestLicenseCCBYDashNC captures enum value "CC BY-NC"
	ProjectPutRequestLicenseCCBYDashNC string = "CC BY-NC"

	// ProjectPutRequestLicenseCCBYDashNCDashND captures enum value "CC BY-NC-ND"
	ProjectPutRequestLicenseCCBYDashNCDashND string = "CC BY-NC-ND"

	// ProjectPutRequestLicenseCCBYDashNCDashSA captures enum value "CC BY-NC-SA"
	ProjectPutRequestLicenseCCBYDashNCDashSA string = "CC BY-NC-SA"

	// ProjectPutRequestLicenseCCBYDashND captures enum value "CC BY-ND"
	ProjectPutRequestLicenseCCBYDashND string = "CC BY-ND"

	// ProjectPutRequestLicenseCCBYDashND3Dot0 captures enum value "CC BY-ND 3.0"
	ProjectPutRequestLicenseCCBYDashND3Dot0 string = "CC BY-ND 3.0"

	// ProjectPutRequestLicenseCCDash0 captures enum value "CC-0"
	ProjectPutRequestLicenseCCDash0 string = "CC-0"

	// ProjectPutRequestLicenseCCDashBY captures enum value "CC-BY"
	ProjectPutRequestLicenseCCDashBY string = "CC-BY"

	// ProjectPutRequestLicenseCCDashBY3Dot0 captures enum value "CC-BY 3.0"
	ProjectPutRequestLicenseCCDashBY3Dot0 string = "CC-BY 3.0"

	// ProjectPutRequestLicenseCCDashBY3Dot0AU captures enum value "CC-BY 3.0 AU"
	ProjectPutRequestLicenseCCDashBY3Dot0AU string = "CC-BY 3.0 AU"

	// ProjectPutRequestLicenseCCDashBY3Dot0IGO captures enum value "CC-BY 3.0 IGO"
	ProjectPutRequestLicenseCCDashBY3Dot0IGO string = "CC-BY 3.0 IGO"

	// ProjectPutRequestLicenseCCDashBY3Dot0NZ captures enum value "CC-BY 3.0 NZ"
	ProjectPutRequestLicenseCCDashBY3Dot0NZ string = "CC-BY 3.0 NZ"

	// ProjectPutRequestLicenseCCDashBYDashIGO captures enum value "CC-BY-IGO"
	ProjectPutRequestLicenseCCDashBYDashIGO string = "CC-BY-IGO"

	// ProjectPutRequestLicenseCCDashBYDashNC3Dot0 captures enum value "CC-BY-NC 3.0"
	ProjectPutRequestLicenseCCDashBYDashNC3Dot0 string = "CC-BY-NC 3.0"

	// ProjectPutRequestLicenseCCDashBYDashNC3Dot0NZ captures enum value "CC-BY-NC 3.0 NZ"
	ProjectPutRequestLicenseCCDashBYDashNC3Dot0NZ string = "CC-BY-NC 3.0 NZ"

	// ProjectPutRequestLicenseCCDashBYDashNCDashNDDashNZDash3Dot0 captures enum value "CC-BY-NC-ND-NZ-3.0"
	ProjectPutRequestLicenseCCDashBYDashNCDashNDDashNZDash3Dot0 string = "CC-BY-NC-ND-NZ-3.0"

	// ProjectPutRequestLicenseCCDashBYDashNCDashSA3Dot0 captures enum value "CC-BY-NC-SA 3.0"
	ProjectPutRequestLicenseCCDashBYDashNCDashSA3Dot0 string = "CC-BY-NC-SA 3.0"

	// ProjectPutRequestLicenseCCDashBYDashNCDashSA3Dot0NZ captures enum value "CC-BY-NC-SA 3.0 NZ"
	ProjectPutRequestLicenseCCDashBYDashNCDashSA3Dot0NZ string = "CC-BY-NC-SA 3.0 NZ"

	// ProjectPutRequestLicenseCCDashBYDashSA captures enum value "CC-BY-SA"
	ProjectPutRequestLicenseCCDashBYDashSA string = "CC-BY-SA"

	// ProjectPutRequestLicenseCCDashBYDashSA3Dot0 captures enum value "CC-BY-SA 3.0"
	ProjectPutRequestLicenseCCDashBYDashSA3Dot0 string = "CC-BY-SA 3.0"

	// ProjectPutRequestLicenseCCDashBYDashSA3Dot0NZ captures enum value "CC-BY-SA 3.0 NZ"
	ProjectPutRequestLicenseCCDashBYDashSA3Dot0NZ string = "CC-BY-SA 3.0 NZ"

	// ProjectPutRequestLicenseCDLADashPermissiveDash1Dot0 captures enum value "CDLA-Permissive-1.0"
	ProjectPutRequestLicenseCDLADashPermissiveDash1Dot0 string = "CDLA-Permissive-1.0"

	// ProjectPutRequestLicenseCDLADashSharingDash1Dot0 captures enum value "CDLA-Sharing-1.0"
	ProjectPutRequestLicenseCDLADashSharingDash1Dot0 string = "CDLA-Sharing-1.0"

	// ProjectPutRequestLicenseItalianDashODL captures enum value "Italian-ODL"
	ProjectPutRequestLicenseItalianDashODL string = "Italian-ODL"

	// ProjectPutRequestLicenseMITLicense captures enum value "MIT License"
	ProjectPutRequestLicenseMITLicense string = "MIT License"

	// ProjectPutRequestLicenseODCDashBY captures enum value "ODC-BY"
	ProjectPutRequestLicenseODCDashBY string = "ODC-BY"

	// ProjectPutRequestLicenseODCDashODbL captures enum value "ODC-ODbL"
	ProjectPutRequestLicenseODCDashODbL string = "ODC-ODbL"

	// ProjectPutRequestLicenseOGL captures enum value "OGL"
	ProjectPutRequestLicenseOGL string = "OGL"

	// ProjectPutRequestLicenseOGLDashCanada captures enum value "OGL-Canada"
	ProjectPutRequestLicenseOGLDashCanada string = "OGL-Canada"

	// ProjectPutRequestLicenseOGLDashNovaScotia captures enum value "OGL-Nova Scotia"
	ProjectPutRequestLicenseOGLDashNovaScotia string = "OGL-Nova Scotia"

	// ProjectPutRequestLicenseOGLDashUK captures enum value "OGL-UK"
	ProjectPutRequestLicenseOGLDashUK string = "OGL-UK"

	// ProjectPutRequestLicenseOSODL captures enum value "OSODL"
	ProjectPutRequestLicenseOSODL string = "OSODL"

	// ProjectPutRequestLicenseOther captures enum value "Other"
	ProjectPutRequestLicenseOther string = "Other"

	// ProjectPutRequestLicensePDDL captures enum value "PDDL"
	ProjectPutRequestLicensePDDL string = "PDDL"

	// ProjectPutRequestLicensePublicDomain captures enum value "Public Domain"
	ProjectPutRequestLicensePublicDomain string = "Public Domain"
)

// prop value enum
func (m *ProjectPutRequest) validateLicenseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, projectPutRequestTypeLicensePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProjectPutRequest) validateLicense(formats strfmt.Registry) error {
	if swag.IsZero(m.License) { // not required
		return nil
	}

	// value enum
	if err := m.validateLicenseEnum("license", "body", m.License); err != nil {
		return err
	}

	return nil
}

func (m *ProjectPutRequest) validateLinkedDatasets(formats strfmt.Registry) error {
	if swag.IsZero(m.LinkedDatasets) { // not required
		return nil
	}

	for i := 0; i < len(m.LinkedDatasets); i++ {
		if swag.IsZero(m.LinkedDatasets[i]) { // not required
			continue
		}

		if m.LinkedDatasets[i] != nil {
			if err := m.LinkedDatasets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("linkedDatasets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("linkedDatasets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectPutRequest) validateObjective(formats strfmt.Registry) error {
	if swag.IsZero(m.Objective) { // not required
		return nil
	}

	if err := validate.MinLength("objective", "body", *m.Objective, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("objective", "body", *m.Objective, 120); err != nil {
		return err
	}

	return nil
}

func (m *ProjectPutRequest) validateProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	for k := range m.Properties {

		if err := validate.Required("properties"+"."+k, "body", m.Properties[k]); err != nil {
			return err
		}
		if val, ok := m.Properties[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("properties" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectPutRequest) validateSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.Summary) { // not required
		return nil
	}

	if err := validate.MinLength("summary", "body", *m.Summary, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("summary", "body", *m.Summary, 25000); err != nil {
		return err
	}

	return nil
}

func (m *ProjectPutRequest) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 60); err != nil {
		return err
	}

	return nil
}

var projectPutRequestTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectPutRequestTypeVisibilityPropEnum = append(projectPutRequestTypeVisibilityPropEnum, v)
	}
}

const (

	// ProjectPutRequestVisibilityOPEN captures enum value "OPEN"
	ProjectPutRequestVisibilityOPEN string = "OPEN"

	// ProjectPutRequestVisibilityPRIVATE captures enum value "PRIVATE"
	ProjectPutRequestVisibilityPRIVATE string = "PRIVATE"
)

// prop value enum
func (m *ProjectPutRequest) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, projectPutRequestTypeVisibilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProjectPutRequest) validateVisibility(formats strfmt.Registry) error {

	if err := validate.Required("visibility", "body", m.Visibility); err != nil {
		return err
	}

	// value enum
	if err := m.validateVisibilityEnum("visibility", "body", *m.Visibility); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this project put request based on the context it is used
func (m *ProjectPutRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinkedDatasets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectPutRequest) contextValidateLinkedDatasets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LinkedDatasets); i++ {

		if m.LinkedDatasets[i] != nil {
			if err := m.LinkedDatasets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("linkedDatasets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("linkedDatasets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectPutRequest) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Properties {

		if val, ok := m.Properties[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectPutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectPutRequest) UnmarshalBinary(b []byte) error {
	var res ProjectPutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}