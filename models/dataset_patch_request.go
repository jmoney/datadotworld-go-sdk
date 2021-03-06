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

// DatasetPatchRequest dataset patch request
//
// swagger:model DatasetPatchRequest
type DatasetPatchRequest struct {

	// asset status iri
	AssetStatusIri string `json:"assetStatusIri,omitempty"`

	// Interval in which files should be synced
	// Enum: [NEVER HOURLY DAILY WEEKLY]
	AutoSyncInterval string `json:"autoSyncInterval,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// license
	// Enum: [CC BY-NC CC BY-NC-ND CC BY-NC-SA CC BY-ND CC BY-ND 3.0 CC-0 CC-BY CC-BY 3.0 CC-BY 3.0 AU CC-BY 3.0 IGO CC-BY 3.0 NZ CC-BY-IGO CC-BY-NC 3.0 CC-BY-NC 3.0 NZ CC-BY-NC-ND-NZ-3.0 CC-BY-NC-SA 3.0 CC-BY-NC-SA 3.0 NZ CC-BY-SA CC-BY-SA 3.0 CC-BY-SA 3.0 NZ CDLA-Permissive-1.0 CDLA-Sharing-1.0 Italian-ODL MIT License ODC-BY ODC-ODbL OGL OGL-Canada OGL-Nova Scotia OGL-UK OSODL Other PDDL Public Domain]
	License string `json:"license,omitempty"`

	// properties
	Properties map[string]JSONNode `json:"properties,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// title
	// Max Length: 60
	// Min Length: 1
	Title string `json:"title,omitempty"`

	// visibility
	// Enum: [OPEN PRIVATE]
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this dataset patch request
func (m *DatasetPatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoSyncInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicense(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
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

var datasetPatchRequestTypeAutoSyncIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NEVER","HOURLY","DAILY","WEEKLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datasetPatchRequestTypeAutoSyncIntervalPropEnum = append(datasetPatchRequestTypeAutoSyncIntervalPropEnum, v)
	}
}

const (

	// DatasetPatchRequestAutoSyncIntervalNEVER captures enum value "NEVER"
	DatasetPatchRequestAutoSyncIntervalNEVER string = "NEVER"

	// DatasetPatchRequestAutoSyncIntervalHOURLY captures enum value "HOURLY"
	DatasetPatchRequestAutoSyncIntervalHOURLY string = "HOURLY"

	// DatasetPatchRequestAutoSyncIntervalDAILY captures enum value "DAILY"
	DatasetPatchRequestAutoSyncIntervalDAILY string = "DAILY"

	// DatasetPatchRequestAutoSyncIntervalWEEKLY captures enum value "WEEKLY"
	DatasetPatchRequestAutoSyncIntervalWEEKLY string = "WEEKLY"
)

// prop value enum
func (m *DatasetPatchRequest) validateAutoSyncIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, datasetPatchRequestTypeAutoSyncIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DatasetPatchRequest) validateAutoSyncInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoSyncInterval) { // not required
		return nil
	}

	// value enum
	if err := m.validateAutoSyncIntervalEnum("autoSyncInterval", "body", m.AutoSyncInterval); err != nil {
		return err
	}

	return nil
}

var datasetPatchRequestTypeLicensePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CC BY-NC","CC BY-NC-ND","CC BY-NC-SA","CC BY-ND","CC BY-ND 3.0","CC-0","CC-BY","CC-BY 3.0","CC-BY 3.0 AU","CC-BY 3.0 IGO","CC-BY 3.0 NZ","CC-BY-IGO","CC-BY-NC 3.0","CC-BY-NC 3.0 NZ","CC-BY-NC-ND-NZ-3.0","CC-BY-NC-SA 3.0","CC-BY-NC-SA 3.0 NZ","CC-BY-SA","CC-BY-SA 3.0","CC-BY-SA 3.0 NZ","CDLA-Permissive-1.0","CDLA-Sharing-1.0","Italian-ODL","MIT License","ODC-BY","ODC-ODbL","OGL","OGL-Canada","OGL-Nova Scotia","OGL-UK","OSODL","Other","PDDL","Public Domain"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datasetPatchRequestTypeLicensePropEnum = append(datasetPatchRequestTypeLicensePropEnum, v)
	}
}

