// Code generated by go-swagger; DO NOT EDIT.

package demo_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new demo data API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for demo data API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DemoDataDeleteDemoData(params *DemoDataDeleteDemoDataParams, authInfo runtime.ClientAuthInfoWriter) (*DemoDataDeleteDemoDataNoContent, error)

	DemoDataPostDemoData(params *DemoDataPostDemoDataParams, authInfo runtime.ClientAuthInfoWriter) (*DemoDataPostDemoDataNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DemoDataDeleteDemoData deletes demo data
*/
func (a *Client) DemoDataDeleteDemoData(params *DemoDataDeleteDemoDataParams, authInfo runtime.ClientAuthInfoWriter) (*DemoDataDeleteDemoDataNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDemoDataDeleteDemoDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DemoData_DeleteDemoData",
		Method:             "DELETE",
		PathPattern:        "/api/demodata/{accountId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DemoDataDeleteDemoDataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DemoDataDeleteDemoDataNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DemoData_DeleteDemoData: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DemoDataPostDemoData generates demo data
*/
func (a *Client) DemoDataPostDemoData(params *DemoDataPostDemoDataParams, authInfo runtime.ClientAuthInfoWriter) (*DemoDataPostDemoDataNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDemoDataPostDemoDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DemoData_PostDemoData",
		Method:             "POST",
		PathPattern:        "/api/demodata/{accountId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DemoDataPostDemoDataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DemoDataPostDemoDataNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DemoData_PostDemoData: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
