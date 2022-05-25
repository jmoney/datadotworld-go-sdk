// Code generated by go-swagger; DO NOT EDIT.

package collections

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

// NewCreateCatalogParams creates a new CreateCatalogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCatalogParams() *CreateCatalogParams {
	return &CreateCatalogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCatalogParamsWithTimeout creates a new CreateCatalogParams object
// with the ability to set a timeout on a request.
func NewCreateCatalogParamsWithTimeout(timeout time.Duration) *CreateCatalogParams {
	return &CreateCatalogParams{
		timeout: timeout,
	}
}

// NewCreateCatalogParamsWithContext creates a new CreateCatalogParams object
// with the ability to set a context for a request.
func NewCreateCatalogParamsWithContext(ctx context.Context) *CreateCatalogParams {
	return &CreateCatalogParams{
		Context: ctx,
	}
}

// NewCreateCatalogParamsWithHTTPClient creates a new CreateCatalogParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCatalogParamsWithHTTPClient(client *http.Client) *CreateCatalogParams {
	return &CreateCatalogParams{
		HTTPClient: client,
	}
}

/* CreateCatalogParams contains all the parameters to send to the API endpoint
   for the create catalog operation.

   Typically these are written to a http.Request.
*/
type CreateCatalogParams struct {

	// Body.
	Body *models.CatalogRequest

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCatalogParams) WithDefaults() *CreateCatalogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCatalogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create catalog params
func (o *CreateCatalogParams) WithTimeout(timeout time.Duration) *CreateCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create catalog params
func (o *CreateCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create catalog params
func (o *CreateCatalogParams) WithContext(ctx context.Context) *CreateCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create catalog params
func (o *CreateCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create catalog params
func (o *CreateCatalogParams) WithHTTPClient(client *http.Client) *CreateCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create catalog params
func (o *CreateCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create catalog params
func (o *CreateCatalogParams) WithBody(body *models.CatalogRequest) *CreateCatalogParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create catalog params
func (o *CreateCatalogParams) SetBody(body *models.CatalogRequest) {
	o.Body = body
}

// WithOwner adds the owner to the create catalog params
func (o *CreateCatalogParams) WithOwner(owner string) *CreateCatalogParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create catalog params
func (o *CreateCatalogParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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
