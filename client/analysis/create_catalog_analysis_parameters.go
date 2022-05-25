// Code generated by go-swagger; DO NOT EDIT.

package analysis

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

// NewCreateCatalogAnalysisParams creates a new CreateCatalogAnalysisParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCatalogAnalysisParams() *CreateCatalogAnalysisParams {
	return &CreateCatalogAnalysisParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCatalogAnalysisParamsWithTimeout creates a new CreateCatalogAnalysisParams object
// with the ability to set a timeout on a request.
func NewCreateCatalogAnalysisParamsWithTimeout(timeout time.Duration) *CreateCatalogAnalysisParams {
	return &CreateCatalogAnalysisParams{
		timeout: timeout,
	}
}

// NewCreateCatalogAnalysisParamsWithContext creates a new CreateCatalogAnalysisParams object
// with the ability to set a context for a request.
func NewCreateCatalogAnalysisParamsWithContext(ctx context.Context) *CreateCatalogAnalysisParams {
	return &CreateCatalogAnalysisParams{
		Context: ctx,
	}
}

// NewCreateCatalogAnalysisParamsWithHTTPClient creates a new CreateCatalogAnalysisParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCatalogAnalysisParamsWithHTTPClient(client *http.Client) *CreateCatalogAnalysisParams {
	return &CreateCatalogAnalysisParams{
		HTTPClient: client,
	}
}

/* CreateCatalogAnalysisParams contains all the parameters to send to the API endpoint
   for the create catalog analysis operation.

   Typically these are written to a http.Request.
*/
type CreateCatalogAnalysisParams struct {

	// Body.
	Body *models.MetadataRequest

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create catalog analysis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCatalogAnalysisParams) WithDefaults() *CreateCatalogAnalysisParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create catalog analysis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCatalogAnalysisParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create catalog analysis params
func (o *CreateCatalogAnalysisParams) WithTimeout(timeout time.Duration) *CreateCatalogAnalysisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create catalog analysis params
func (o *CreateCatalogAnalysisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create catalog analysis params
func (o *CreateCatalogAnalysisParams) WithContext(ctx context.Context) *CreateCatalogAnalysisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create catalog analysis params
func (o *CreateCatalogAnalysisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create catalog analysis params
func (o *CreateCatalogAnalysisParams) WithHTTPClient(client *http.Client) *CreateCatalogAnalysisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create catalog analysis params
func (o *CreateCatalogAnalysisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create catalog analysis params
func (o *CreateCatalogAnalysisParams) WithBody(body *models.MetadataRequest) *CreateCatalogAnalysisParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create catalog analysis params
func (o *CreateCatalogAnalysisParams) SetBody(body *models.MetadataRequest) {
	o.Body = body
}

// WithOwner adds the owner to the create catalog analysis params
func (o *CreateCatalogAnalysisParams) WithOwner(owner string) *CreateCatalogAnalysisParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create catalog analysis params
func (o *CreateCatalogAnalysisParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCatalogAnalysisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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