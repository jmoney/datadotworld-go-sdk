// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewUnsubscribeFromProjectParams creates a new UnsubscribeFromProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnsubscribeFromProjectParams() *UnsubscribeFromProjectParams {
	return &UnsubscribeFromProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnsubscribeFromProjectParamsWithTimeout creates a new UnsubscribeFromProjectParams object
// with the ability to set a timeout on a request.
func NewUnsubscribeFromProjectParamsWithTimeout(timeout time.Duration) *UnsubscribeFromProjectParams {
	return &UnsubscribeFromProjectParams{
		timeout: timeout,
	}
}

// NewUnsubscribeFromProjectParamsWithContext creates a new UnsubscribeFromProjectParams object
// with the ability to set a context for a request.
func NewUnsubscribeFromProjectParamsWithContext(ctx context.Context) *UnsubscribeFromProjectParams {
	return &UnsubscribeFromProjectParams{
		Context: ctx,
	}
}

// NewUnsubscribeFromProjectParamsWithHTTPClient creates a new UnsubscribeFromProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnsubscribeFromProjectParamsWithHTTPClient(client *http.Client) *UnsubscribeFromProjectParams {
	return &UnsubscribeFromProjectParams{
		HTTPClient: client,
	}
}

/* UnsubscribeFromProjectParams contains all the parameters to send to the API endpoint
   for the unsubscribe from project operation.

   Typically these are written to a http.Request.
*/
type UnsubscribeFromProjectParams struct {

	// ID.
	ID string

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unsubscribe from project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnsubscribeFromProjectParams) WithDefaults() *UnsubscribeFromProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unsubscribe from project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnsubscribeFromProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unsubscribe from project params
func (o *UnsubscribeFromProjectParams) WithTimeout(timeout time.Duration) *UnsubscribeFromProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unsubscribe from project params
func (o *UnsubscribeFromProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unsubscribe from project params
func (o *UnsubscribeFromProjectParams) WithContext(ctx context.Context) *UnsubscribeFromProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unsubscribe from project params
func (o *UnsubscribeFromProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unsubscribe from project params
func (o *UnsubscribeFromProjectParams) WithHTTPClient(client *http.Client) *UnsubscribeFromProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unsubscribe from project params
func (o *UnsubscribeFromProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the unsubscribe from project params
func (o *UnsubscribeFromProjectParams) WithID(id string) *UnsubscribeFromProjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the unsubscribe from project params
func (o *UnsubscribeFromProjectParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the unsubscribe from project params
func (o *UnsubscribeFromProjectParams) WithOwner(owner string) *UnsubscribeFromProjectParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the unsubscribe from project params
func (o *UnsubscribeFromProjectParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UnsubscribeFromProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
