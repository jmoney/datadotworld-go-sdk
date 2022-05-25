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
)

// NewGetRequestParams creates a new GetRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRequestParams() *GetRequestParams {
	return &GetRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRequestParamsWithTimeout creates a new GetRequestParams object
// with the ability to set a timeout on a request.
func NewGetRequestParamsWithTimeout(timeout time.Duration) *GetRequestParams {
	return &GetRequestParams{
		timeout: timeout,
	}
}

// NewGetRequestParamsWithContext creates a new GetRequestParams object
// with the ability to set a context for a request.
func NewGetRequestParamsWithContext(ctx context.Context) *GetRequestParams {
	return &GetRequestParams{
		Context: ctx,
	}
}

// NewGetRequestParamsWithHTTPClient creates a new GetRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRequestParamsWithHTTPClient(client *http.Client) *GetRequestParams {
	return &GetRequestParams{
		HTTPClient: client,
	}
}

/* GetRequestParams contains all the parameters to send to the API endpoint
   for the get request operation.

   Typically these are written to a http.Request.
*/
type GetRequestParams struct {

	// Requestid.
	Requestid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRequestParams) WithDefaults() *GetRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get request params
func (o *GetRequestParams) WithTimeout(timeout time.Duration) *GetRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get request params
func (o *GetRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get request params
func (o *GetRequestParams) WithContext(ctx context.Context) *GetRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get request params
func (o *GetRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get request params
func (o *GetRequestParams) WithHTTPClient(client *http.Client) *GetRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get request params
func (o *GetRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestid adds the requestid to the get request params
func (o *GetRequestParams) WithRequestid(requestid string) *GetRequestParams {
	o.SetRequestid(requestid)
	return o
}

// SetRequestid adds the requestid to the get request params
func (o *GetRequestParams) SetRequestid(requestid string) {
	o.Requestid = requestid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param requestid
	if err := r.SetPathParam("requestid", o.Requestid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
