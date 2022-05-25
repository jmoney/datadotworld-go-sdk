// Code generated by go-swagger; DO NOT EDIT.

package requests

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

// NewAcceptRequestParams creates a new AcceptRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAcceptRequestParams() *AcceptRequestParams {
	return &AcceptRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptRequestParamsWithTimeout creates a new AcceptRequestParams object
// with the ability to set a timeout on a request.
func NewAcceptRequestParamsWithTimeout(timeout time.Duration) *AcceptRequestParams {
	return &AcceptRequestParams{
		timeout: timeout,
	}
}

// NewAcceptRequestParamsWithContext creates a new AcceptRequestParams object
// with the ability to set a context for a request.
func NewAcceptRequestParamsWithContext(ctx context.Context) *AcceptRequestParams {
	return &AcceptRequestParams{
		Context: ctx,
	}
}

// NewAcceptRequestParamsWithHTTPClient creates a new AcceptRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewAcceptRequestParamsWithHTTPClient(client *http.Client) *AcceptRequestParams {
	return &AcceptRequestParams{
		HTTPClient: client,
	}
}

/* AcceptRequestParams contains all the parameters to send to the API endpoint
   for the accept request operation.

   Typically these are written to a http.Request.
*/
type AcceptRequestParams struct {

	// Body.
	Body *models.AcceptRequestDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accept request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptRequestParams) WithDefaults() *AcceptRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accept request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accept request params
func (o *AcceptRequestParams) WithTimeout(timeout time.Duration) *AcceptRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept request params
func (o *AcceptRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept request params
func (o *AcceptRequestParams) WithContext(ctx context.Context) *AcceptRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept request params
func (o *AcceptRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept request params
func (o *AcceptRequestParams) WithHTTPClient(client *http.Client) *AcceptRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept request params
func (o *AcceptRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the accept request params
func (o *AcceptRequestParams) WithBody(body *models.AcceptRequestDto) *AcceptRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the accept request params
func (o *AcceptRequestParams) SetBody(body *models.AcceptRequestDto) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
