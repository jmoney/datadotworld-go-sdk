// Code generated by go-swagger; DO NOT EDIT.

package search

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
	"github.com/go-openapi/swag"

	"datadotworld-go-sdk/models"
)

// NewSearchResourcesParams creates a new SearchResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchResourcesParams() *SearchResourcesParams {
	return &SearchResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchResourcesParamsWithTimeout creates a new SearchResourcesParams object
// with the ability to set a timeout on a request.
func NewSearchResourcesParamsWithTimeout(timeout time.Duration) *SearchResourcesParams {
	return &SearchResourcesParams{
		timeout: timeout,
	}
}

// NewSearchResourcesParamsWithContext creates a new SearchResourcesParams object
// with the ability to set a context for a request.
func NewSearchResourcesParamsWithContext(ctx context.Context) *SearchResourcesParams {
	return &SearchResourcesParams{
		Context: ctx,
	}
}

// NewSearchResourcesParamsWithHTTPClient creates a new SearchResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchResourcesParamsWithHTTPClient(client *http.Client) *SearchResourcesParams {
	return &SearchResourcesParams{
		HTTPClient: client,
	}
}

/* SearchResourcesParams contains all the parameters to send to the API endpoint
   for the search resources operation.

   Typically these are written to a http.Request.
*/
type SearchResourcesParams struct {

	// Body.
	Body *models.SearchRequest

	// From.
	From *string

	// Hydration.
	Hydration []string

	// Size.
	Size *string

	// Sort.
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchResourcesParams) WithDefaults() *SearchResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchResourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search resources params
func (o *SearchResourcesParams) WithTimeout(timeout time.Duration) *SearchResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search resources params
func (o *SearchResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search resources params
func (o *SearchResourcesParams) WithContext(ctx context.Context) *SearchResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search resources params
func (o *SearchResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search resources params
func (o *SearchResourcesParams) WithHTTPClient(client *http.Client) *SearchResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search resources params
func (o *SearchResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search resources params
func (o *SearchResourcesParams) WithBody(body *models.SearchRequest) *SearchResourcesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search resources params
func (o *SearchResourcesParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithFrom adds the from to the search resources params
func (o *SearchResourcesParams) WithFrom(from *string) *SearchResourcesParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the search resources params
func (o *SearchResourcesParams) SetFrom(from *string) {
	o.From = from
}

// WithHydration adds the hydration to the search resources params
func (o *SearchResourcesParams) WithHydration(hydration []string) *SearchResourcesParams {
	o.SetHydration(hydration)
	return o
}

// SetHydration adds the hydration to the search resources params
func (o *SearchResourcesParams) SetHydration(hydration []string) {
	o.Hydration = hydration
}

// WithSize adds the size to the search resources params
func (o *SearchResourcesParams) WithSize(size *string) *SearchResourcesParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search resources params
func (o *SearchResourcesParams) SetSize(size *string) {
	o.Size = size
}

// WithSort adds the sort to the search resources params
func (o *SearchResourcesParams) WithSort(sort *string) *SearchResourcesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search resources params
func (o *SearchResourcesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *SearchResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.From != nil {

		// query param from
		var qrFrom string

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	if o.Hydration != nil {

		// binding items for hydration
		joinedHydration := o.bindParamHydration(reg)

		// query array param hydration
		if err := r.SetQueryParam("hydration", joinedHydration...); err != nil {
			return err
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize string

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := qrSize
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchResources binds the parameter hydration
func (o *SearchResourcesParams) bindParamHydration(formats strfmt.Registry) []string {
	hydrationIR := o.Hydration

	var hydrationIC []string
	for _, hydrationIIR := range hydrationIR { // explode []string

		hydrationIIV := hydrationIIR // string as string
		hydrationIC = append(hydrationIC, hydrationIIV)
	}

	// items.CollectionFormat: "multi"
	hydrationIS := swag.JoinByFormat(hydrationIC, "multi")

	return hydrationIS
}
