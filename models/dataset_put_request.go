// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DatasetPutRequest dataset put request
//
// swagger:model DatasetPutRequest
type DatasetPutRequest struct {

	// asset status iri
	AssetStatusIri string `json:"assetStatusIri,omitempty"`

	// description
	// Max Length: 120
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// license
	// Enum: [CC BY-NC CC BY-NC-ND CC BY-NC-SA CC BY-ND CC BY-ND 3.0 CC-0 CC-BY CC-BY 3.0 CC-BY 3.0 AU CC-BY 3.0 IGO CC-BY 3.0 NZ CC-BY-IGO CC-BY-NC 3.0 CC-BY-NC 3.0 NZ CC-BY-NC-ND-NZ-3.0 CC-BY-NC-SA 3.0 CC-BY-NC-SA 3.0 NZ CC-BY-SA CC-BY-SA 3.0 CC-BY-SA 3.0 NZ CDLA-Permissive-1.0 CDLA-Sharing-1.0 Italian-ODL MIT License ODC-BY ODC-ODbL OGL OGL-Canada OGL-Nova Scotia OGL-UK OSODL Other PDDL Public Domain]
	License string `json:"license,omitempty"`

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

// Validate validates this dataset put request
func (m *DatasetPutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicense(formats); err != nil {
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

func (m *DatasetPutRequest) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", *m.Description, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 120); err != nil {
		return err
	}

	return nil
}

var datasetPutRequestTypeLicensePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CC BY-NC","CC BY-NC-ND","CC BY-NC-SA","CC BY-ND","CC BY-ND 3.0","CC-0","CC-BY","CC-BY 3.0","CC-BY 3.0 AU","CC-BY 3.0 IGO","CC-BY 3.0 NZ","CC-BY-IGO","CC-BY-NC 3.0","CC-BY-NC 3.0 NZ","CC-BY-NC-ND-NZ-3.0","CC-BY-NC-SA 3.0","CC-BY-NC-SA 3.0 NZ","CC-BY-SA","CC-BY-SA 3.0","CC-BY-SA 3.0 NZ","CDLA-Permissive-1.0","CDLA-Sharing-1.0","Italian-ODL","MIT License","ODC-BY","ODC-ODbL","OGL","OGL-Canada","OGL-Nova Scotia","OGL-UK","OSODL","Other","PDDL","Public Domain"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datasetPutRequestTypeLicensePropEnum = append(datasetPutRequestTypeLicensePropEnum, v)
	}
}

