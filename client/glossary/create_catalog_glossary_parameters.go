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

	"datadotworld-go-sdk/models"
)

// NewCreateCatalogGlossaryParams creates a new CreateCatalogGlossaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCatalogGlossaryParams() *CreateCatalogGlossaryParams {
	return &CreateCatalogGlossaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCatalogGlossaryParamsWithTimeout creates a new CreateCatalogGlossaryParams object
// with the ability to set a timeout on a request.
func NewCreateCatalogGlossaryParamsWithTimeout(timeout time.Duration) *CreateCatalogGlossaryParams {
	return &CreateCatalogGlossaryParams{
		timeout: timeout,
	}
}

// NewCreateCatalogGlossaryParamsWithContext creates a new CreateCatalogGlossaryParams object
// with the ability to set a context for a request.
func NewCreateCatalogGlossaryParamsWithContext(ctx context.Context) *CreateCatalogGlossaryParams {
	return &CreateCatalogGlossaryParams{
		Context: ctx,
	}
}

// NewCreateCatalogGlossaryParamsWithHTTPClient creates a new CreateCatalogGlossaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCatalogGlossaryParamsWithHTTPClient(client *http.Client) *CreateCatalogGlossaryParams {
	return &CreateCatalogGlossaryParams{
		HTTPClient: client,
	}
}

/* CreateCatalogGlossaryParams contains all the parameters to send to the API endpoint
   for the create catalog glossary operation.

   Typically these are written to a http.Request.
*/
type CreateCatalogGlossaryParams struct {

	// Body.
	Body *models.CatalogGlossaryRequest

	// Owner.
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create catalog glossary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCatalogGlossaryParams) WithDefaults() *CreateCatalogGlossaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create catalog glossary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCatalogGlossaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create catalog glossary params
func (o *CreateCatalogGlossaryParams) WithTimeout(timeout time.Duration) *CreateCatalogGlossaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create catalog glossary params
func (o *CreateCatalogGlossaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create catalog glossary params
func (o *CreateCatalogGlossaryParams) WithContext(ctx context.Context) *CreateCatalogGlossaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create catalog glossary params
func (o *CreateCatalogGlossaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create catalog glossary params
func (o *CreateCatalogGlossaryParams) WithHTTPClient(client *http.Client) *CreateCatalogGlossaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create catalog glossary params
func (o *CreateCatalogGlossaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create catalog glossary params
func (o *CreateCatalogGlossaryParams) WithBody(body *models.CatalogGlossaryRequest) *CreateCatalogGlossaryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create catalog glossary params
func (o *CreateCatalogGlossaryParams) SetBody(body *models.CatalogGlossaryRequest) {
	o.Body = body
}

// WithOwner adds the owner to the create catalog glossary params
func (o *CreateCatalogGlossaryParams) WithOwner(owner string) *CreateCatalogGlossaryParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create catalog glossary params
func (o *CreateCatalogGlossaryParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCatalogGlossaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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