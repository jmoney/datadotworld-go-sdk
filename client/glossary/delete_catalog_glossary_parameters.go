// Code generated by go-swagger; DO NOT EDIT.

package glossary

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

// NewDeleteCatalogGlossaryParams creates a new DeleteCatalogGlossaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCatalogGlossaryParams() *DeleteCatalogGlossaryParams {
	return &DeleteCatalogGlossaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCatalogGlossaryParamsWithTimeout creates a new DeleteCatalogGlossaryParams object
// with the ability to set a timeout on a request.
func NewDeleteCatalogGlossaryParamsWithTimeout(timeout time.Duration) *DeleteCatalogGlossaryParams {
	return &DeleteCatalogGlossaryParams{
		timeout: timeout,
	}
}

// NewDeleteCatalogGlossaryParamsWithContext creates a new DeleteCatalogGlossaryParams object
// with the ability to set a context for a request.
func NewDeleteCatalogGlossaryParamsWithContext(ctx context.Context) *DeleteCatalogGlossaryParams {
	return &DeleteCatalogGlossaryParams{
		Context: ctx,
	}
}

// NewDeleteCatalogGlossaryParamsWithHTTPClient creates a new DeleteCatalogGlossaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCatalogGlossaryParamsWithHTTPClient(client *http.Client) *DeleteCatalogGlossaryParams {
	return &DeleteCatalogGlossaryParams{
		HTTPClient: client,
	}
}

/* DeleteCatalogGlossaryParams contains all the parameters to send to the API endpoint
   for the delete catalog glossary operation.

   Typically these are written to a http.Request.
*/
type DeleteCatalogGlossaryParams struct {

	// ID.
	ID string

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete catalog glossary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCatalogGlossaryParams) WithDefaults() *DeleteCatalogGlossaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete catalog glossary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCatalogGlossaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete catalog glossary params
func (o *DeleteCatalogGlossaryParams) WithTimeout(timeout time.Duration) *DeleteCatalogGlossaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete catalog glossary params
func (o *DeleteCatalogGlossaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete catalog glossary params
func (o *DeleteCatalogGlossaryParams) WithContext(ctx context.Context) *DeleteCatalogGlossaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete catalog glossary params
func (o *DeleteCatalogGlossaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete catalog glossary params
func (o *DeleteCatalogGlossaryParams) WithHTTPClient(client *http.Client) *DeleteCatalogGlossaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete catalog glossary params
func (o *DeleteCatalogGlossaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete catalog glossary params
func (o *DeleteCatalogGlossaryParams) WithID(id string) *DeleteCatalogGlossaryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete catalog glossary params
func (o *DeleteCatalogGlossaryParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the delete catalog glossary params
func (o *DeleteCatalogGlossaryParams) WithOwner(owner string) *DeleteCatalogGlossaryParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete catalog glossary params
func (o *DeleteCatalogGlossaryParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCatalogGlossaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
