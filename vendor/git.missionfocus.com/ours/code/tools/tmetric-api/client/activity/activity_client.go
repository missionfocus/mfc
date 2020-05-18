// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new activity API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for activity API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ActivityDeleteActivity(params *ActivityDeleteActivityParams, authInfo runtime.ClientAuthInfoWriter) (*ActivityDeleteActivityNoContent, error)

	ActivityDeleteAllActivity(params *ActivityDeleteAllActivityParams, authInfo runtime.ClientAuthInfoWriter) (*ActivityDeleteAllActivityNoContent, error)

	ActivityGetDates(params *ActivityGetDatesParams, authInfo runtime.ClientAuthInfoWriter) (*ActivityGetDatesOK, error)

	ActivityGetScreenshots(params *ActivityGetScreenshotsParams, authInfo runtime.ClientAuthInfoWriter) (*ActivityGetScreenshotsOK, error)

	ActivityGetToken(params *ActivityGetTokenParams, authInfo runtime.ClientAuthInfoWriter) (*ActivityGetTokenOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ActivityDeleteActivity activity delete activity API
*/
func (a *Client) ActivityDeleteActivity(params *ActivityDeleteActivityParams, authInfo runtime.ClientAuthInfoWriter) (*ActivityDeleteActivityNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActivityDeleteActivityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Activity_DeleteActivity",
		Method:             "DELETE",
		PathPattern:        "/api/accounts/{accountId}/activity/{userProfileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActivityDeleteActivityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ActivityDeleteActivityNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Activity_DeleteActivity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ActivityDeleteAllActivity activity delete all activity API
*/
func (a *Client) ActivityDeleteAllActivity(params *ActivityDeleteAllActivityParams, authInfo runtime.ClientAuthInfoWriter) (*ActivityDeleteAllActivityNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActivityDeleteAllActivityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Activity_DeleteAllActivity",
		Method:             "DELETE",
		PathPattern:        "/api/accounts/{accountId}/activity/{userProfileId}/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActivityDeleteAllActivityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ActivityDeleteAllActivityNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Activity_DeleteAllActivity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ActivityGetDates activity get dates API
*/
func (a *Client) ActivityGetDates(params *ActivityGetDatesParams, authInfo runtime.ClientAuthInfoWriter) (*ActivityGetDatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActivityGetDatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Activity_GetDates",
		Method:             "GET",
		PathPattern:        "/api/accounts/{accountId}/activity/screenshots/dates",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActivityGetDatesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ActivityGetDatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Activity_GetDates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ActivityGetScreenshots activity get screenshots API
*/
func (a *Client) ActivityGetScreenshots(params *ActivityGetScreenshotsParams, authInfo runtime.ClientAuthInfoWriter) (*ActivityGetScreenshotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActivityGetScreenshotsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Activity_GetScreenshots",
		Method:             "GET",
		PathPattern:        "/api/accounts/{accountId}/activity/screenshots",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActivityGetScreenshotsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ActivityGetScreenshotsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Activity_GetScreenshots: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ActivityGetToken activity get token API
*/
func (a *Client) ActivityGetToken(params *ActivityGetTokenParams, authInfo runtime.ClientAuthInfoWriter) (*ActivityGetTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActivityGetTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Activity_GetToken",
		Method:             "GET",
		PathPattern:        "/api/accounts/{accountId}/activity/screenshots/token",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActivityGetTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ActivityGetTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Activity_GetToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
