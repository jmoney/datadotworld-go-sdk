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

// NewReplaceCatalogAnalysisParams creates a new ReplaceCatalogAnalysisParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceCatalogAnalysisParams() *ReplaceCatalogAnalysisParams {
	return &ReplaceCatalogAnalysisParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceCatalogAnalysisParamsWithTimeout creates a new ReplaceCatalogAnalysisParams object
// with the ability to set a timeout on a request.
func NewReplaceCatalogAnalysisParamsWithTimeout(timeout time.Duration) *ReplaceCatalogAnalysisParams {
	return &ReplaceCatalogAnalysisParams{
		timeout: timeout,
	}
}

// NewReplaceCatalogAnalysisParamsWithContext creates a new ReplaceCatalogAnalysisParams object
// with the ability to set a context for a request.
func NewReplaceCatalogAnalysisParamsWithContext(ctx context.Context) *ReplaceCatalogAnalysisParams {
	return &ReplaceCatalogAnalysisParams{
		Context: ctx,
	}
}

// NewReplaceCatalogAnalysisParamsWithHTTPClient creates a new ReplaceCatalogAnalysisParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceCatalogAnalysisParamsWithHTTPClient(client *http.Client) *ReplaceCatalogAnalysisParams {
	return &ReplaceCatalogAnalysisParams{
		HTTPClient: client,
	}
}

/* ReplaceCatalogAnalysisParams contains all the parameters to send to the API endpoint
   for the replace catalog analysis operation.

   Typically these are written to a http.Request.
*/
type ReplaceCatalogAnalysisParams struct {

	// Body.
	Body *models.MetadataRequest

	// ID.
	ID string

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace catalog analysis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceCatalogAnalysisParams) WithDefaults() *ReplaceCatalogAnalysisParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace catalog analysis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceCatalogAnalysisParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) WithTimeout(timeout time.Duration) *ReplaceCatalogAnalysisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) WithContext(ctx context.Context) *ReplaceCatalogAnalysisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) WithHTTPClient(client *http.Client) *ReplaceCatalogAnalysisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) WithBody(body *models.MetadataRequest) *ReplaceCatalogAnalysisParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) SetBody(body *models.MetadataRequest) {
	o.Body = body
}

// WithID adds the id to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) WithID(id string) *ReplaceCatalogAnalysisParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) WithOwner(owner string) *ReplaceCatalogAnalysisParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the replace catalog analysis params
func (o *ReplaceCatalogAnalysisParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceCatalogAnalysisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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