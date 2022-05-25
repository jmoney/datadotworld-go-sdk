// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SearchFacetResult search facet result
//
// swagger:model SearchFacetResult
type SearchFacetResult struct {

	// count
	Count int64 `json:"count,omitempty"`

	// term
	Term string `json:"term,omitempty"`
}

// Validate validates this search facet result
func (m *SearchFacetResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search facet result based on context it is used
func (m *SearchFacetResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SearchFacetResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchFacetResult) UnmarshalBinary(b []byte) error {
	var res SearchFacetResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
