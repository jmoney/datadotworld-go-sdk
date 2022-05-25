// Code generated by go-swagger; DO NOT EDIT.

package relationships

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

// NewGetTableRelationshipsParams creates a new GetTableRelationshipsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTableRelationshipsParams() *GetTableRelationshipsParams {
	return &GetTableRelationshipsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTableRelationshipsParamsWithTimeout creates a new GetTableRelationshipsParams object
// with the ability to set a timeout on a request.
func NewGetTableRelationshipsParamsWithTimeout(timeout time.Duration) *GetTableRelationshipsParams {
	return &GetTableRelationshipsParams{
		timeout: timeout,
	}
}

// NewGetTableRelationshipsParamsWithContext creates a new GetTableRelationshipsParams object
// with the ability to set a context for a request.
func NewGetTableRelationshipsParamsWithContext(ctx context.Context) *GetTableRelationshipsParams {
	return &GetTableRelationshipsParams{
		Context: ctx,
	}
}

// NewGetTableRelationshipsParamsWithHTTPClient creates a new GetTableRelationshipsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTableRelationshipsParamsWithHTTPClient(client *http.Client) *GetTableRelationshipsParams {
	return &GetTableRelationshipsParams{
		HTTPClient: client,
	}
}

/* GetTableRelationshipsParams contains all the parameters to send to the API endpoint
   for the get table relationships operation.

   Typically these are written to a http.Request.
*/
type GetTableRelationshipsParams struct {

	// Body.
	Body *models.RelationshipGetTableRequest

	// Limit.
	Limit *string

	// Next.
	Next *string

	/* Owner.

	   ID of organization
	*/
	Owner string

	/* SourceID.

	   ID of dataset that the table is in
	*/
	SourceID string

	// TableID.
	TableID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get table relationships params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTableRelationshipsParams) WithDefaults() *GetTableRelationshipsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get table relationships params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTableRelationshipsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get table relationships params
func (o *GetTableRelationshipsParams) WithTimeout(timeout time.Duration) *GetTableRelationshipsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get table relationships params
func (o *GetTableRelationshipsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get table relationships params
func (o *GetTableRelationshipsParams) WithContext(ctx context.Context) *GetTableRelationshipsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get table relationships params
func (o *GetTableRelationshipsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get table relationships params
func (o *GetTableRelationshipsParams) WithHTTPClient(client *http.Client) *GetTableRelationshipsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get table relationships params
func (o *GetTableRelationshipsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get table relationships params
func (o *GetTableRelationshipsParams) WithBody(body *models.RelationshipGetTableRequest) *GetTableRelationshipsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get table relationships params
func (o *GetTableRelationshipsParams) SetBody(body *models.RelationshipGetTableRequest) {
	o.Body = body
}

// WithLimit adds the limit to the get table relationships params
func (o *GetTableRelationshipsParams) WithLimit(limit *string) *GetTableRelationshipsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get table relationships params
func (o *GetTableRelationshipsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNext adds the next to the get table relationships params
func (o *GetTableRelationshipsParams) WithNext(next *string) *GetTableRelationshipsParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the get table relationships params
func (o *GetTableRelationshipsParams) SetNext(next *string) {
	o.Next = next
}

// WithOwner adds the owner to the get table relationships params
func (o *GetTableRelationshipsParams) WithOwner(owner string) *GetTableRelationshipsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get table relationships params
func (o *GetTableRelationshipsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithSourceID adds the sourceID to the get table relationships params
func (o *GetTableRelationshipsParams) WithSourceID(sourceID string) *GetTableRelationshipsParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the get table relationships params
func (o *GetTableRelationshipsParams) SetSourceID(sourceID string) {
	o.SourceID = sourceID
}

// WithTableID adds the tableID to the get table relationships params
func (o *GetTableRelationshipsParams) WithTableID(tableID string) *GetTableRelationshipsParams {
	o.SetTableID(tableID)
	return o
}

// SetTableID adds the tableId to the get table relationships params
func (o *GetTableRelationshipsParams) SetTableID(tableID string) {
	o.TableID = tableID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTableRelationshipsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param sourceId
	if err := r.SetPathParam("sourceId", o.SourceID); err != nil {
		return err
	}

	// path param tableId
	if err := r.SetPathParam("tableId", o.TableID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}