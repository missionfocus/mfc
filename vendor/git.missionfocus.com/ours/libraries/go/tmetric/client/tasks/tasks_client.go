// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	TasksDeleteTask(params *TasksDeleteTaskParams, authInfo runtime.ClientAuthInfoWriter) (*TasksDeleteTaskNoContent, error)

	TasksGetTask(params *TasksGetTaskParams, authInfo runtime.ClientAuthInfoWriter) (*TasksGetTaskOK, error)

	TasksGetTasks(params *TasksGetTasksParams, authInfo runtime.ClientAuthInfoWriter) (*TasksGetTasksOK, error)

	TasksPostTask(params *TasksPostTaskParams, authInfo runtime.ClientAuthInfoWriter) (*TasksPostTaskOK, error)

	TasksPutTask(params *TasksPutTaskParams, authInfo runtime.ClientAuthInfoWriter) (*TasksPutTaskNoContent, error)

	TasksPutTasks(params *TasksPutTasksParams, authInfo runtime.ClientAuthInfoWriter) (*TasksPutTasksOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TasksDeleteTask deletes task
*/
func (a *Client) TasksDeleteTask(params *TasksDeleteTaskParams, authInfo runtime.ClientAuthInfoWriter) (*TasksDeleteTaskNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTasksDeleteTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Tasks_DeleteTask",
		Method:             "DELETE",
		PathPattern:        "/api/accounts/{accountId}/tasks/{taskId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TasksDeleteTaskReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TasksDeleteTaskNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Tasks_DeleteTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TasksGetTask gets task by task id
*/
func (a *Client) TasksGetTask(params *TasksGetTaskParams, authInfo runtime.ClientAuthInfoWriter) (*TasksGetTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTasksGetTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Tasks_GetTask",
		Method:             "GET",
		PathPattern:        "/api/accounts/{accountId}/tasks/{taskId}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TasksGetTaskReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TasksGetTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Tasks_GetTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TasksGetTasks gets tasks list
*/
func (a *Client) TasksGetTasks(params *TasksGetTasksParams, authInfo runtime.ClientAuthInfoWriter) (*TasksGetTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTasksGetTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Tasks_GetTasks",
		Method:             "GET",
		PathPattern:        "/api/accounts/{accountId}/tasks",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TasksGetTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TasksGetTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Tasks_GetTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TasksPostTask creates task
*/
func (a *Client) TasksPostTask(params *TasksPostTaskParams, authInfo runtime.ClientAuthInfoWriter) (*TasksPostTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTasksPostTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Tasks_PostTask",
		Method:             "POST",
		PathPattern:        "/api/accounts/{accountId}/tasks",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TasksPostTaskReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TasksPostTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Tasks_PostTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TasksPutTask edits task
*/
func (a *Client) TasksPutTask(params *TasksPutTaskParams, authInfo runtime.ClientAuthInfoWriter) (*TasksPutTaskNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTasksPutTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Tasks_PutTask",
		Method:             "PUT",
		PathPattern:        "/api/accounts/{accountId}/tasks/{taskId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TasksPutTaskReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TasksPutTaskNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Tasks_PutTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TasksPutTasks edits multiple tasks
*/
func (a *Client) TasksPutTasks(params *TasksPutTasksParams, authInfo runtime.ClientAuthInfoWriter) (*TasksPutTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTasksPutTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Tasks_PutTasks",
		Method:             "PUT",
		PathPattern:        "/api/accounts/{accountId}/tasks/bulk",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TasksPutTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TasksPutTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Tasks_PutTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
