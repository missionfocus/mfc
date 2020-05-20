// Code generated by go-swagger; DO NOT EDIT.

package tasks

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
)

// NewTasksGetTasksParams creates a new TasksGetTasksParams object
// with the default values initialized.
func NewTasksGetTasksParams() *TasksGetTasksParams {
	var ()
	return &TasksGetTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTasksGetTasksParamsWithTimeout creates a new TasksGetTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTasksGetTasksParamsWithTimeout(timeout time.Duration) *TasksGetTasksParams {
	var ()
	return &TasksGetTasksParams{

		timeout: timeout,
	}
}

// NewTasksGetTasksParamsWithContext creates a new TasksGetTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewTasksGetTasksParamsWithContext(ctx context.Context) *TasksGetTasksParams {
	var ()
	return &TasksGetTasksParams{

		Context: ctx,
	}
}

// NewTasksGetTasksParamsWithHTTPClient creates a new TasksGetTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTasksGetTasksParamsWithHTTPClient(client *http.Client) *TasksGetTasksParams {
	var ()
	return &TasksGetTasksParams{
		HTTPClient: client,
	}
}

/*TasksGetTasksParams contains all the parameters to send to the API endpoint
for the tasks get tasks operation typically these are written to a http.Request
*/
type TasksGetTasksParams struct {

	/*AccountID*/
	AccountID int32
	/*FilterAssigneeList
	  Gets or sets profile list for report filter.

	*/
	FilterAssigneeList []int32
	/*FilterCompleted
	  Gets or sets status for task filter.

	*/
	FilterCompleted *bool
	/*FilterGroupList
	  Gets or sets group list for report filter.

	*/
	FilterGroupList []int32
	/*FilterProjectList
	  Gets or sets project list for report filter.

	*/
	FilterProjectList []int32
	/*FilterTagList
	  Gets or sets tag list for report filter.

	*/
	FilterTagList []int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tasks get tasks params
func (o *TasksGetTasksParams) WithTimeout(timeout time.Duration) *TasksGetTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tasks get tasks params
func (o *TasksGetTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tasks get tasks params
func (o *TasksGetTasksParams) WithContext(ctx context.Context) *TasksGetTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tasks get tasks params
func (o *TasksGetTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tasks get tasks params
func (o *TasksGetTasksParams) WithHTTPClient(client *http.Client) *TasksGetTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tasks get tasks params
func (o *TasksGetTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the tasks get tasks params
func (o *TasksGetTasksParams) WithAccountID(accountID int32) *TasksGetTasksParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the tasks get tasks params
func (o *TasksGetTasksParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithFilterAssigneeList adds the filterAssigneeList to the tasks get tasks params
func (o *TasksGetTasksParams) WithFilterAssigneeList(filterAssigneeList []int32) *TasksGetTasksParams {
	o.SetFilterAssigneeList(filterAssigneeList)
	return o
}

// SetFilterAssigneeList adds the filterAssigneeList to the tasks get tasks params
func (o *TasksGetTasksParams) SetFilterAssigneeList(filterAssigneeList []int32) {
	o.FilterAssigneeList = filterAssigneeList
}

// WithFilterCompleted adds the filterCompleted to the tasks get tasks params
func (o *TasksGetTasksParams) WithFilterCompleted(filterCompleted *bool) *TasksGetTasksParams {
	o.SetFilterCompleted(filterCompleted)
	return o
}

// SetFilterCompleted adds the filterCompleted to the tasks get tasks params
func (o *TasksGetTasksParams) SetFilterCompleted(filterCompleted *bool) {
	o.FilterCompleted = filterCompleted
}

// WithFilterGroupList adds the filterGroupList to the tasks get tasks params
func (o *TasksGetTasksParams) WithFilterGroupList(filterGroupList []int32) *TasksGetTasksParams {
	o.SetFilterGroupList(filterGroupList)
	return o
}

// SetFilterGroupList adds the filterGroupList to the tasks get tasks params
func (o *TasksGetTasksParams) SetFilterGroupList(filterGroupList []int32) {
	o.FilterGroupList = filterGroupList
}

// WithFilterProjectList adds the filterProjectList to the tasks get tasks params
func (o *TasksGetTasksParams) WithFilterProjectList(filterProjectList []int32) *TasksGetTasksParams {
	o.SetFilterProjectList(filterProjectList)
	return o
}

// SetFilterProjectList adds the filterProjectList to the tasks get tasks params
func (o *TasksGetTasksParams) SetFilterProjectList(filterProjectList []int32) {
	o.FilterProjectList = filterProjectList
}

// WithFilterTagList adds the filterTagList to the tasks get tasks params
func (o *TasksGetTasksParams) WithFilterTagList(filterTagList []int32) *TasksGetTasksParams {
	o.SetFilterTagList(filterTagList)
	return o
}

// SetFilterTagList adds the filterTagList to the tasks get tasks params
func (o *TasksGetTasksParams) SetFilterTagList(filterTagList []int32) {
	o.FilterTagList = filterTagList
}

// WriteToRequest writes these params to a swagger request
func (o *TasksGetTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	var valuesFilterAssigneeList []string
	for _, v := range o.FilterAssigneeList {
		valuesFilterAssigneeList = append(valuesFilterAssigneeList, swag.FormatInt32(v))
	}

	joinedFilterAssigneeList := swag.JoinByFormat(valuesFilterAssigneeList, "multi")
	// query array param filter.assigneeList
	if err := r.SetQueryParam("filter.assigneeList", joinedFilterAssigneeList...); err != nil {
		return err
	}

	if o.FilterCompleted != nil {

		// query param filter.completed
		var qrFilterCompleted bool
		if o.FilterCompleted != nil {
			qrFilterCompleted = *o.FilterCompleted
		}
		qFilterCompleted := swag.FormatBool(qrFilterCompleted)
		if qFilterCompleted != "" {
			if err := r.SetQueryParam("filter.completed", qFilterCompleted); err != nil {
				return err
			}
		}

	}

	var valuesFilterGroupList []string
	for _, v := range o.FilterGroupList {
		valuesFilterGroupList = append(valuesFilterGroupList, swag.FormatInt32(v))
	}

	joinedFilterGroupList := swag.JoinByFormat(valuesFilterGroupList, "multi")
	// query array param filter.groupList
	if err := r.SetQueryParam("filter.groupList", joinedFilterGroupList...); err != nil {
		return err
	}

	var valuesFilterProjectList []string
	for _, v := range o.FilterProjectList {
		valuesFilterProjectList = append(valuesFilterProjectList, swag.FormatInt32(v))
	}

	joinedFilterProjectList := swag.JoinByFormat(valuesFilterProjectList, "multi")
	// query array param filter.projectList
	if err := r.SetQueryParam("filter.projectList", joinedFilterProjectList...); err != nil {
		return err
	}

	var valuesFilterTagList []string
	for _, v := range o.FilterTagList {
		valuesFilterTagList = append(valuesFilterTagList, swag.FormatInt32(v))
	}

	joinedFilterTagList := swag.JoinByFormat(valuesFilterTagList, "multi")
	// query array param filter.tagList
	if err := r.SetQueryParam("filter.tagList", joinedFilterTagList...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
