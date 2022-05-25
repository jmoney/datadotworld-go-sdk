// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DwccSpecDto dwcc spec dto
//
// swagger:model DwccSpecDto
type DwccSpecDto struct {

	// database Id
	DatabaseID string `json:"databaseId,omitempty"`

	// source database
	SourceDatabase string `json:"sourceDatabase,omitempty"`

	// source schemas
	SourceSchemas []string `json:"sourceSchemas"`

	// target catalog
	TargetCatalog string `json:"targetCatalog,omitempty"`
}

// Validate validates this dwcc spec dto
func (m *DwccSpecDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dwcc spec dto based on context it is used
func (m *DwccSpecDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DwccSpecDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DwccSpecDto) UnmarshalBinary(b []byte) error {
	var res DwccSpecDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}