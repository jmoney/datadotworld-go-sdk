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

// NewFetchContributingDatasetsParams creates a new FetchContributingDatasetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFetchContributingDatasetsParams() *FetchContributingDatasetsParams {
	return &FetchContributingDatasetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFetchContributingDatasetsParamsWithTimeout creates a new FetchContributingDatasetsParams object
// with the ability to set a timeout on a request.
func NewFetchContributingDatasetsParamsWithTimeout(timeout time.Duration) *FetchContributingDatasetsParams {
	return &FetchContributingDatasetsParams{
		timeout: timeout,
	}
}

// NewFetchContributingDatasetsParamsWithContext creates a new FetchContributingDatasetsParams object
// with the ability to set a context for a request.
func NewFetchContributingDatasetsParamsWithContext(ctx context.Context) *FetchContributingDatasetsParams {
	return &FetchContributingDatasetsParams{
		Context: ctx,
	}
}

// NewFetchContributingDatasetsParamsWithHTTPClient creates a new FetchContributingDatasetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFetchContributingDatasetsParamsWithHTTPClient(client *http.Client) *FetchContributingDatasetsParams {
	return &FetchContributingDatasetsParams{
		HTTPClient: client,
	}
}

/* FetchContributingDatasetsParams contains all the parameters to send to the API endpoint
   for the fetch contributing datasets operation.

   Typically these are written to a http.Request.
*/
type FetchContributingDatasetsParams struct {

	// Fields.
	Fields *string

	// Limit.
	Limit *string

	// Next.
	Next *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fetch contributing datasets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FetchContributingDatasetsParams) WithDefaults() *FetchContributingDatasetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fetch contributing datasets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FetchContributingDatasetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) WithTimeout(timeout time.Duration) *FetchContributingDatasetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) WithContext(ctx context.Context) *FetchContributingDatasetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) WithHTTPClient(client *http.Client) *FetchContributingDatasetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) WithFields(fields *string) *FetchContributingDatasetsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) WithLimit(limit *string) *FetchContributingDatasetsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNext adds the next to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) WithNext(next *string) *FetchContributingDatasetsParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the fetch contributing datasets params
func (o *FetchContributingDatasetsParams) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *FetchContributingDatasetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Next != nil {

		// query param next
		var qrNext string

		if o.Next != nil {
			qrNext = *o.Next
		}
		qNext := qrNext
		if qNext != "" {

			if err := r.SetQueryParam("next", qNext); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
