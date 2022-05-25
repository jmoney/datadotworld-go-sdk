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

	"datadotworld-go-sdk/models"
)

// NewPatchFileMetadataParams creates a new PatchFileMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchFileMetadataParams() *PatchFileMetadataParams {
	return &PatchFileMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchFileMetadataParamsWithTimeout creates a new PatchFileMetadataParams object
// with the ability to set a timeout on a request.
func NewPatchFileMetadataParamsWithTimeout(timeout time.Duration) *PatchFileMetadataParams {
	return &PatchFileMetadataParams{
		timeout: timeout,
	}
}

// NewPatchFileMetadataParamsWithContext creates a new PatchFileMetadataParams object
// with the ability to set a context for a request.
func NewPatchFileMetadataParamsWithContext(ctx context.Context) *PatchFileMetadataParams {
	return &PatchFileMetadataParams{
		Context: ctx,
	}
}

// NewPatchFileMetadataParamsWithHTTPClient creates a new PatchFileMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchFileMetadataParamsWithHTTPClient(client *http.Client) *PatchFileMetadataParams {
	return &PatchFileMetadataParams{
		HTTPClient: client,
	}
}

/* PatchFileMetadataParams contains all the parameters to send to the API endpoint
   for the patch file metadata operation.

   Typically these are written to a http.Request.
*/
type PatchFileMetadataParams struct {

	// Body.
	Body *models.FileMetadataUpdateRequest

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

// WithDefaults hydrates default values in the patch file metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchFileMetadataParams) WithDefaults() *PatchFileMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch file metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchFileMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch file metadata params
func (o *PatchFileMetadataParams) WithTimeout(timeout time.Duration) *PatchFileMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch file metadata params
func (o *PatchFileMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch file metadata params
func (o *PatchFileMetadataParams) WithContext(ctx context.Context) *PatchFileMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch file metadata params
func (o *PatchFileMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch file metadata params
func (o *PatchFileMetadataParams) WithHTTPClient(client *http.Client) *PatchFileMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch file metadata params
func (o *PatchFileMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch file metadata params
func (o *PatchFileMetadataParams) WithBody(body *models.FileMetadataUpdateRequest) *PatchFileMetadataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch file metadata params
func (o *PatchFileMetadataParams) SetBody(body *models.FileMetadataUpdateRequest) {
	o.Body = body
}

// WithFile adds the file to the patch file metadata params
func (o *PatchFileMetadataParams) WithFile(file string) *PatchFileMetadataParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the patch file metadata params
func (o *PatchFileMetadataParams) SetFile(file string) {
	o.File = file
}

// WithID adds the id to the patch file metadata params
func (o *PatchFileMetadataParams) WithID(id string) *PatchFileMetadataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch file metadata params
func (o *PatchFileMetadataParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the patch file metadata params
func (o *PatchFileMetadataParams) WithOwner(owner string) *PatchFileMetadataParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch file metadata params
func (o *PatchFileMetadataParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *PatchFileMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
