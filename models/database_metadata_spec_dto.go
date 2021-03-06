// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DatabaseMetadataSpecDto database metadata spec dto
//
// swagger:model DatabaseMetadataSpecDto
type DatabaseMetadataSpecDto struct {

	// table specs
	// Max Items: 100
	// Min Items: 0
	TableSpecs []*TableMetadataSpecDto `json:"tableSpecs"`
}

// Validate validates this database metadata spec dto
func (m *DatabaseMetadataSpecDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTableSpecs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseMetadataSpecDto) validateTableSpecs(formats strfmt.Registry) error {
	if swag.IsZero(m.TableSpecs) { // not required
		return nil
	}

	iTableSpecsSize := int64(len(m.TableSpecs))

	if err := validate.MinItems("tableSpecs", "body", iTableSpecsSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("tableSpecs", "body", iTableSpecsSize, 100); err != nil {
		return err
	}

	for i := 0; i < len(m.TableSpecs); i++ {
		if swag.IsZero(m.TableSpecs[i]) { // not required
			continue
		}

		if m.TableSpecs[i] != nil {
			if err := m.TableSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tableSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tableSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this database metadata spec dto based on the context it is used
func (m *DatabaseMetadataSpecDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTableSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseMetadataSpecDto) contextValidateTableSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TableSpecs); i++ {

		if m.TableSpecs[i] != nil {
			if err := m.TableSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tableSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tableSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseMetadataSpecDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseMetadataSpecDto) UnmarshalBinary(b []byte) error {
	var res DatabaseMetadataSpecDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
