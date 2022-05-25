// Code generated by go-swagger; DO NOT EDIT.

package properties

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

// NewGetPropertiesByOwnerParams creates a new GetPropertiesByOwnerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPropertiesByOwnerParams() *GetPropertiesByOwnerParams {
	return &GetPropertiesByOwnerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPropertiesByOwnerParamsWithTimeout creates a new GetPropertiesByOwnerParams object
// with the ability to set a timeout on a request.
func NewGetPropertiesByOwnerParamsWithTimeout(timeout time.Duration) *GetPropertiesByOwnerParams {
	return &GetPropertiesByOwnerParams{
		timeout: timeout,
	}
}

// NewGetPropertiesByOwnerParamsWithContext creates a new GetPropertiesByOwnerParams object
// with the ability to set a context for a request.
func NewGetPropertiesByOwnerParamsWithContext(ctx context.Context) *GetPropertiesByOwnerParams {
	return &GetPropertiesByOwnerParams{
		Context: ctx,
	}
}

// NewGetPropertiesByOwnerParamsWithHTTPClient creates a new GetPropertiesByOwnerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPropertiesByOwnerParamsWithHTTPClient(client *http.Client) *GetPropertiesByOwnerParams {
	return &GetPropertiesByOwnerParams{
		HTTPClient: client,
	}
}

/* GetPropertiesByOwnerParams contains all the parameters to send to the API endpoint
   for the get properties by owner operation.

   Typically these are written to a http.Request.
*/
type GetPropertiesByOwnerParams struct {

	// Categorylabel.
	Categorylabel *string

	// Owner.
	Owner string

	// Typelabel.
	Typelabel *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get properties by owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPropertiesByOwnerParams) WithDefaults() *GetPropertiesByOwnerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get properties by owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPropertiesByOwnerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get properties by owner params
func (o *GetPropertiesByOwnerParams) WithTimeout(timeout time.Duration) *GetPropertiesByOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get properties by owner params
func (o *GetPropertiesByOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get properties by owner params
func (o *GetPropertiesByOwnerParams) WithContext(ctx context.Context) *GetPropertiesByOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get properties by owner params
func (o *GetPropertiesByOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get properties by owner params
func (o *GetPropertiesByOwnerParams) WithHTTPClient(client *http.Client) *GetPropertiesByOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get properties by owner params
func (o *GetPropertiesByOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategorylabel adds the categorylabel to the get properties by owner params
func (o *GetPropertiesByOwnerParams) WithCategorylabel(categorylabel *string) *GetPropertiesByOwnerParams {
	o.SetCategorylabel(categorylabel)
	return o
}

// SetCategorylabel adds the categorylabel to the get properties by owner params
func (o *GetPropertiesByOwnerParams) SetCategorylabel(categorylabel *string) {
	o.Categorylabel = categorylabel
}

// WithOwner adds the owner to the get properties by owner params
func (o *GetPropertiesByOwnerParams) WithOwner(owner string) *GetPropertiesByOwnerParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get properties by owner params
func (o *GetPropertiesByOwnerParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithTypelabel adds the typelabel to the get properties by owner params
func (o *GetPropertiesByOwnerParams) WithTypelabel(typelabel *string) *GetPropertiesByOwnerParams {
	o.SetTypelabel(typelabel)
	return o
}

// SetTypelabel adds the typelabel to the get properties by owner params
func (o *GetPropertiesByOwnerParams) SetTypelabel(typelabel *string) {
	o.Typelabel = typelabel
}

// WriteToRequest writes these params to a swagger request
func (o *GetPropertiesByOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Categorylabel != nil {

		// query param categorylabel
		var qrCategorylabel string

		if o.Categorylabel != nil {
			qrCategorylabel = *o.Categorylabel
		}
		qCategorylabel := qrCategorylabel
		if qCategorylabel != "" {

			if err := r.SetQueryParam("categorylabel", qCategorylabel); err != nil {
				return err
			}
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Typelabel != nil {

		// query param typelabel
		var qrTypelabel string

		if o.Typelabel != nil {
			qrTypelabel = *o.Typelabel
		}
		qTypelabel := qrTypelabel
		if qTypelabel != "" {

			if err := r.SetQueryParam("typelabel", qTypelabel); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
