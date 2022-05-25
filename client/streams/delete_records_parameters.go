// Code generated by go-swagger; DO NOT EDIT.

package streams

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

// NewDeleteRecordsParams creates a new DeleteRecordsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRecordsParams() *DeleteRecordsParams {
	return &DeleteRecordsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRecordsParamsWithTimeout creates a new DeleteRecordsParams object
// with the ability to set a timeout on a request.
func NewDeleteRecordsParamsWithTimeout(timeout time.Duration) *DeleteRecordsParams {
	return &DeleteRecordsParams{
		timeout: timeout,
	}
}

// NewDeleteRecordsParamsWithContext creates a new DeleteRecordsParams object
// with the ability to set a context for a request.
func NewDeleteRecordsParamsWithContext(ctx context.Context) *DeleteRecordsParams {
	return &DeleteRecordsParams{
		Context: ctx,
	}
}

// NewDeleteRecordsParamsWithHTTPClient creates a new DeleteRecordsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRecordsParamsWithHTTPClient(client *http.Client) *DeleteRecordsParams {
	return &DeleteRecordsParams{
		HTTPClient: client,
	}
}

/* DeleteRecordsParams contains all the parameters to send to the API endpoint
   for the delete records operation.

   Typically these are written to a http.Request.
*/
type DeleteRecordsParams struct {

	// ID.
	ID string

	// Owner.
	Owner string

	// StreamID.
	StreamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete records params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRecordsParams) WithDefaults() *DeleteRecordsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete records params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRecordsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete records params
func (o *DeleteRecordsParams) WithTimeout(timeout time.Duration) *DeleteRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete records params
func (o *DeleteRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete records params
func (o *DeleteRecordsParams) WithContext(ctx context.Context) *DeleteRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete records params
func (o *DeleteRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete records params
func (o *DeleteRecordsParams) WithHTTPClient(client *http.Client) *DeleteRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete records params
func (o *DeleteRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete records params
func (o *DeleteRecordsParams) WithID(id string) *DeleteRecordsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete records params
func (o *DeleteRecordsParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the delete records params
func (o *DeleteRecordsParams) WithOwner(owner string) *DeleteRecordsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete records params
func (o *DeleteRecordsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithStreamID adds the streamID to the delete records params
func (o *DeleteRecordsParams) WithStreamID(streamID string) *DeleteRecordsParams {
	o.SetStreamID(streamID)
	return o
}

// SetStreamID adds the streamId to the delete records params
func (o *DeleteRecordsParams) SetStreamID(streamID string) {
	o.StreamID = streamID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param streamId
	if err := r.SetPathParam("streamId", o.StreamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
