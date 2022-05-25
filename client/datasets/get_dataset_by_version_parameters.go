// Code generated by go-swagger; DO NOT EDIT.

package datasets

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

// NewGetDatasetByVersionParams creates a new GetDatasetByVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDatasetByVersionParams() *GetDatasetByVersionParams {
	return &GetDatasetByVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatasetByVersionParamsWithTimeout creates a new GetDatasetByVersionParams object
// with the ability to set a timeout on a request.
func NewGetDatasetByVersionParamsWithTimeout(timeout time.Duration) *GetDatasetByVersionParams {
	return &GetDatasetByVersionParams{
		timeout: timeout,
	}
}

// NewGetDatasetByVersionParamsWithContext creates a new GetDatasetByVersionParams object
// with the ability to set a context for a request.
func NewGetDatasetByVersionParamsWithContext(ctx context.Context) *GetDatasetByVersionParams {
	return &GetDatasetByVersionParams{
		Context: ctx,
	}
}

// NewGetDatasetByVersionParamsWithHTTPClient creates a new GetDatasetByVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDatasetByVersionParamsWithHTTPClient(client *http.Client) *GetDatasetByVersionParams {
	return &GetDatasetByVersionParams{
		HTTPClient: client,
	}
}

/* GetDatasetByVersionParams contains all the parameters to send to the API endpoint
   for the get dataset by version operation.

   Typically these are written to a http.Request.
*/
type GetDatasetByVersionParams struct {

	// ID.
	ID string

	// Owner.
	Owner string

	// VersionID.
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dataset by version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasetByVersionParams) WithDefaults() *GetDatasetByVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dataset by version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasetByVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dataset by version params
func (o *GetDatasetByVersionParams) WithTimeout(timeout time.Duration) *GetDatasetByVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dataset by version params
func (o *GetDatasetByVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dataset by version params
func (o *GetDatasetByVersionParams) WithContext(ctx context.Context) *GetDatasetByVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dataset by version params
func (o *GetDatasetByVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dataset by version params
func (o *GetDatasetByVersionParams) WithHTTPClient(client *http.Client) *GetDatasetByVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dataset by version params
func (o *GetDatasetByVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get dataset by version params
func (o *GetDatasetByVersionParams) WithID(id string) *GetDatasetByVersionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get dataset by version params
func (o *GetDatasetByVersionParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the get dataset by version params
func (o *GetDatasetByVersionParams) WithOwner(owner string) *GetDatasetByVersionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get dataset by version params
func (o *GetDatasetByVersionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithVersionID adds the versionID to the get dataset by version params
func (o *GetDatasetByVersionParams) WithVersionID(versionID string) *GetDatasetByVersionParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get dataset by version params
func (o *GetDatasetByVersionParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatasetByVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
