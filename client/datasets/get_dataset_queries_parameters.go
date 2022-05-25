// Code generated by go-swagger; DO NOT EDIT.

package datasets

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

// NewGetDatasetQueriesParams creates a new GetDatasetQueriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDatasetQueriesParams() *GetDatasetQueriesParams {
	return &GetDatasetQueriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatasetQueriesParamsWithTimeout creates a new GetDatasetQueriesParams object
// with the ability to set a timeout on a request.
func NewGetDatasetQueriesParamsWithTimeout(timeout time.Duration) *GetDatasetQueriesParams {
	return &GetDatasetQueriesParams{
		timeout: timeout,
	}
}

// NewGetDatasetQueriesParamsWithContext creates a new GetDatasetQueriesParams object
// with the ability to set a context for a request.
func NewGetDatasetQueriesParamsWithContext(ctx context.Context) *GetDatasetQueriesParams {
	return &GetDatasetQueriesParams{
		Context: ctx,
	}
}

// NewGetDatasetQueriesParamsWithHTTPClient creates a new GetDatasetQueriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDatasetQueriesParamsWithHTTPClient(client *http.Client) *GetDatasetQueriesParams {
	return &GetDatasetQueriesParams{
		HTTPClient: client,
	}
}

/* GetDatasetQueriesParams contains all the parameters to send to the API endpoint
   for the get dataset queries operation.

   Typically these are written to a http.Request.
*/
type GetDatasetQueriesParams struct {

	// ID.
	ID string

	// Limit.
	Limit *string

	// Next.
	Next *string

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dataset queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasetQueriesParams) WithDefaults() *GetDatasetQueriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dataset queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasetQueriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dataset queries params
func (o *GetDatasetQueriesParams) WithTimeout(timeout time.Duration) *GetDatasetQueriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dataset queries params
func (o *GetDatasetQueriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dataset queries params
func (o *GetDatasetQueriesParams) WithContext(ctx context.Context) *GetDatasetQueriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dataset queries params
func (o *GetDatasetQueriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dataset queries params
func (o *GetDatasetQueriesParams) WithHTTPClient(client *http.Client) *GetDatasetQueriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dataset queries params
func (o *GetDatasetQueriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get dataset queries params
func (o *GetDatasetQueriesParams) WithID(id string) *GetDatasetQueriesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get dataset queries params
func (o *GetDatasetQueriesParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the get dataset queries params
func (o *GetDatasetQueriesParams) WithLimit(limit *string) *GetDatasetQueriesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get dataset queries params
func (o *GetDatasetQueriesParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNext adds the next to the get dataset queries params
func (o *GetDatasetQueriesParams) WithNext(next *string) *GetDatasetQueriesParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the get dataset queries params
func (o *GetDatasetQueriesParams) SetNext(next *string) {
	o.Next = next
}

// WithOwner adds the owner to the get dataset queries params
func (o *GetDatasetQueriesParams) WithOwner(owner string) *GetDatasetQueriesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get dataset queries params
func (o *GetDatasetQueriesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatasetQueriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
