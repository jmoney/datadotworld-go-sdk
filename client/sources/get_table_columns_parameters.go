// Code generated by go-swagger; DO NOT EDIT.

package sources

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

// NewGetTableColumnsParams creates a new GetTableColumnsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTableColumnsParams() *GetTableColumnsParams {
	return &GetTableColumnsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTableColumnsParamsWithTimeout creates a new GetTableColumnsParams object
// with the ability to set a timeout on a request.
func NewGetTableColumnsParamsWithTimeout(timeout time.Duration) *GetTableColumnsParams {
	return &GetTableColumnsParams{
		timeout: timeout,
	}
}

// NewGetTableColumnsParamsWithContext creates a new GetTableColumnsParams object
// with the ability to set a context for a request.
func NewGetTableColumnsParamsWithContext(ctx context.Context) *GetTableColumnsParams {
	return &GetTableColumnsParams{
		Context: ctx,
	}
}

// NewGetTableColumnsParamsWithHTTPClient creates a new GetTableColumnsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTableColumnsParamsWithHTTPClient(client *http.Client) *GetTableColumnsParams {
	return &GetTableColumnsParams{
		HTTPClient: client,
	}
}

/* GetTableColumnsParams contains all the parameters to send to the API endpoint
   for the get table columns operation.

   Typically these are written to a http.Request.
*/
type GetTableColumnsParams struct {

	// From.
	From *string

	/* Owner.

	   id of the user or organization
	*/
	Owner string

	// Size.
	Size *string

	// Sort.
	Sort *string

	/* Sourceid.

	   database source id
	*/
	Sourceid string

	// Tableid.
	Tableid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get table columns params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTableColumnsParams) WithDefaults() *GetTableColumnsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get table columns params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTableColumnsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get table columns params
func (o *GetTableColumnsParams) WithTimeout(timeout time.Duration) *GetTableColumnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get table columns params
func (o *GetTableColumnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get table columns params
func (o *GetTableColumnsParams) WithContext(ctx context.Context) *GetTableColumnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get table columns params
func (o *GetTableColumnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get table columns params
func (o *GetTableColumnsParams) WithHTTPClient(client *http.Client) *GetTableColumnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get table columns params
func (o *GetTableColumnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get table columns params
func (o *GetTableColumnsParams) WithFrom(from *string) *GetTableColumnsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get table columns params
func (o *GetTableColumnsParams) SetFrom(from *string) {
	o.From = from
}

// WithOwner adds the owner to the get table columns params
func (o *GetTableColumnsParams) WithOwner(owner string) *GetTableColumnsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get table columns params
func (o *GetTableColumnsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithSize adds the size to the get table columns params
func (o *GetTableColumnsParams) WithSize(size *string) *GetTableColumnsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get table columns params
func (o *GetTableColumnsParams) SetSize(size *string) {
	o.Size = size
}

// WithSort adds the sort to the get table columns params
func (o *GetTableColumnsParams) WithSort(sort *string) *GetTableColumnsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get table columns params
func (o *GetTableColumnsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithSourceid adds the sourceid to the get table columns params
func (o *GetTableColumnsParams) WithSourceid(sourceid string) *GetTableColumnsParams {
	o.SetSourceid(sourceid)
	return o
}

// SetSourceid adds the sourceid to the get table columns params
func (o *GetTableColumnsParams) SetSourceid(sourceid string) {
	o.Sourceid = sourceid
}

// WithTableid adds the tableid to the get table columns params
func (o *GetTableColumnsParams) WithTableid(tableid string) *GetTableColumnsParams {
	o.SetTableid(tableid)
	return o
}

// SetTableid adds the tableid to the get table columns params
func (o *GetTableColumnsParams) SetTableid(tableid string) {
	o.Tableid = tableid
}

// WriteToRequest writes these params to a swagger request
func (o *GetTableColumnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param sourceid
	if err := r.SetPathParam("sourceid", o.Sourceid); err != nil {
		return err
	}

	// path param tableid
	if err := r.SetPathParam("tableid", o.Tableid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
