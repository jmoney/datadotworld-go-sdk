// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QueryParameter query parameter
//
// swagger:model QueryParameter
type QueryParameter struct {

	// datatype
	Datatype string `json:"datatype,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this query parameter
func (m *QueryParameter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query parameter based on context it is used
func (m *QueryParameter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryParameter) UnmarshalBinary(b []byte) error {
	var res QueryParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
