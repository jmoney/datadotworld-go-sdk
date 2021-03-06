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

// NewDeleteDatasetSavedQueryParams creates a new DeleteDatasetSavedQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDatasetSavedQueryParams() *DeleteDatasetSavedQueryParams {
	return &DeleteDatasetSavedQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDatasetSavedQueryParamsWithTimeout creates a new DeleteDatasetSavedQueryParams object
// with the ability to set a timeout on a request.
func NewDeleteDatasetSavedQueryParamsWithTimeout(timeout time.Duration) *DeleteDatasetSavedQueryParams {
	return &DeleteDatasetSavedQueryParams{
		timeout: timeout,
	}
}

// NewDeleteDatasetSavedQueryParamsWithContext creates a new DeleteDatasetSavedQueryParams object
// with the ability to set a context for a request.
func NewDeleteDatasetSavedQueryParamsWithContext(ctx context.Context) *DeleteDatasetSavedQueryParams {
	return &DeleteDatasetSavedQueryParams{
		Context: ctx,
	}
}

// NewDeleteDatasetSavedQueryParamsWithHTTPClient creates a new DeleteDatasetSavedQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDatasetSavedQueryParamsWithHTTPClient(client *http.Client) *DeleteDatasetSavedQueryParams {
	return &DeleteDatasetSavedQueryParams{
		HTTPClient: client,
	}
}

/* DeleteDatasetSavedQueryParams contains all the parameters to send to the API endpoint
   for the delete dataset saved query operation.

   Typically these are written to a http.Request.
*/
type DeleteDatasetSavedQueryParams struct {

	// ID.
	ID string

	// Owner.
	Owner string

	// QueryID.
	QueryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete dataset saved query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDatasetSavedQueryParams) WithDefaults() *DeleteDatasetSavedQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete dataset saved query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDatasetSavedQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) WithTimeout(timeout time.Duration) *DeleteDatasetSavedQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) WithContext(ctx context.Context) *DeleteDatasetSavedQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) WithHTTPClient(client *http.Client) *DeleteDatasetSavedQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) WithID(id string) *DeleteDatasetSavedQueryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) WithOwner(owner string) *DeleteDatasetSavedQueryParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithQueryID adds the queryID to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) WithQueryID(queryID string) *DeleteDatasetSavedQueryParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the delete dataset saved query params
func (o *DeleteDatasetSavedQueryParams) SetQueryID(queryID string) {
	o.QueryID = queryID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDatasetSavedQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param queryId
	if err := r.SetPathParam("queryId", o.QueryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
