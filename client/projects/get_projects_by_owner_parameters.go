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
)

// NewGetProjectsByOwnerParams creates a new GetProjectsByOwnerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectsByOwnerParams() *GetProjectsByOwnerParams {
	return &GetProjectsByOwnerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsByOwnerParamsWithTimeout creates a new GetProjectsByOwnerParams object
// with the ability to set a timeout on a request.
func NewGetProjectsByOwnerParamsWithTimeout(timeout time.Duration) *GetProjectsByOwnerParams {
	return &GetProjectsByOwnerParams{
		timeout: timeout,
	}
}

// NewGetProjectsByOwnerParamsWithContext creates a new GetProjectsByOwnerParams object
// with the ability to set a context for a request.
func NewGetProjectsByOwnerParamsWithContext(ctx context.Context) *GetProjectsByOwnerParams {
	return &GetProjectsByOwnerParams{
		Context: ctx,
	}
}

// NewGetProjectsByOwnerParamsWithHTTPClient creates a new GetProjectsByOwnerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectsByOwnerParamsWithHTTPClient(client *http.Client) *GetProjectsByOwnerParams {
	return &GetProjectsByOwnerParams{
		HTTPClient: client,
	}
}

/* GetProjectsByOwnerParams contains all the parameters to send to the API endpoint
   for the get projects by owner operation.

   Typically these are written to a http.Request.
*/
type GetProjectsByOwnerParams struct {

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

// WithDefaults hydrates default values in the get projects by owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsByOwnerParams) WithDefaults() *GetProjectsByOwnerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get projects by owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsByOwnerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get projects by owner params
func (o *GetProjectsByOwnerParams) WithTimeout(timeout time.Duration) *GetProjectsByOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects by owner params
func (o *GetProjectsByOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects by owner params
func (o *GetProjectsByOwnerParams) WithContext(ctx context.Context) *GetProjectsByOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects by owner params
func (o *GetProjectsByOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects by owner params
func (o *GetProjectsByOwnerParams) WithHTTPClient(client *http.Client) *GetProjectsByOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects by owner params
func (o *GetProjectsByOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get projects by owner params
func (o *GetProjectsByOwnerParams) WithLimit(limit *string) *GetProjectsByOwnerParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get projects by owner params
func (o *GetProjectsByOwnerParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNext adds the next to the get projects by owner params
func (o *GetProjectsByOwnerParams) WithNext(next *string) *GetProjectsByOwnerParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the get projects by owner params
func (o *GetProjectsByOwnerParams) SetNext(next *string) {
	o.Next = next
}

// WithOwner adds the owner to the get projects by owner params
func (o *GetProjectsByOwnerParams) WithOwner(owner string) *GetProjectsByOwnerParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get projects by owner params
func (o *GetProjectsByOwnerParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsByOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
