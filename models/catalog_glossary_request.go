// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CatalogGlossaryRequest catalog glossary request
//
// swagger:model CatalogGlossaryRequest
type CatalogGlossaryRequest struct {

	// Catalog Collection to which this metadata resource is added. Required for POST and PUT.Available catalog collection can be discovered via appropriate discovery endpoints.
	// Example: Aid Option
	Collections []string `json:"collections"`

	// A short, but descriptive statement about the metadata resource.
	// Example: Aid given to charity is an option that can be exercised by individuals and corporations
	// Max Length: 120
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// Custom properties for the metadata resource mapped to API BindingsCan be simple name-value string pairs or nested values for a string name. See examples for details.
	Properties map[string]JSONNode `json:"properties,omitempty"`

	// summary
	// Max Length: 120
	// Min Length: 0
	Summary *string `json:"summary,omitempty"`

	// A collection of tags to identify the relevance of metadata resource. Tags with no spaces is defacto standard
	// Example: HumanitarianAid
	Tags []string `json:"tags"`

	// Title of the metadata resource. Required for POST and PUT.
	// Example: crypto.snowflake
	// Max Length: 60
	// Min Length: 1
	Title string `json:"title,omitempty"`

	// Indicates the type of metadata resource. Some examples of valid values are Report, Tableau dashboard, Glossary, Table, Database view etc. Default values will be assumed if not provided. Defaults to Report for analysis resources, Glossary for Glossary resources, Table for Table resources and Column for Column resources. Once specified during creation, this cannot be changed via PATCH. Only a PUT can change the type
	// Example: Tableau dashboard
	TypeLabel string `json:"typeLabel,omitempty"`
}

// Validate validates this catalog glossary request
func (m *CatalogGlossaryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogGlossaryRequest) validateDescription(formats strfmt.Registry) error {
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

func (m *CatalogGlossaryRequest) validateProperties(formats strfmt.Registry) error {
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

func (m *CatalogGlossaryRequest) validateSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.Summary) { // not required
		return nil
	}

	if err := validate.MinLength("summary", "body", *m.Summary, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("summary", "body", *m.Summary, 120); err != nil {
		return err
	}

	return nil
}

func (m *CatalogGlossaryRequest) validateTitle(formats strfmt.Registry) error {
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

// ContextValidate validate this catalog glossary request based on the context it is used
func (m *CatalogGlossaryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogGlossaryRequest) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CatalogGlossaryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogGlossaryRequest) UnmarshalBinary(b []byte) error {
	var res CatalogGlossaryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
