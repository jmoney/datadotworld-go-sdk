// Code generated by go-swagger; DO NOT EDIT.

package queries

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

// NewGetQueryVersionParams creates a new GetQueryVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQueryVersionParams() *GetQueryVersionParams {
	return &GetQueryVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQueryVersionParamsWithTimeout creates a new GetQueryVersionParams object
// with the ability to set a timeout on a request.
func NewGetQueryVersionParamsWithTimeout(timeout time.Duration) *GetQueryVersionParams {
	return &GetQueryVersionParams{
		timeout: timeout,
	}
}

// NewGetQueryVersionParamsWithContext creates a new GetQueryVersionParams object
// with the ability to set a context for a request.
func NewGetQueryVersionParamsWithContext(ctx context.Context) *GetQueryVersionParams {
	return &GetQueryVersionParams{
		Context: ctx,
	}
}

// NewGetQueryVersionParamsWithHTTPClient creates a new GetQueryVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQueryVersionParamsWithHTTPClient(client *http.Client) *GetQueryVersionParams {
	return &GetQueryVersionParams{
		HTTPClient: client,
	}
}

/* GetQueryVersionParams contains all the parameters to send to the API endpoint
   for the get query version operation.

   Typically these are written to a http.Request.
*/
type GetQueryVersionParams struct {

	// ID.
	ID string

	// VersionID.
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get query version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueryVersionParams) WithDefaults() *GetQueryVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get query version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueryVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get query version params
func (o *GetQueryVersionParams) WithTimeout(timeout time.Duration) *GetQueryVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get query version params
func (o *GetQueryVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get query version params
func (o *GetQueryVersionParams) WithContext(ctx context.Context) *GetQueryVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get query version params
func (o *GetQueryVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get query version params
func (o *GetQueryVersionParams) WithHTTPClient(client *http.Client) *GetQueryVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get query version params
func (o *GetQueryVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get query version params
func (o *GetQueryVersionParams) WithID(id string) *GetQueryVersionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get query version params
func (o *GetQueryVersionParams) SetID(id string) {
	o.ID = id
}

// WithVersionID adds the versionID to the get query version params
func (o *GetQueryVersionParams) WithVersionID(versionID string) *GetQueryVersionParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get query version params
func (o *GetQueryVersionParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetQueryVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
