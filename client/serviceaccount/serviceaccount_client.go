// Code generated by go-swagger; DO NOT EDIT.

package serviceaccount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new serviceaccount API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for serviceaccount API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateServiceAccount(params *CreateServiceAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateServiceAccountOK, error)

	DeleteServiceAccount(params *DeleteServiceAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServiceAccountOK, error)

	GetServiceAccounts(params *GetServiceAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceAccountsOK, error)

	ResetServiceAccount(params *ResetServiceAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetServiceAccountOK, error)

	UpdateServiceAccount(params *UpdateServiceAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateServiceAccountOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateServiceAccount creates a new service account
*/
func (a *Client) CreateServiceAccount(params *CreateServiceAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateServiceAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServiceAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createServiceAccount",
		Method:             "POST",
		PathPattern:        "/serviceaccount/{owner}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateServiceAccountReader{formats: a.formats},
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
	success, ok := result.(*CreateServiceAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createServiceAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteServiceAccount deletes service account
*/
func (a *Client) DeleteServiceAccount(params *DeleteServiceAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServiceAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteServiceAccount",
		Method:             "DELETE",
		PathPattern:        "/serviceaccount/{owner}/{serviceAccount}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServiceAccountReader{formats: a.formats},
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
	success, ok := result.(*DeleteServiceAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteServiceAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetServiceAccounts gets all service accounts by owner
*/
func (a *Client) GetServiceAccounts(params *GetServiceAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getServiceAccounts",
		Method:             "GET",
		PathPattern:        "/serviceaccount/{owner}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceAccountsReader{formats: a.formats},
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
	success, ok := result.(*GetServiceAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServiceAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ResetServiceAccount refreshes service account token
*/
func (a *Client) ResetServiceAccount(params *ResetServiceAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetServiceAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetServiceAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetServiceAccount",
		Method:             "POST",
		PathPattern:        "/serviceaccount/{owner}/{serviceAccount}/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResetServiceAccountReader{formats: a.formats},
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
	success, ok := result.(*ResetServiceAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resetServiceAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateServiceAccount updates service account
*/
func (a *Client) UpdateServiceAccount(params *UpdateServiceAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateServiceAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServiceAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateServiceAccount",
		Method:             "PATCH",
		PathPattern:        "/serviceaccount/{owner}/{serviceAccount}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServiceAccountReader{formats: a.formats},
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
	success, ok := result.(*UpdateServiceAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateServiceAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
