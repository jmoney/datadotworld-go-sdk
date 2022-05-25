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

// CreateQueryRequest create query request
//
// swagger:model CreateQueryRequest
type CreateQueryRequest struct {

	// content
	// Required: true
	Content *string `json:"content"`

	// description
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// language
	// Required: true
	// Enum: [SPARQL SQL]
	Language *string `json:"language"`

	// name
	// Required: true
	// Max Length: 128
	// Min Length: 0
	Name *string `json:"name"`

	// published
	Published *bool `json:"published,omitempty"`
}

// Validate validates this create query request
func (m *CreateQueryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateQueryRequest) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *CreateQueryRequest) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", *m.Description, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 1000); err != nil {
		return err
	}

	return nil
}

var createQueryRequestTypeLanguagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SPARQL","SQL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createQueryRequestTypeLanguagePropEnum = append(createQueryRequestTypeLanguagePropEnum, v)
	}
}

const (

	// CreateQueryRequestLanguageSPARQL captures enum value "SPARQL"
	CreateQueryRequestLanguageSPARQL string = "SPARQL"

	// CreateQueryRequestLanguageSQL captures enum value "SQL"
	CreateQueryRequestLanguageSQL string = "SQL"
)

// prop value enum
func (m *CreateQueryRequest) validateLanguageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createQueryRequestTypeLanguagePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateQueryRequest) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	// value enum
	if err := m.validateLanguageEnum("language", "body", *m.Language); err != nil {
		return err
	}

	return nil
}

func (m *CreateQueryRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 128); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create query request based on context it is used
func (m *CreateQueryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateQueryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateQueryRequest) UnmarshalBinary(b []byte) error {
	var res CreateQueryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}