const (

	// DatasetPatchRequestLicenseCCBYDashNC captures enum value "CC BY-NC"
	DatasetPatchRequestLicenseCCBYDashNC string = "CC BY-NC"

	// DatasetPatchRequestLicenseCCBYDashNCDashND captures enum value "CC BY-NC-ND"
	DatasetPatchRequestLicenseCCBYDashNCDashND string = "CC BY-NC-ND"

	// DatasetPatchRequestLicenseCCBYDashNCDashSA captures enum value "CC BY-NC-SA"
	DatasetPatchRequestLicenseCCBYDashNCDashSA string = "CC BY-NC-SA"

	// DatasetPatchRequestLicenseCCBYDashND captures enum value "CC BY-ND"
	DatasetPatchRequestLicenseCCBYDashND string = "CC BY-ND"

	// DatasetPatchRequestLicenseCCBYDashND3Dot0 captures enum value "CC BY-ND 3.0"
	DatasetPatchRequestLicenseCCBYDashND3Dot0 string = "CC BY-ND 3.0"

	// DatasetPatchRequestLicenseCCDash0 captures enum value "CC-0"
	DatasetPatchRequestLicenseCCDash0 string = "CC-0"

	// DatasetPatchRequestLicenseCCDashBY captures enum value "CC-BY"
	DatasetPatchRequestLicenseCCDashBY string = "CC-BY"

	// DatasetPatchRequestLicenseCCDashBY3Dot0 captures enum value "CC-BY 3.0"
	DatasetPatchRequestLicenseCCDashBY3Dot0 string = "CC-BY 3.0"

	// DatasetPatchRequestLicenseCCDashBY3Dot0AU captures enum value "CC-BY 3.0 AU"
	DatasetPatchRequestLicenseCCDashBY3Dot0AU string = "CC-BY 3.0 AU"

	// DatasetPatchRequestLicenseCCDashBY3Dot0IGO captures enum value "CC-BY 3.0 IGO"
	DatasetPatchRequestLicenseCCDashBY3Dot0IGO string = "CC-BY 3.0 IGO"

	// DatasetPatchRequestLicenseCCDashBY3Dot0NZ captures enum value "CC-BY 3.0 NZ"
	DatasetPatchRequestLicenseCCDashBY3Dot0NZ string = "CC-BY 3.0 NZ"

	// DatasetPatchRequestLicenseCCDashBYDashIGO captures enum value "CC-BY-IGO"
	DatasetPatchRequestLicenseCCDashBYDashIGO string = "CC-BY-IGO"

	// DatasetPatchRequestLicenseCCDashBYDashNC3Dot0 captures enum value "CC-BY-NC 3.0"
	DatasetPatchRequestLicenseCCDashBYDashNC3Dot0 string = "CC-BY-NC 3.0"

	// DatasetPatchRequestLicenseCCDashBYDashNC3Dot0NZ captures enum value "CC-BY-NC 3.0 NZ"
	DatasetPatchRequestLicenseCCDashBYDashNC3Dot0NZ string = "CC-BY-NC 3.0 NZ"

	// DatasetPatchRequestLicenseCCDashBYDashNCDashNDDashNZDash3Dot0 captures enum value "CC-BY-NC-ND-NZ-3.0"
	DatasetPatchRequestLicenseCCDashBYDashNCDashNDDashNZDash3Dot0 string = "CC-BY-NC-ND-NZ-3.0"

	// DatasetPatchRequestLicenseCCDashBYDashNCDashSA3Dot0 captures enum value "CC-BY-NC-SA 3.0"
	DatasetPatchRequestLicenseCCDashBYDashNCDashSA3Dot0 string = "CC-BY-NC-SA 3.0"

	// DatasetPatchRequestLicenseCCDashBYDashNCDashSA3Dot0NZ captures enum value "CC-BY-NC-SA 3.0 NZ"
	DatasetPatchRequestLicenseCCDashBYDashNCDashSA3Dot0NZ string = "CC-BY-NC-SA 3.0 NZ"

	// DatasetPatchRequestLicenseCCDashBYDashSA captures enum value "CC-BY-SA"
	DatasetPatchRequestLicenseCCDashBYDashSA string = "CC-BY-SA"

	// DatasetPatchRequestLicenseCCDashBYDashSA3Dot0 captures enum value "CC-BY-SA 3.0"
	DatasetPatchRequestLicenseCCDashBYDashSA3Dot0 string = "CC-BY-SA 3.0"

	// DatasetPatchRequestLicenseCCDashBYDashSA3Dot0NZ captures enum value "CC-BY-SA 3.0 NZ"
	DatasetPatchRequestLicenseCCDashBYDashSA3Dot0NZ string = "CC-BY-SA 3.0 NZ"

	// DatasetPatchRequestLicenseCDLADashPermissiveDash1Dot0 captures enum value "CDLA-Permissive-1.0"
	DatasetPatchRequestLicenseCDLADashPermissiveDash1Dot0 string = "CDLA-Permissive-1.0"

	// DatasetPatchRequestLicenseCDLADashSharingDash1Dot0 captures enum value "CDLA-Sharing-1.0"
	DatasetPatchRequestLicenseCDLADashSharingDash1Dot0 string = "CDLA-Sharing-1.0"

	// DatasetPatchRequestLicenseItalianDashODL captures enum value "Italian-ODL"
	DatasetPatchRequestLicenseItalianDashODL string = "Italian-ODL"

	// DatasetPatchRequestLicenseMITLicense captures enum value "MIT License"
	DatasetPatchRequestLicenseMITLicense string = "MIT License"

	// DatasetPatchRequestLicenseODCDashBY captures enum value "ODC-BY"
	DatasetPatchRequestLicenseODCDashBY string = "ODC-BY"

	// DatasetPatchRequestLicenseODCDashODbL captures enum value "ODC-ODbL"
	DatasetPatchRequestLicenseODCDashODbL string = "ODC-ODbL"

	// DatasetPatchRequestLicenseOGL captures enum value "OGL"
	DatasetPatchRequestLicenseOGL string = "OGL"

	// DatasetPatchRequestLicenseOGLDashCanada captures enum value "OGL-Canada"
	DatasetPatchRequestLicenseOGLDashCanada string = "OGL-Canada"

	// DatasetPatchRequestLicenseOGLDashNovaScotia captures enum value "OGL-Nova Scotia"
	DatasetPatchRequestLicenseOGLDashNovaScotia string = "OGL-Nova Scotia"

	// DatasetPatchRequestLicenseOGLDashUK captures enum value "OGL-UK"
	DatasetPatchRequestLicenseOGLDashUK string = "OGL-UK"

	// DatasetPatchRequestLicenseOSODL captures enum value "OSODL"
	DatasetPatchRequestLicenseOSODL string = "OSODL"

	// DatasetPatchRequestLicenseOther captures enum value "Other"
	DatasetPatchRequestLicenseOther string = "Other"

	// DatasetPatchRequestLicensePDDL captures enum value "PDDL"
	DatasetPatchRequestLicensePDDL string = "PDDL"

	// DatasetPatchRequestLicensePublicDomain captures enum value "Public Domain"
	DatasetPatchRequestLicensePublicDomain string = "Public Domain"
)