const (

	// DatasetPutRequestLicenseCCBYDashNC captures enum value "CC BY-NC"
	DatasetPutRequestLicenseCCBYDashNC string = "CC BY-NC"

	// DatasetPutRequestLicenseCCBYDashNCDashND captures enum value "CC BY-NC-ND"
	DatasetPutRequestLicenseCCBYDashNCDashND string = "CC BY-NC-ND"

	// DatasetPutRequestLicenseCCBYDashNCDashSA captures enum value "CC BY-NC-SA"
	DatasetPutRequestLicenseCCBYDashNCDashSA string = "CC BY-NC-SA"

	// DatasetPutRequestLicenseCCBYDashND captures enum value "CC BY-ND"
	DatasetPutRequestLicenseCCBYDashND string = "CC BY-ND"

	// DatasetPutRequestLicenseCCBYDashND3Dot0 captures enum value "CC BY-ND 3.0"
	DatasetPutRequestLicenseCCBYDashND3Dot0 string = "CC BY-ND 3.0"

	// DatasetPutRequestLicenseCCDash0 captures enum value "CC-0"
	DatasetPutRequestLicenseCCDash0 string = "CC-0"

	// DatasetPutRequestLicenseCCDashBY captures enum value "CC-BY"
	DatasetPutRequestLicenseCCDashBY string = "CC-BY"

	// DatasetPutRequestLicenseCCDashBY3Dot0 captures enum value "CC-BY 3.0"
	DatasetPutRequestLicenseCCDashBY3Dot0 string = "CC-BY 3.0"

	// DatasetPutRequestLicenseCCDashBY3Dot0AU captures enum value "CC-BY 3.0 AU"
	DatasetPutRequestLicenseCCDashBY3Dot0AU string = "CC-BY 3.0 AU"

	// DatasetPutRequestLicenseCCDashBY3Dot0IGO captures enum value "CC-BY 3.0 IGO"
	DatasetPutRequestLicenseCCDashBY3Dot0IGO string = "CC-BY 3.0 IGO"

	// DatasetPutRequestLicenseCCDashBY3Dot0NZ captures enum value "CC-BY 3.0 NZ"
	DatasetPutRequestLicenseCCDashBY3Dot0NZ string = "CC-BY 3.0 NZ"

	// DatasetPutRequestLicenseCCDashBYDashIGO captures enum value "CC-BY-IGO"
	DatasetPutRequestLicenseCCDashBYDashIGO string = "CC-BY-IGO"

	// DatasetPutRequestLicenseCCDashBYDashNC3Dot0 captures enum value "CC-BY-NC 3.0"
	DatasetPutRequestLicenseCCDashBYDashNC3Dot0 string = "CC-BY-NC 3.0"

	// DatasetPutRequestLicenseCCDashBYDashNC3Dot0NZ captures enum value "CC-BY-NC 3.0 NZ"
	DatasetPutRequestLicenseCCDashBYDashNC3Dot0NZ string = "CC-BY-NC 3.0 NZ"

	// DatasetPutRequestLicenseCCDashBYDashNCDashNDDashNZDash3Dot0 captures enum value "CC-BY-NC-ND-NZ-3.0"
	DatasetPutRequestLicenseCCDashBYDashNCDashNDDashNZDash3Dot0 string = "CC-BY-NC-ND-NZ-3.0"

	// DatasetPutRequestLicenseCCDashBYDashNCDashSA3Dot0 captures enum value "CC-BY-NC-SA 3.0"
	DatasetPutRequestLicenseCCDashBYDashNCDashSA3Dot0 string = "CC-BY-NC-SA 3.0"

	// DatasetPutRequestLicenseCCDashBYDashNCDashSA3Dot0NZ captures enum value "CC-BY-NC-SA 3.0 NZ"
	DatasetPutRequestLicenseCCDashBYDashNCDashSA3Dot0NZ string = "CC-BY-NC-SA 3.0 NZ"

	// DatasetPutRequestLicenseCCDashBYDashSA captures enum value "CC-BY-SA"
	DatasetPutRequestLicenseCCDashBYDashSA string = "CC-BY-SA"

	// DatasetPutRequestLicenseCCDashBYDashSA3Dot0 captures enum value "CC-BY-SA 3.0"
	DatasetPutRequestLicenseCCDashBYDashSA3Dot0 string = "CC-BY-SA 3.0"

	// DatasetPutRequestLicenseCCDashBYDashSA3Dot0NZ captures enum value "CC-BY-SA 3.0 NZ"
	DatasetPutRequestLicenseCCDashBYDashSA3Dot0NZ string = "CC-BY-SA 3.0 NZ"

	// DatasetPutRequestLicenseCDLADashPermissiveDash1Dot0 captures enum value "CDLA-Permissive-1.0"
	DatasetPutRequestLicenseCDLADashPermissiveDash1Dot0 string = "CDLA-Permissive-1.0"

	// DatasetPutRequestLicenseCDLADashSharingDash1Dot0 captures enum value "CDLA-Sharing-1.0"
	DatasetPutRequestLicenseCDLADashSharingDash1Dot0 string = "CDLA-Sharing-1.0"

	// DatasetPutRequestLicenseItalianDashODL captures enum value "Italian-ODL"
	DatasetPutRequestLicenseItalianDashODL string = "Italian-ODL"

	// DatasetPutRequestLicenseMITLicense captures enum value "MIT License"
	DatasetPutRequestLicenseMITLicense string = "MIT License"

	// DatasetPutRequestLicenseODCDashBY captures enum value "ODC-BY"
	DatasetPutRequestLicenseODCDashBY string = "ODC-BY"

	// DatasetPutRequestLicenseODCDashODbL captures enum value "ODC-ODbL"
	DatasetPutRequestLicenseODCDashODbL string = "ODC-ODbL"

	// DatasetPutRequestLicenseOGL captures enum value "OGL"
	DatasetPutRequestLicenseOGL string = "OGL"

	// DatasetPutRequestLicenseOGLDashCanada captures enum value "OGL-Canada"
	DatasetPutRequestLicenseOGLDashCanada string = "OGL-Canada"

	// DatasetPutRequestLicenseOGLDashNovaScotia captures enum value "OGL-Nova Scotia"
	DatasetPutRequestLicenseOGLDashNovaScotia string = "OGL-Nova Scotia"

	// DatasetPutRequestLicenseOGLDashUK captures enum value "OGL-UK"
	DatasetPutRequestLicenseOGLDashUK string = "OGL-UK"

	// DatasetPutRequestLicenseOSODL captures enum value "OSODL"
	DatasetPutRequestLicenseOSODL string = "OSODL"

	// DatasetPutRequestLicenseOther captures enum value "Other"
	DatasetPutRequestLicenseOther string = "Other"

	// DatasetPutRequestLicensePDDL captures enum value "PDDL"
	DatasetPutRequestLicensePDDL string = "PDDL"

	// DatasetPutRequestLicensePublicDomain captures enum value "Public Domain"
	DatasetPutRequestLicensePublicDomain string = "Public Domain"
)

// prop value enum
func (m *DatasetPutRequest) validateLicenseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, datasetPutRequestTypeLicensePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DatasetPutRequest) validateLicense(formats strfmt.Registry) error {
	if swag.IsZero(m.License) { // not required
		return nil
	}

	// value enum
	if err := m.validateLicenseEnum("license", "body", m.License); err != nil {
		return err
	}

	return nil
}

func (m *DatasetPutRequest) validateProperties(formats strfmt.Registry) error {
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

func (m *DatasetPutRequest) validateSummary(formats strfmt.Registry) error {
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

func (m *DatasetPutRequest) validateTitle(formats strfmt.Registry) error {

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

var datasetPutRequestTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datasetPutRequestTypeVisibilityPropEnum = append(datasetPutRequestTypeVisibilityPropEnum, v)
	}
}

const (

	// DatasetPutRequestVisibilityOPEN captures enum value "OPEN"
	DatasetPutRequestVisibilityOPEN string = "OPEN"

	// DatasetPutRequestVisibilityPRIVATE captures enum value "PRIVATE"
	DatasetPutRequestVisibilityPRIVATE string = "PRIVATE"
)

// prop value enum
func (m *DatasetPutRequest) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, datasetPutRequestTypeVisibilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DatasetPutRequest) validateVisibility(formats strfmt.Registry) error {

	if err := validate.Required("visibility", "body", m.Visibility); err != nil {
		return err
	}

	// value enum
	if err := m.validateVisibilityEnum("visibility", "body", *m.Visibility); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dataset put request based on the context it is used
func (m *DatasetPutRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasetPutRequest) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DatasetPutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasetPutRequest) UnmarshalBinary(b []byte) error {
	var res DatasetPutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
