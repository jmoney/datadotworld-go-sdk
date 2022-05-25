// Code generated by go-swagger; DO NOT EDIT.

package insights

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

// NewCreateInsightParams creates a new CreateInsightParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInsightParams() *CreateInsightParams {
	return &CreateInsightParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInsightParamsWithTimeout creates a new CreateInsightParams object
// with the ability to set a timeout on a request.
func NewCreateInsightParamsWithTimeout(timeout time.Duration) *CreateInsightParams {
	return &CreateInsightParams{
		timeout: timeout,
	}
}

// NewCreateInsightParamsWithContext creates a new CreateInsightParams object
// with the ability to set a context for a request.
func NewCreateInsightParamsWithContext(ctx context.Context) *CreateInsightParams {
	return &CreateInsightParams{
		Context: ctx,
	}
}

// NewCreateInsightParamsWithHTTPClient creates a new CreateInsightParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInsightParamsWithHTTPClient(client *http.Client) *CreateInsightParams {
	return &CreateInsightParams{
		HTTPClient: client,
	}
}

/* CreateInsightParams contains all the parameters to send to the API endpoint
   for the create insight operation.

   Typically these are written to a http.Request.
*/
type CreateInsightParams struct {

	// Body.
	Body *models.InsightCreateRequest

	// ProjectID.
	ProjectID string

	// ProjectOwner.
	ProjectOwner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create insight params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInsightParams) WithDefaults() *CreateInsightParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create insight params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInsightParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create insight params
func (o *CreateInsightParams) WithTimeout(timeout time.Duration) *CreateInsightParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create insight params
func (o *CreateInsightParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create insight params
func (o *CreateInsightParams) WithContext(ctx context.Context) *CreateInsightParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create insight params
func (o *CreateInsightParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create insight params
func (o *CreateInsightParams) WithHTTPClient(client *http.Client) *CreateInsightParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create insight params
func (o *CreateInsightParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create insight params
func (o *CreateInsightParams) WithBody(body *models.InsightCreateRequest) *CreateInsightParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create insight params
func (o *CreateInsightParams) SetBody(body *models.InsightCreateRequest) {
	o.Body = body
}

// WithProjectID adds the projectID to the create insight params
func (o *CreateInsightParams) WithProjectID(projectID string) *CreateInsightParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create insight params
func (o *CreateInsightParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithProjectOwner adds the projectOwner to the create insight params
func (o *CreateInsightParams) WithProjectOwner(projectOwner string) *CreateInsightParams {
	o.SetProjectOwner(projectOwner)
	return o
}

// SetProjectOwner adds the projectOwner to the create insight params
func (o *CreateInsightParams) SetProjectOwner(projectOwner string) {
	o.ProjectOwner = projectOwner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInsightParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	// path param projectOwner
	if err := r.SetPathParam("projectOwner", o.ProjectOwner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}