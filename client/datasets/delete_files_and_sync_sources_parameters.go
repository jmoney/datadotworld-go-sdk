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
	"github.com/go-openapi/swag"
)

// NewDeleteFilesAndSyncSourcesParams creates a new DeleteFilesAndSyncSourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteFilesAndSyncSourcesParams() *DeleteFilesAndSyncSourcesParams {
	return &DeleteFilesAndSyncSourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFilesAndSyncSourcesParamsWithTimeout creates a new DeleteFilesAndSyncSourcesParams object
// with the ability to set a timeout on a request.
func NewDeleteFilesAndSyncSourcesParamsWithTimeout(timeout time.Duration) *DeleteFilesAndSyncSourcesParams {
	return &DeleteFilesAndSyncSourcesParams{
		timeout: timeout,
	}
}

// NewDeleteFilesAndSyncSourcesParamsWithContext creates a new DeleteFilesAndSyncSourcesParams object
// with the ability to set a context for a request.
func NewDeleteFilesAndSyncSourcesParamsWithContext(ctx context.Context) *DeleteFilesAndSyncSourcesParams {
	return &DeleteFilesAndSyncSourcesParams{
		Context: ctx,
	}
}

// NewDeleteFilesAndSyncSourcesParamsWithHTTPClient creates a new DeleteFilesAndSyncSourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteFilesAndSyncSourcesParamsWithHTTPClient(client *http.Client) *DeleteFilesAndSyncSourcesParams {
	return &DeleteFilesAndSyncSourcesParams{
		HTTPClient: client,
	}
}

/* DeleteFilesAndSyncSourcesParams contains all the parameters to send to the API endpoint
   for the delete files and sync sources operation.

   Typically these are written to a http.Request.
*/
type DeleteFilesAndSyncSourcesParams struct {

	// ID.
	ID string

	// Name.
	Name []string

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete files and sync sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFilesAndSyncSourcesParams) WithDefaults() *DeleteFilesAndSyncSourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete files and sync sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFilesAndSyncSourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) WithTimeout(timeout time.Duration) *DeleteFilesAndSyncSourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) WithContext(ctx context.Context) *DeleteFilesAndSyncSourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) WithHTTPClient(client *http.Client) *DeleteFilesAndSyncSourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) WithID(id string) *DeleteFilesAndSyncSourcesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) WithName(name []string) *DeleteFilesAndSyncSourcesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) SetName(name []string) {
	o.Name = name
}

// WithOwner adds the owner to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) WithOwner(owner string) *DeleteFilesAndSyncSourcesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete files and sync sources params
func (o *DeleteFilesAndSyncSourcesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFilesAndSyncSourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Name != nil {

		// binding items for name
		joinedName := o.bindParamName(reg)

		// query array param name
		if err := r.SetQueryParam("name", joinedName...); err != nil {
			return err
		}
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

// bindParamDeleteFilesAndSyncSources binds the parameter name
func (o *DeleteFilesAndSyncSourcesParams) bindParamName(formats strfmt.Registry) []string {
	nameIR := o.Name

	var nameIC []string
	for _, nameIIR := range nameIR { // explode []string

		nameIIV := nameIIR // string as string
		nameIC = append(nameIC, nameIIV)
	}

	// items.CollectionFormat: "multi"
	nameIS := swag.JoinByFormat(nameIC, "multi")

	return nameIS
}
