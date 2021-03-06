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

// NewDeleteCatalogColumnParams creates a new DeleteCatalogColumnParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCatalogColumnParams() *DeleteCatalogColumnParams {
	return &DeleteCatalogColumnParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCatalogColumnParamsWithTimeout creates a new DeleteCatalogColumnParams object
// with the ability to set a timeout on a request.
func NewDeleteCatalogColumnParamsWithTimeout(timeout time.Duration) *DeleteCatalogColumnParams {
	return &DeleteCatalogColumnParams{
		timeout: timeout,
	}
}

// NewDeleteCatalogColumnParamsWithContext creates a new DeleteCatalogColumnParams object
// with the ability to set a context for a request.
func NewDeleteCatalogColumnParamsWithContext(ctx context.Context) *DeleteCatalogColumnParams {
	return &DeleteCatalogColumnParams{
		Context: ctx,
	}
}

// NewDeleteCatalogColumnParamsWithHTTPClient creates a new DeleteCatalogColumnParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCatalogColumnParamsWithHTTPClient(client *http.Client) *DeleteCatalogColumnParams {
	return &DeleteCatalogColumnParams{
		HTTPClient: client,
	}
}

/* DeleteCatalogColumnParams contains all the parameters to send to the API endpoint
   for the delete catalog column operation.

   Typically these are written to a http.Request.
*/
type DeleteCatalogColumnParams struct {

	// Columnid.
	Columnid string

	/* Owner.

	   id of the user or organization
	*/
	Owner string

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

// WithDefaults hydrates default values in the delete catalog column params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCatalogColumnParams) WithDefaults() *DeleteCatalogColumnParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete catalog column params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCatalogColumnParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete catalog column params
func (o *DeleteCatalogColumnParams) WithTimeout(timeout time.Duration) *DeleteCatalogColumnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete catalog column params
func (o *DeleteCatalogColumnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete catalog column params
func (o *DeleteCatalogColumnParams) WithContext(ctx context.Context) *DeleteCatalogColumnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete catalog column params
func (o *DeleteCatalogColumnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete catalog column params
func (o *DeleteCatalogColumnParams) WithHTTPClient(client *http.Client) *DeleteCatalogColumnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete catalog column params
func (o *DeleteCatalogColumnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColumnid adds the columnid to the delete catalog column params
func (o *DeleteCatalogColumnParams) WithColumnid(columnid string) *DeleteCatalogColumnParams {
	o.SetColumnid(columnid)
	return o
}

// SetColumnid adds the columnid to the delete catalog column params
func (o *DeleteCatalogColumnParams) SetColumnid(columnid string) {
	o.Columnid = columnid
}

// WithOwner adds the owner to the delete catalog column params
func (o *DeleteCatalogColumnParams) WithOwner(owner string) *DeleteCatalogColumnParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete catalog column params
func (o *DeleteCatalogColumnParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithSourceid adds the sourceid to the delete catalog column params
func (o *DeleteCatalogColumnParams) WithSourceid(sourceid string) *DeleteCatalogColumnParams {
	o.SetSourceid(sourceid)
	return o
}

// SetSourceid adds the sourceid to the delete catalog column params
func (o *DeleteCatalogColumnParams) SetSourceid(sourceid string) {
	o.Sourceid = sourceid
}

// WithTableid adds the tableid to the delete catalog column params
func (o *DeleteCatalogColumnParams) WithTableid(tableid string) *DeleteCatalogColumnParams {
	o.SetTableid(tableid)
	return o
}

// SetTableid adds the tableid to the delete catalog column params
func (o *DeleteCatalogColumnParams) SetTableid(tableid string) {
	o.Tableid = tableid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCatalogColumnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param columnid
	if err := r.SetPathParam("columnid", o.Columnid); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
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
