// Code generated by go-swagger; DO NOT EDIT.

package item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new item API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for item API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddItem(params *AddItemParams, opts ...ClientOption) (*AddItemOK, error)

	RemoveItem(params *RemoveItemParams, opts ...ClientOption) (*RemoveItemOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddItem adds or add quantity to an item in the shopping basket
*/
func (a *Client) AddItem(params *AddItemParams, opts ...ClientOption) (*AddItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddItemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addItem",
		Method:             "PUT",
		PathPattern:        "/api/baskets/{id}/addItem",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddItemReader{formats: a.formats},
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
	success, ok := result.(*AddItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddItemDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RemoveItem removes or remove quantity to an item in the shopping basket
*/
func (a *Client) RemoveItem(params *RemoveItemParams, opts ...ClientOption) (*RemoveItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveItemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeItem",
		Method:             "PUT",
		PathPattern:        "/api/baskets/{id}/removeItem",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveItemReader{formats: a.formats},
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
	success, ok := result.(*RemoveItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RemoveItemDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}