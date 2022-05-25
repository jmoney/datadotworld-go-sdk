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

// NewDeleteCatalogTableParams creates a new DeleteCatalogTableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCatalogTableParams() *DeleteCatalogTableParams {
	return &DeleteCatalogTableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCatalogTableParamsWithTimeout creates a new DeleteCatalogTableParams object
// with the ability to set a timeout on a request.
func NewDeleteCatalogTableParamsWithTimeout(timeout time.Duration) *DeleteCatalogTableParams {
	return &DeleteCatalogTableParams{
		timeout: timeout,
	}
}

// NewDeleteCatalogTableParamsWithContext creates a new DeleteCatalogTableParams object
// with the ability to set a context for a request.
func NewDeleteCatalogTableParamsWithContext(ctx context.Context) *DeleteCatalogTableParams {
	return &DeleteCatalogTableParams{
		Context: ctx,
	}
}

// NewDeleteCatalogTableParamsWithHTTPClient creates a new DeleteCatalogTableParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCatalogTableParamsWithHTTPClient(client *http.Client) *DeleteCatalogTableParams {
	return &DeleteCatalogTableParams{
		HTTPClient: client,
	}
}

/* DeleteCatalogTableParams contains all the parameters to send to the API endpoint
   for the delete catalog table operation.

   Typically these are written to a http.Request.
*/
type DeleteCatalogTableParams struct {

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

// WithDefaults hydrates default values in the delete catalog table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCatalogTableParams) WithDefaults() *DeleteCatalogTableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete catalog table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCatalogTableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete catalog table params
func (o *DeleteCatalogTableParams) WithTimeout(timeout time.Duration) *DeleteCatalogTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete catalog table params
func (o *DeleteCatalogTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete catalog table params
func (o *DeleteCatalogTableParams) WithContext(ctx context.Context) *DeleteCatalogTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete catalog table params
func (o *DeleteCatalogTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete catalog table params
func (o *DeleteCatalogTableParams) WithHTTPClient(client *http.Client) *DeleteCatalogTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete catalog table params
func (o *DeleteCatalogTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the delete catalog table params
func (o *DeleteCatalogTableParams) WithOwner(owner string) *DeleteCatalogTableParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete catalog table params
func (o *DeleteCatalogTableParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithSourceid adds the sourceid to the delete catalog table params
func (o *DeleteCatalogTableParams) WithSourceid(sourceid string) *DeleteCatalogTableParams {
	o.SetSourceid(sourceid)
	return o
}

// SetSourceid adds the sourceid to the delete catalog table params
func (o *DeleteCatalogTableParams) SetSourceid(sourceid string) {
	o.Sourceid = sourceid
}

// WithTableid adds the tableid to the delete catalog table params
func (o *DeleteCatalogTableParams) WithTableid(tableid string) *DeleteCatalogTableParams {
	o.SetTableid(tableid)
	return o
}

// SetTableid adds the tableid to the delete catalog table params
func (o *DeleteCatalogTableParams) SetTableid(tableid string) {
	o.Tableid = tableid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCatalogTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
