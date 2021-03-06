// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	FetchContributingDatasets(params *FetchContributingDatasetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchContributingDatasetsOK, error)

	FetchContributingProjects(params *FetchContributingProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchContributingProjectsOK, error)

	FetchDatasets(params *FetchDatasetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchDatasetsOK, error)

	FetchLikedDatasets(params *FetchLikedDatasetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchLikedDatasetsOK, error)

	FetchLikedProjects(params *FetchLikedProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchLikedProjectsOK, error)

	FetchOrganizations(params *FetchOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchOrganizationsOK, error)

	FetchProjects(params *FetchProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchProjectsOK, error)

	GetForDataset(params *GetForDatasetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetForDatasetOK, error)

	GetForProject(params *GetForProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetForProjectOK, error)

	GetForUser(params *GetForUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetForUserOK, error)

	GetUserData(params *GetUserDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserDataOK, error)

	GetWebhooks(params *GetWebhooksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebhooksOK, error)

	SubscribeToDataset(params *SubscribeToDatasetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscribeToDatasetOK, error)

	SubscribeToProject(params *SubscribeToProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscribeToProjectOK, error)

	SubscribeToUser(params *SubscribeToUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscribeToUserOK, error)

	UnsubscribeFromDataset(params *UnsubscribeFromDatasetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnsubscribeFromDatasetOK, error)

	UnsubscribeFromProject(params *UnsubscribeFromProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnsubscribeFromProjectOK, error)

	UnsubscribeFromUser(params *UnsubscribeFromUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnsubscribeFromUserOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  FetchContributingDatasets fetches datasets to which the requesting user is a contributor
*/
func (a *Client) FetchContributingDatasets(params *FetchContributingDatasetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchContributingDatasetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchContributingDatasetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchContributingDatasets",
		Method:             "GET",
		PathPattern:        "/user/datasets/contributing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchContributingDatasetsReader{formats: a.formats},
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
	success, ok := result.(*FetchContributingDatasetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fetchContributingDatasets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FetchContributingProjects fetches projects to which the requesting user is a contributor
*/
func (a *Client) FetchContributingProjects(params *FetchContributingProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchContributingProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchContributingProjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchContributingProjects",
		Method:             "GET",
		PathPattern:        "/user/projects/contributing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchContributingProjectsReader{formats: a.formats},
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
	success, ok := result.(*FetchContributingProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fetchContributingProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FetchDatasets fetches datasets the requesting user owns
*/
func (a *Client) FetchDatasets(params *FetchDatasetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchDatasetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchDatasetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchDatasets",
		Method:             "GET",
		PathPattern:        "/user/datasets/own",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchDatasetsReader{formats: a.formats},
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
	success, ok := result.(*FetchDatasetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fetchDatasets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FetchLikedDatasets fetches datasets the requesting user likes
*/
func (a *Client) FetchLikedDatasets(params *FetchLikedDatasetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchLikedDatasetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchLikedDatasetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchLikedDatasets",
		Method:             "GET",
		PathPattern:        "/user/datasets/liked",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchLikedDatasetsReader{formats: a.formats},
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
	success, ok := result.(*FetchLikedDatasetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fetchLikedDatasets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FetchLikedProjects fetches projects the requesting user likes
*/
func (a *Client) FetchLikedProjects(params *FetchLikedProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchLikedProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchLikedProjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchLikedProjects",
		Method:             "GET",
		PathPattern:        "/user/projects/liked",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchLikedProjectsReader{formats: a.formats},
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
	success, ok := result.(*FetchLikedProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fetchLikedProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FetchOrganizations fetches organizations to which the requesting user is a member
*/
func (a *Client) FetchOrganizations(params *FetchOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchOrganizationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchOrganizations",
		Method:             "GET",
		PathPattern:        "/user/organizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchOrganizationsReader{formats: a.formats},
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
	success, ok := result.(*FetchOrganizationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fetchOrganizations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FetchProjects fetches projects the requesting user owns
*/
func (a *Client) FetchProjects(params *FetchProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchProjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchProjects",
		Method:             "GET",
		PathPattern:        "/user/projects/own",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchProjectsReader{formats: a.formats},
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
	success, ok := result.(*FetchProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fetchProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetForDataset fetches a webhook subscription for a dataset
*/
func (a *Client) GetForDataset(params *GetForDatasetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetForDatasetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetForDatasetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getForDataset",
		Method:             "GET",
		PathPattern:        "/user/webhooks/datasets/{owner}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetForDatasetReader{formats: a.formats},
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
	success, ok := result.(*GetForDatasetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getForDataset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetForProject fetches a webhook subscription for a project
*/
func (a *Client) GetForProject(params *GetForProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetForProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetForProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getForProject",
		Method:             "GET",
		PathPattern:        "/user/webhooks/projects/{owner}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetForProjectReader{formats: a.formats},
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
	success, ok := result.(*GetForProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getForProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetForUser fetches a webhook subscription for a user account
*/
func (a *Client) GetForUser(params *GetForUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetForUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetForUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getForUser",
		Method:             "GET",
		PathPattern:        "/user/webhooks/users/{account}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetForUserReader{formats: a.formats},
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
	success, ok := result.(*GetForUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getForUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUserData gets user data
*/
func (a *Client) GetUserData(params *GetUserDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserDataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserData",
		Method:             "GET",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserDataReader{formats: a.formats},
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
	success, ok := result.(*GetUserDataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserData: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWebhooks returns existing webhook subscriptions
*/
func (a *Client) GetWebhooks(params *GetWebhooksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWebhooksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWebhooksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getWebhooks",
		Method:             "GET",
		PathPattern:        "/user/webhooks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWebhooksReader{formats: a.formats},
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
	success, ok := result.(*GetWebhooksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWebhooks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SubscribeToDataset creates a webhook subscription for a dataset including objects owned by it
*/
func (a *Client) SubscribeToDataset(params *SubscribeToDatasetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscribeToDatasetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscribeToDatasetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "subscribeToDataset",
		Method:             "PUT",
		PathPattern:        "/user/webhooks/datasets/{owner}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscribeToDatasetReader{formats: a.formats},
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
	success, ok := result.(*SubscribeToDatasetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for subscribeToDataset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SubscribeToProject creates a webhook subscription for a project including objects owned by it
*/
func (a *Client) SubscribeToProject(params *SubscribeToProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscribeToProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscribeToProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "subscribeToProject",
		Method:             "PUT",
		PathPattern:        "/user/webhooks/projects/{owner}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscribeToProjectReader{formats: a.formats},
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
	success, ok := result.(*SubscribeToProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for subscribeToProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SubscribeToUser creates a webhook subscription for a user account including objects owned by it
*/
func (a *Client) SubscribeToUser(params *SubscribeToUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscribeToUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscribeToUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "subscribeToUser",
		Method:             "PUT",
		PathPattern:        "/user/webhooks/users/{account}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscribeToUserReader{formats: a.formats},
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
	success, ok := result.(*SubscribeToUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for subscribeToUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UnsubscribeFromDataset deletes a webhook subscription for a dataset
*/
func (a *Client) UnsubscribeFromDataset(params *UnsubscribeFromDatasetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnsubscribeFromDatasetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnsubscribeFromDatasetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "unsubscribeFromDataset",
		Method:             "DELETE",
		PathPattern:        "/user/webhooks/datasets/{owner}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnsubscribeFromDatasetReader{formats: a.formats},
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
	success, ok := result.(*UnsubscribeFromDatasetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unsubscribeFromDataset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UnsubscribeFromProject deletes a webhook subscription for a project
*/
func (a *Client) UnsubscribeFromProject(params *UnsubscribeFromProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnsubscribeFromProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnsubscribeFromProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "unsubscribeFromProject",
		Method:             "DELETE",
		PathPattern:        "/user/webhooks/projects/{owner}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnsubscribeFromProjectReader{formats: a.formats},
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
	success, ok := result.(*UnsubscribeFromProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unsubscribeFromProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UnsubscribeFromUser deletes a webhook subscription for a user account
*/
func (a *Client) UnsubscribeFromUser(params *UnsubscribeFromUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnsubscribeFromUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnsubscribeFromUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "unsubscribeFromUser",
		Method:             "DELETE",
		PathPattern:        "/user/webhooks/users/{account}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnsubscribeFromUserReader{formats: a.formats},
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
	success, ok := result.(*UnsubscribeFromUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unsubscribeFromUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
