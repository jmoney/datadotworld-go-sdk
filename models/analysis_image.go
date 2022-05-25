// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AnalysisImage analysis image
//
// swagger:model AnalysisImage
type AnalysisImage struct {

	// external Url
	ExternalURL string `json:"externalUrl,omitempty"`
}

// Validate validates this analysis image
func (m *AnalysisImage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this analysis image based on context it is used
func (m *AnalysisImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnalysisImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalysisImage) UnmarshalBinary(b []byte) error {
	var res AnalysisImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
