// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewCreateProjectSavedQueryParams creates a new CreateProjectSavedQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProjectSavedQueryParams() *CreateProjectSavedQueryParams {
	return &CreateProjectSavedQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectSavedQueryParamsWithTimeout creates a new CreateProjectSavedQueryParams object
// with the ability to set a timeout on a request.
func NewCreateProjectSavedQueryParamsWithTimeout(timeout time.Duration) *CreateProjectSavedQueryParams {
	return &CreateProjectSavedQueryParams{
		timeout: timeout,
	}
}

// NewCreateProjectSavedQueryParamsWithContext creates a new CreateProjectSavedQueryParams object
// with the ability to set a context for a request.
func NewCreateProjectSavedQueryParamsWithContext(ctx context.Context) *CreateProjectSavedQueryParams {
	return &CreateProjectSavedQueryParams{
		Context: ctx,
	}
}

// NewCreateProjectSavedQueryParamsWithHTTPClient creates a new CreateProjectSavedQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProjectSavedQueryParamsWithHTTPClient(client *http.Client) *CreateProjectSavedQueryParams {
	return &CreateProjectSavedQueryParams{
		HTTPClient: client,
	}
}

/* CreateProjectSavedQueryParams contains all the parameters to send to the API endpoint
   for the create project saved query operation.

   Typically these are written to a http.Request.
*/
type CreateProjectSavedQueryParams struct {

	// Body.
	Body *models.CreateQueryRequest

	// ID.
	ID string

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create project saved query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProjectSavedQueryParams) WithDefaults() *CreateProjectSavedQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create project saved query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProjectSavedQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create project saved query params
func (o *CreateProjectSavedQueryParams) WithTimeout(timeout time.Duration) *CreateProjectSavedQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project saved query params
func (o *CreateProjectSavedQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project saved query params
func (o *CreateProjectSavedQueryParams) WithContext(ctx context.Context) *CreateProjectSavedQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project saved query params
func (o *CreateProjectSavedQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project saved query params
func (o *CreateProjectSavedQueryParams) WithHTTPClient(client *http.Client) *CreateProjectSavedQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project saved query params
func (o *CreateProjectSavedQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create project saved query params
func (o *CreateProjectSavedQueryParams) WithBody(body *models.CreateQueryRequest) *CreateProjectSavedQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create project saved query params
func (o *CreateProjectSavedQueryParams) SetBody(body *models.CreateQueryRequest) {
	o.Body = body
}

// WithID adds the id to the create project saved query params
func (o *CreateProjectSavedQueryParams) WithID(id string) *CreateProjectSavedQueryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create project saved query params
func (o *CreateProjectSavedQueryParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the create project saved query params
func (o *CreateProjectSavedQueryParams) WithOwner(owner string) *CreateProjectSavedQueryParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create project saved query params
func (o *CreateProjectSavedQueryParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProjectSavedQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
