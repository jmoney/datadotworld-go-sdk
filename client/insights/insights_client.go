// Code generated by go-swagger; DO NOT EDIT.

package insights

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new insights API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for insights API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateInsight(params *CreateInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInsightOK, error)

	DeleteInsight(params *DeleteInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInsightOK, error)

	GetInsight(params *GetInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInsightOK, error)

	GetInsightByVersion(params *GetInsightByVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInsightByVersionOK, error)

	GetInsightsForProject(params *GetInsightsForProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInsightsForProjectOK, error)

	ReplaceInsight(params *ReplaceInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceInsightOK, error)

	UpdateInsight(params *UpdateInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateInsightOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateInsight creates new insight for a project
*/
func (a *Client) CreateInsight(params *CreateInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInsightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInsightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createInsight",
		Method:             "POST",
		PathPattern:        "/insights/{projectOwner}/{projectId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateInsightReader{formats: a.formats},
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
	success, ok := result.(*CreateInsightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createInsight: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteInsight deletes an insight
*/
func (a *Client) DeleteInsight(params *DeleteInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInsightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInsightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteInsight",
		Method:             "DELETE",
		PathPattern:        "/insights/{projectOwner}/{projectId}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteInsightReader{formats: a.formats},
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
	success, ok := result.(*DeleteInsightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInsight: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInsight retrieves a project insight
*/
func (a *Client) GetInsight(params *GetInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInsightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInsightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInsight",
		Method:             "GET",
		PathPattern:        "/insights/{projectOwner}/{projectId}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInsightReader{formats: a.formats},
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
	success, ok := result.(*GetInsightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInsight: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInsightByVersion retrieves a project insight by version
*/
func (a *Client) GetInsightByVersion(params *GetInsightByVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInsightByVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInsightByVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInsightByVersion",
		Method:             "GET",
		PathPattern:        "/insights/{projectOwner}/{projectId}/{id}/v/{versionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInsightByVersionReader{formats: a.formats},
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
	success, ok := result.(*GetInsightByVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInsightByVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInsightsForProject gets insights for project
*/
func (a *Client) GetInsightsForProject(params *GetInsightsForProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInsightsForProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInsightsForProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInsightsForProject",
		Method:             "GET",
		PathPattern:        "/insights/{projectOwner}/{projectId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInsightsForProjectReader{formats: a.formats},
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
	success, ok := result.(*GetInsightsForProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInsightsForProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceInsight replaces an existing insight or create a new insight
*/
func (a *Client) ReplaceInsight(params *ReplaceInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceInsightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceInsightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceInsight",
		Method:             "PUT",
		PathPattern:        "/insights/{projectOwner}/{projectId}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReplaceInsightReader{formats: a.formats},
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
	success, ok := result.(*ReplaceInsightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for replaceInsight: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateInsight updates an existing insight
*/
func (a *Client) UpdateInsight(params *UpdateInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateInsightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateInsightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateInsight",
		Method:             "PATCH",
		PathPattern:        "/insights/{projectOwner}/{projectId}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateInsightReader{formats: a.formats},
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
	success, ok := result.(*UpdateInsightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateInsight: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
