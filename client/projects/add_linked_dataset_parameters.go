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

// NewAddLinkedDatasetParams creates a new AddLinkedDatasetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddLinkedDatasetParams() *AddLinkedDatasetParams {
	return &AddLinkedDatasetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddLinkedDatasetParamsWithTimeout creates a new AddLinkedDatasetParams object
// with the ability to set a timeout on a request.
func NewAddLinkedDatasetParamsWithTimeout(timeout time.Duration) *AddLinkedDatasetParams {
	return &AddLinkedDatasetParams{
		timeout: timeout,
	}
}

// NewAddLinkedDatasetParamsWithContext creates a new AddLinkedDatasetParams object
// with the ability to set a context for a request.
func NewAddLinkedDatasetParamsWithContext(ctx context.Context) *AddLinkedDatasetParams {
	return &AddLinkedDatasetParams{
		Context: ctx,
	}
}

// NewAddLinkedDatasetParamsWithHTTPClient creates a new AddLinkedDatasetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddLinkedDatasetParamsWithHTTPClient(client *http.Client) *AddLinkedDatasetParams {
	return &AddLinkedDatasetParams{
		HTTPClient: client,
	}
}

/* AddLinkedDatasetParams contains all the parameters to send to the API endpoint
   for the add linked dataset operation.

   Typically these are written to a http.Request.
*/
type AddLinkedDatasetParams struct {

	// ID.
	ID string

	// LinkedDatasetID.
	LinkedDatasetID string

	// LinkedDatasetOwner.
	LinkedDatasetOwner string

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add linked dataset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddLinkedDatasetParams) WithDefaults() *AddLinkedDatasetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add linked dataset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddLinkedDatasetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add linked dataset params
func (o *AddLinkedDatasetParams) WithTimeout(timeout time.Duration) *AddLinkedDatasetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add linked dataset params
func (o *AddLinkedDatasetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add linked dataset params
func (o *AddLinkedDatasetParams) WithContext(ctx context.Context) *AddLinkedDatasetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add linked dataset params
func (o *AddLinkedDatasetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add linked dataset params
func (o *AddLinkedDatasetParams) WithHTTPClient(client *http.Client) *AddLinkedDatasetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add linked dataset params
func (o *AddLinkedDatasetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the add linked dataset params
func (o *AddLinkedDatasetParams) WithID(id string) *AddLinkedDatasetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add linked dataset params
func (o *AddLinkedDatasetParams) SetID(id string) {
	o.ID = id
}

// WithLinkedDatasetID adds the linkedDatasetID to the add linked dataset params
func (o *AddLinkedDatasetParams) WithLinkedDatasetID(linkedDatasetID string) *AddLinkedDatasetParams {
	o.SetLinkedDatasetID(linkedDatasetID)
	return o
}

// SetLinkedDatasetID adds the linkedDatasetId to the add linked dataset params
func (o *AddLinkedDatasetParams) SetLinkedDatasetID(linkedDatasetID string) {
	o.LinkedDatasetID = linkedDatasetID
}

// WithLinkedDatasetOwner adds the linkedDatasetOwner to the add linked dataset params
func (o *AddLinkedDatasetParams) WithLinkedDatasetOwner(linkedDatasetOwner string) *AddLinkedDatasetParams {
	o.SetLinkedDatasetOwner(linkedDatasetOwner)
	return o
}

// SetLinkedDatasetOwner adds the linkedDatasetOwner to the add linked dataset params
func (o *AddLinkedDatasetParams) SetLinkedDatasetOwner(linkedDatasetOwner string) {
	o.LinkedDatasetOwner = linkedDatasetOwner
}

// WithOwner adds the owner to the add linked dataset params
func (o *AddLinkedDatasetParams) WithOwner(owner string) *AddLinkedDatasetParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the add linked dataset params
func (o *AddLinkedDatasetParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *AddLinkedDatasetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param linkedDatasetId
	if err := r.SetPathParam("linkedDatasetId", o.LinkedDatasetID); err != nil {
		return err
	}

	// path param linkedDatasetOwner
	if err := r.SetPathParam("linkedDatasetOwner", o.LinkedDatasetOwner); err != nil {
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
