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

	"datadotworld-go-sdk/models"
)

// NewSubscribeToProjectParams creates a new SubscribeToProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscribeToProjectParams() *SubscribeToProjectParams {
	return &SubscribeToProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscribeToProjectParamsWithTimeout creates a new SubscribeToProjectParams object
// with the ability to set a timeout on a request.
func NewSubscribeToProjectParamsWithTimeout(timeout time.Duration) *SubscribeToProjectParams {
	return &SubscribeToProjectParams{
		timeout: timeout,
	}
}

// NewSubscribeToProjectParamsWithContext creates a new SubscribeToProjectParams object
// with the ability to set a context for a request.
func NewSubscribeToProjectParamsWithContext(ctx context.Context) *SubscribeToProjectParams {
	return &SubscribeToProjectParams{
		Context: ctx,
	}
}

// NewSubscribeToProjectParamsWithHTTPClient creates a new SubscribeToProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscribeToProjectParamsWithHTTPClient(client *http.Client) *SubscribeToProjectParams {
	return &SubscribeToProjectParams{
		HTTPClient: client,
	}
}

/* SubscribeToProjectParams contains all the parameters to send to the API endpoint
   for the subscribe to project operation.

   Typically these are written to a http.Request.
*/
type SubscribeToProjectParams struct {

	// Body.
	Body *models.SubscriptionCreateRequest

	// ID.
	ID string

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscribe to project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscribeToProjectParams) WithDefaults() *SubscribeToProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscribe to project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscribeToProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscribe to project params
func (o *SubscribeToProjectParams) WithTimeout(timeout time.Duration) *SubscribeToProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscribe to project params
func (o *SubscribeToProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscribe to project params
func (o *SubscribeToProjectParams) WithContext(ctx context.Context) *SubscribeToProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscribe to project params
func (o *SubscribeToProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscribe to project params
func (o *SubscribeToProjectParams) WithHTTPClient(client *http.Client) *SubscribeToProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscribe to project params
func (o *SubscribeToProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscribe to project params
func (o *SubscribeToProjectParams) WithBody(body *models.SubscriptionCreateRequest) *SubscribeToProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscribe to project params
func (o *SubscribeToProjectParams) SetBody(body *models.SubscriptionCreateRequest) {
	o.Body = body
}

// WithID adds the id to the subscribe to project params
func (o *SubscribeToProjectParams) WithID(id string) *SubscribeToProjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the subscribe to project params
func (o *SubscribeToProjectParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the subscribe to project params
func (o *SubscribeToProjectParams) WithOwner(owner string) *SubscribeToProjectParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the subscribe to project params
func (o *SubscribeToProjectParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *SubscribeToProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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