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

// NewGetFileMetadataParams creates a new GetFileMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFileMetadataParams() *GetFileMetadataParams {
	return &GetFileMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFileMetadataParamsWithTimeout creates a new GetFileMetadataParams object
// with the ability to set a timeout on a request.
func NewGetFileMetadataParamsWithTimeout(timeout time.Duration) *GetFileMetadataParams {
	return &GetFileMetadataParams{
		timeout: timeout,
	}
}

// NewGetFileMetadataParamsWithContext creates a new GetFileMetadataParams object
// with the ability to set a context for a request.
func NewGetFileMetadataParamsWithContext(ctx context.Context) *GetFileMetadataParams {
	return &GetFileMetadataParams{
		Context: ctx,
	}
}

// NewGetFileMetadataParamsWithHTTPClient creates a new GetFileMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFileMetadataParamsWithHTTPClient(client *http.Client) *GetFileMetadataParams {
	return &GetFileMetadataParams{
		HTTPClient: client,
	}
}

/* GetFileMetadataParams contains all the parameters to send to the API endpoint
   for the get file metadata operation.

   Typically these are written to a http.Request.
*/
type GetFileMetadataParams struct {

	// File.
	File string

	// ID.
	ID string

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get file metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileMetadataParams) WithDefaults() *GetFileMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get file metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get file metadata params
func (o *GetFileMetadataParams) WithTimeout(timeout time.Duration) *GetFileMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file metadata params
func (o *GetFileMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file metadata params
func (o *GetFileMetadataParams) WithContext(ctx context.Context) *GetFileMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file metadata params
func (o *GetFileMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file metadata params
func (o *GetFileMetadataParams) WithHTTPClient(client *http.Client) *GetFileMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file metadata params
func (o *GetFileMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the get file metadata params
func (o *GetFileMetadataParams) WithFile(file string) *GetFileMetadataParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the get file metadata params
func (o *GetFileMetadataParams) SetFile(file string) {
	o.File = file
}

// WithID adds the id to the get file metadata params
func (o *GetFileMetadataParams) WithID(id string) *GetFileMetadataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get file metadata params
func (o *GetFileMetadataParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the get file metadata params
func (o *GetFileMetadataParams) WithOwner(owner string) *GetFileMetadataParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get file metadata params
func (o *GetFileMetadataParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *GetFileMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param file
	if err := r.SetPathParam("file", o.File); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
