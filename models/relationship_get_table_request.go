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

// RelationshipGetTableRequest relationship get table request
//
// swagger:model RelationshipGetTableRequest
type RelationshipGetTableRequest struct {

	// Filter by these relation types. Relation types should be IRIs.
	// Example: https://dwec.data.world/v0/usesDataFrom
	ByRelationTypes []string `json:"byRelationTypes"`

	// Filter by these resource types.
	ByResourceTypes []string `json:"byResourceTypes"`
}

// Validate validates this relationship get table request
func (m *RelationshipGetTableRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateByResourceTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var relationshipGetTableRequestByResourceTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CATALOG","ANALYSIS","BUSINESS_TERM","COLUMN","DATA_TYPE","DATASET","PROJECT","TABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		relationshipGetTableRequestByResourceTypesItemsEnum = append(relationshipGetTableRequestByResourceTypesItemsEnum, v)
	}
}

func (m *RelationshipGetTableRequest) validateByResourceTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, relationshipGetTableRequestByResourceTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RelationshipGetTableRequest) validateByResourceTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.ByResourceTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ByResourceTypes); i++ {

		// value enum
		if err := m.validateByResourceTypesItemsEnum("byResourceTypes"+"."+strconv.Itoa(i), "body", m.ByResourceTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this relationship get table request based on context it is used
func (m *RelationshipGetTableRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RelationshipGetTableRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipGetTableRequest) UnmarshalBinary(b []byte) error {
	var res RelationshipGetTableRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
