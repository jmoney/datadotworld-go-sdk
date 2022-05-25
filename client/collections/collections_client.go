// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new collections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for collections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCatalog(params *CreateCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCatalogOK, error)

	DeleteCatalog(params *DeleteCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCatalogOK, error)

	GetCatalog(params *GetCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCatalogOK, error)

	GetCatalogs(params *GetCatalogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCatalogsOK, error)

	ReplaceCatalog(params *ReplaceCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceCatalogOK, error)

	UpdateCatalog(params *UpdateCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCatalogOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCatalog creates a new collection
*/
func (a *Client) CreateCatalog(params *CreateCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createCatalog",
		Method:             "POST",
		PathPattern:        "/metadata/collections/{owner}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCatalogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCatalog deletes a collection
*/
func (a *Client) DeleteCatalog(params *DeleteCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteCatalog",
		Method:             "DELETE",
		PathPattern:        "/metadata/collections/{owner}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCatalogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCatalog gets a collection by ID
*/
func (a *Client) GetCatalog(params *GetCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCatalog",
		Method:             "GET",
		PathPattern:        "/metadata/collections/{owner}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCatalogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCatalogs gets collections owned by specified owner
*/
func (a *Client) GetCatalogs(params *GetCatalogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCatalogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCatalogs",
		Method:             "GET",
		PathPattern:        "/metadata/collections/{owner}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCatalogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCatalogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceCatalog replaces a collection
*/
func (a *Client) ReplaceCatalog(params *ReplaceCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceCatalog",
		Method:             "PUT",
		PathPattern:        "/metadata/collections/{owner}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReplaceCatalogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReplaceCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for replaceCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCatalog updates a collection
*/
func (a *Client) UpdateCatalog(params *UpdateCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateCatalog",
		Method:             "PATCH",
		PathPattern:        "/metadata/collections/{owner}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCatalogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}