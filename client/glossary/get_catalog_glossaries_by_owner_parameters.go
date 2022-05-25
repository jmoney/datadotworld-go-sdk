// Code generated by go-swagger; DO NOT EDIT.

package glossary

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetCatalogGlossariesByOwnerParams creates a new GetCatalogGlossariesByOwnerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCatalogGlossariesByOwnerParams() *GetCatalogGlossariesByOwnerParams {
	return &GetCatalogGlossariesByOwnerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogGlossariesByOwnerParamsWithTimeout creates a new GetCatalogGlossariesByOwnerParams object
// with the ability to set a timeout on a request.
func NewGetCatalogGlossariesByOwnerParamsWithTimeout(timeout time.Duration) *GetCatalogGlossariesByOwnerParams {
	return &GetCatalogGlossariesByOwnerParams{
		timeout: timeout,
	}
}

// NewGetCatalogGlossariesByOwnerParamsWithContext creates a new GetCatalogGlossariesByOwnerParams object
// with the ability to set a context for a request.
func NewGetCatalogGlossariesByOwnerParamsWithContext(ctx context.Context) *GetCatalogGlossariesByOwnerParams {
	return &GetCatalogGlossariesByOwnerParams{
		Context: ctx,
	}
}

// NewGetCatalogGlossariesByOwnerParamsWithHTTPClient creates a new GetCatalogGlossariesByOwnerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCatalogGlossariesByOwnerParamsWithHTTPClient(client *http.Client) *GetCatalogGlossariesByOwnerParams {
	return &GetCatalogGlossariesByOwnerParams{
		HTTPClient: client,
	}
}

/* GetCatalogGlossariesByOwnerParams contains all the parameters to send to the API endpoint
   for the get catalog glossaries by owner operation.

   Typically these are written to a http.Request.
*/
type GetCatalogGlossariesByOwnerParams struct {

	// From.
	From *string

	// Owner.
	Owner string

	// Size.
	Size *string

	// Sort.
	Sort *string

	// Typeiri.
	Typeiri *string

	// Typelabel.
	Typelabel *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get catalog glossaries by owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCatalogGlossariesByOwnerParams) WithDefaults() *GetCatalogGlossariesByOwnerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get catalog glossaries by owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCatalogGlossariesByOwnerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) WithTimeout(timeout time.Duration) *GetCatalogGlossariesByOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) WithContext(ctx context.Context) *GetCatalogGlossariesByOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) WithHTTPClient(client *http.Client) *GetCatalogGlossariesByOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) WithFrom(from *string) *GetCatalogGlossariesByOwnerParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) SetFrom(from *string) {
	o.From = from
}

// WithOwner adds the owner to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) WithOwner(owner string) *GetCatalogGlossariesByOwnerParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithSize adds the size to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) WithSize(size *string) *GetCatalogGlossariesByOwnerParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) SetSize(size *string) {
	o.Size = size
}

// WithSort adds the sort to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) WithSort(sort *string) *GetCatalogGlossariesByOwnerParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTypeiri adds the typeiri to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) WithTypeiri(typeiri *string) *GetCatalogGlossariesByOwnerParams {
	o.SetTypeiri(typeiri)
	return o
}

// SetTypeiri adds the typeiri to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) SetTypeiri(typeiri *string) {
	o.Typeiri = typeiri
}

// WithTypelabel adds the typelabel to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) WithTypelabel(typelabel *string) *GetCatalogGlossariesByOwnerParams {
	o.SetTypelabel(typelabel)
	return o
}

// SetTypelabel adds the typelabel to the get catalog glossaries by owner params
func (o *GetCatalogGlossariesByOwnerParams) SetTypelabel(typelabel *string) {
	o.Typelabel = typelabel
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogGlossariesByOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.From != nil {

		// query param from
		var qrFrom string

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Size != nil {

		// query param size
		var qrSize string

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := qrSize
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.Typeiri != nil {

		// query param typeiri
		var qrTypeiri string

		if o.Typeiri != nil {
			qrTypeiri = *o.Typeiri
		}
		qTypeiri := qrTypeiri
		if qTypeiri != "" {

			if err := r.SetQueryParam("typeiri", qTypeiri); err != nil {
				return err
			}
		}
	}

	if o.Typelabel != nil {

		// query param typelabel
		var qrTypelabel string

		if o.Typelabel != nil {
			qrTypelabel = *o.Typelabel
		}
		qTypelabel := qrTypelabel
		if qTypelabel != "" {

			if err := r.SetQueryParam("typelabel", qTypelabel); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
