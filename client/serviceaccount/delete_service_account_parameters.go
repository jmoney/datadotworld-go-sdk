// Code generated by go-swagger; DO NOT EDIT.

package serviceaccount

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

// NewDeleteServiceAccountParams creates a new DeleteServiceAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteServiceAccountParams() *DeleteServiceAccountParams {
	return &DeleteServiceAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServiceAccountParamsWithTimeout creates a new DeleteServiceAccountParams object
// with the ability to set a timeout on a request.
func NewDeleteServiceAccountParamsWithTimeout(timeout time.Duration) *DeleteServiceAccountParams {
	return &DeleteServiceAccountParams{
		timeout: timeout,
	}
}

// NewDeleteServiceAccountParamsWithContext creates a new DeleteServiceAccountParams object
// with the ability to set a context for a request.
func NewDeleteServiceAccountParamsWithContext(ctx context.Context) *DeleteServiceAccountParams {
	return &DeleteServiceAccountParams{
		Context: ctx,
	}
}

// NewDeleteServiceAccountParamsWithHTTPClient creates a new DeleteServiceAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteServiceAccountParamsWithHTTPClient(client *http.Client) *DeleteServiceAccountParams {
	return &DeleteServiceAccountParams{
		HTTPClient: client,
	}
}

/* DeleteServiceAccountParams contains all the parameters to send to the API endpoint
   for the delete service account operation.

   Typically these are written to a http.Request.
*/
type DeleteServiceAccountParams struct {

	// Owner.
	Owner string

	// ServiceAccount.
	ServiceAccount string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteServiceAccountParams) WithDefaults() *DeleteServiceAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteServiceAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete service account params
func (o *DeleteServiceAccountParams) WithTimeout(timeout time.Duration) *DeleteServiceAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete service account params
func (o *DeleteServiceAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete service account params
func (o *DeleteServiceAccountParams) WithContext(ctx context.Context) *DeleteServiceAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete service account params
func (o *DeleteServiceAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete service account params
func (o *DeleteServiceAccountParams) WithHTTPClient(client *http.Client) *DeleteServiceAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete service account params
func (o *DeleteServiceAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the delete service account params
func (o *DeleteServiceAccountParams) WithOwner(owner string) *DeleteServiceAccountParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete service account params
func (o *DeleteServiceAccountParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithServiceAccount adds the serviceAccount to the delete service account params
func (o *DeleteServiceAccountParams) WithServiceAccount(serviceAccount string) *DeleteServiceAccountParams {
	o.SetServiceAccount(serviceAccount)
	return o
}

// SetServiceAccount adds the serviceAccount to the delete service account params
func (o *DeleteServiceAccountParams) SetServiceAccount(serviceAccount string) {
	o.ServiceAccount = serviceAccount
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServiceAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param serviceAccount
	if err := r.SetPathParam("serviceAccount", o.ServiceAccount); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
