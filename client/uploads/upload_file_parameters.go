// Code generated by go-swagger; DO NOT EDIT.

package uploads

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
	"github.com/go-openapi/swag"
)

// NewUploadFileParams creates a new UploadFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadFileParams() *UploadFileParams {
	return &UploadFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadFileParamsWithTimeout creates a new UploadFileParams object
// with the ability to set a timeout on a request.
func NewUploadFileParamsWithTimeout(timeout time.Duration) *UploadFileParams {
	return &UploadFileParams{
		timeout: timeout,
	}
}

// NewUploadFileParamsWithContext creates a new UploadFileParams object
// with the ability to set a context for a request.
func NewUploadFileParamsWithContext(ctx context.Context) *UploadFileParams {
	return &UploadFileParams{
		Context: ctx,
	}
}

// NewUploadFileParamsWithHTTPClient creates a new UploadFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadFileParamsWithHTTPClient(client *http.Client) *UploadFileParams {
	return &UploadFileParams{
		HTTPClient: client,
	}
}

/* UploadFileParams contains all the parameters to send to the API endpoint
   for the upload file operation.

   Typically these are written to a http.Request.
*/
type UploadFileParams struct {

	// ExpandArchive.
	ExpandArchive *bool

	/* File.

	   Filename
	*/
	File string

	// ID.
	ID string

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadFileParams) WithDefaults() *UploadFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadFileParams) SetDefaults() {
	var (
		expandArchiveDefault = bool(false)
	)

	val := UploadFileParams{
		ExpandArchive: &expandArchiveDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the upload file params
func (o *UploadFileParams) WithTimeout(timeout time.Duration) *UploadFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload file params
func (o *UploadFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload file params
func (o *UploadFileParams) WithContext(ctx context.Context) *UploadFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload file params
func (o *UploadFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload file params
func (o *UploadFileParams) WithHTTPClient(client *http.Client) *UploadFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload file params
func (o *UploadFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpandArchive adds the expandArchive to the upload file params
func (o *UploadFileParams) WithExpandArchive(expandArchive *bool) *UploadFileParams {
	o.SetExpandArchive(expandArchive)
	return o
}

// SetExpandArchive adds the expandArchive to the upload file params
func (o *UploadFileParams) SetExpandArchive(expandArchive *bool) {
	o.ExpandArchive = expandArchive
}

// WithFile adds the file to the upload file params
func (o *UploadFileParams) WithFile(file string) *UploadFileParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the upload file params
func (o *UploadFileParams) SetFile(file string) {
	o.File = file
}

// WithID adds the id to the upload file params
func (o *UploadFileParams) WithID(id string) *UploadFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the upload file params
func (o *UploadFileParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the upload file params
func (o *UploadFileParams) WithOwner(owner string) *UploadFileParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the upload file params
func (o *UploadFileParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UploadFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExpandArchive != nil {

		// query param expandArchive
		var qrExpandArchive bool

		if o.ExpandArchive != nil {
			qrExpandArchive = *o.ExpandArchive
		}
		qExpandArchive := swag.FormatBool(qrExpandArchive)
		if qExpandArchive != "" {

			if err := r.SetQueryParam("expandArchive", qExpandArchive); err != nil {
				return err
			}
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
