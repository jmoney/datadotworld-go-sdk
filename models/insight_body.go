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

// InsightBody insight body
//
// swagger:model InsightBody
type InsightBody struct {

	// embed Url
	// Format: uri
	EmbedURL strfmt.URI `json:"embedUrl,omitempty"`

	// image Url
	// Format: uri
	ImageURL strfmt.URI `json:"imageUrl,omitempty"`

	// markdown body
	// Max Length: 25000
	// Min Length: 0
	MarkdownBody *string `json:"markdownBody,omitempty"`
}

// Validate validates this insight body
func (m *InsightBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmbedURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkdownBody(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InsightBody) validateEmbedURL(formats strfmt.Registry) error {
	if swag.IsZero(m.EmbedURL) { // not required
		return nil
	}

	if err := validate.FormatOf("embedUrl", "body", "uri", m.EmbedURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InsightBody) validateImageURL(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageURL) { // not required
		return nil
	}

	if err := validate.FormatOf("imageUrl", "body", "uri", m.ImageURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InsightBody) validateMarkdownBody(formats strfmt.Registry) error {
	if swag.IsZero(m.MarkdownBody) { // not required
		return nil
	}

	if err := validate.MinLength("markdownBody", "body", *m.MarkdownBody, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("markdownBody", "body", *m.MarkdownBody, 25000); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this insight body based on context it is used
func (m *InsightBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InsightBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InsightBody) UnmarshalBinary(b []byte) error {
	var res InsightBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
