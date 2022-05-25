// Code generated by go-swagger; DO NOT EDIT.

package connections

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

// NewGetConnectionsByOwnerParams creates a new GetConnectionsByOwnerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConnectionsByOwnerParams() *GetConnectionsByOwnerParams {
	return &GetConnectionsByOwnerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConnectionsByOwnerParamsWithTimeout creates a new GetConnectionsByOwnerParams object
// with the ability to set a timeout on a request.
func NewGetConnectionsByOwnerParamsWithTimeout(timeout time.Duration) *GetConnectionsByOwnerParams {
	return &GetConnectionsByOwnerParams{
		timeout: timeout,
	}
}

// NewGetConnectionsByOwnerParamsWithContext creates a new GetConnectionsByOwnerParams object
// with the ability to set a context for a request.
func NewGetConnectionsByOwnerParamsWithContext(ctx context.Context) *GetConnectionsByOwnerParams {
	return &GetConnectionsByOwnerParams{
		Context: ctx,
	}
}

// NewGetConnectionsByOwnerParamsWithHTTPClient creates a new GetConnectionsByOwnerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConnectionsByOwnerParamsWithHTTPClient(client *http.Client) *GetConnectionsByOwnerParams {
	return &GetConnectionsByOwnerParams{
		HTTPClient: client,
	}
}

/* GetConnectionsByOwnerParams contains all the parameters to send to the API endpoint
   for the get connections by owner operation.

   Typically these are written to a http.Request.
*/
type GetConnectionsByOwnerParams struct {

	// From.
	From *string

	// Owner.
	Owner string

	// Size.
	Size *string

	// Sort.
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get connections by owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConnectionsByOwnerParams) WithDefaults() *GetConnectionsByOwnerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get connections by owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConnectionsByOwnerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get connections by owner params
func (o *GetConnectionsByOwnerParams) WithTimeout(timeout time.Duration) *GetConnectionsByOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get connections by owner params
func (o *GetConnectionsByOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get connections by owner params
func (o *GetConnectionsByOwnerParams) WithContext(ctx context.Context) *GetConnectionsByOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get connections by owner params
func (o *GetConnectionsByOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get connections by owner params
func (o *GetConnectionsByOwnerParams) WithHTTPClient(client *http.Client) *GetConnectionsByOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get connections by owner params
func (o *GetConnectionsByOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get connections by owner params
func (o *GetConnectionsByOwnerParams) WithFrom(from *string) *GetConnectionsByOwnerParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get connections by owner params
func (o *GetConnectionsByOwnerParams) SetFrom(from *string) {
	o.From = from
}

// WithOwner adds the owner to the get connections by owner params
func (o *GetConnectionsByOwnerParams) WithOwner(owner string) *GetConnectionsByOwnerParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get connections by owner params
func (o *GetConnectionsByOwnerParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithSize adds the size to the get connections by owner params
func (o *GetConnectionsByOwnerParams) WithSize(size *string) *GetConnectionsByOwnerParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get connections by owner params
func (o *GetConnectionsByOwnerParams) SetSize(size *string) {
	o.Size = size
}

// WithSort adds the sort to the get connections by owner params
func (o *GetConnectionsByOwnerParams) WithSort(sort *string) *GetConnectionsByOwnerParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get connections by owner params
func (o *GetConnectionsByOwnerParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetConnectionsByOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