// prop value enum
func (m *DatasetPatchRequest) validateLicenseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, datasetPatchRequestTypeLicensePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DatasetPatchRequest) validateLicense(formats strfmt.Registry) error {
	if swag.IsZero(m.License) { // not required
		return nil
	}

	// value enum
	if err := m.validateLicenseEnum("license", "body", m.License); err != nil {
		return err
	}

	return nil
}

func (m *DatasetPatchRequest) validateProperties(formats strfmt.Registry) error {
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

func (m *DatasetPatchRequest) validateTitle(formats strfmt.Registry) error {
	if swag.IsZero(m.Title) { // not required
		return nil
	}

	if err := validate.MinLength("title", "body", m.Title, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", m.Title, 60); err != nil {
		return err
	}

	return nil
}

var datasetPatchRequestTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datasetPatchRequestTypeVisibilityPropEnum = append(datasetPatchRequestTypeVisibilityPropEnum, v)
	}
}

const (

	// DatasetPatchRequestVisibilityOPEN captures enum value "OPEN"
	DatasetPatchRequestVisibilityOPEN string = "OPEN"

	// DatasetPatchRequestVisibilityPRIVATE captures enum value "PRIVATE"
	DatasetPatchRequestVisibilityPRIVATE string = "PRIVATE"
)

// prop value enum
func (m *DatasetPatchRequest) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, datasetPatchRequestTypeVisibilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DatasetPatchRequest) validateVisibility(formats strfmt.Registry) error {
	if swag.IsZero(m.Visibility) { // not required
		return nil
	}

	// value enum
	if err := m.validateVisibilityEnum("visibility", "body", m.Visibility); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dataset patch request based on the context it is used
func (m *DatasetPatchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasetPatchRequest) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DatasetPatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasetPatchRequest) UnmarshalBinary(b []byte) error {
	var res DatasetPatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
