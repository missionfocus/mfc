// Code generated by go-swagger; DO NOT EDIT.

package user_groups

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

// NewUserGroupsGetGroupParams creates a new UserGroupsGetGroupParams object
// with the default values initialized.
func NewUserGroupsGetGroupParams() *UserGroupsGetGroupParams {
	var ()
	return &UserGroupsGetGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserGroupsGetGroupParamsWithTimeout creates a new UserGroupsGetGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserGroupsGetGroupParamsWithTimeout(timeout time.Duration) *UserGroupsGetGroupParams {
	var ()
	return &UserGroupsGetGroupParams{

		timeout: timeout,
	}
}

// NewUserGroupsGetGroupParamsWithContext creates a new UserGroupsGetGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserGroupsGetGroupParamsWithContext(ctx context.Context) *UserGroupsGetGroupParams {
	var ()
	return &UserGroupsGetGroupParams{

		Context: ctx,
	}
}

// NewUserGroupsGetGroupParamsWithHTTPClient creates a new UserGroupsGetGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserGroupsGetGroupParamsWithHTTPClient(client *http.Client) *UserGroupsGetGroupParams {
	var ()
	return &UserGroupsGetGroupParams{
		HTTPClient: client,
	}
}

/*UserGroupsGetGroupParams contains all the parameters to send to the API endpoint
for the user groups get group operation typically these are written to a http.Request
*/
type UserGroupsGetGroupParams struct {

	/*AccountID*/
	AccountID int32
	/*UserGroupID*/
	UserGroupID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user groups get group params
func (o *UserGroupsGetGroupParams) WithTimeout(timeout time.Duration) *UserGroupsGetGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user groups get group params
func (o *UserGroupsGetGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user groups get group params
func (o *UserGroupsGetGroupParams) WithContext(ctx context.Context) *UserGroupsGetGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user groups get group params
func (o *UserGroupsGetGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user groups get group params
func (o *UserGroupsGetGroupParams) WithHTTPClient(client *http.Client) *UserGroupsGetGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user groups get group params
func (o *UserGroupsGetGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the user groups get group params
func (o *UserGroupsGetGroupParams) WithAccountID(accountID int32) *UserGroupsGetGroupParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the user groups get group params
func (o *UserGroupsGetGroupParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithUserGroupID adds the userGroupID to the user groups get group params
func (o *UserGroupsGetGroupParams) WithUserGroupID(userGroupID int32) *UserGroupsGetGroupParams {
	o.SetUserGroupID(userGroupID)
	return o
}

// SetUserGroupID adds the userGroupId to the user groups get group params
func (o *UserGroupsGetGroupParams) SetUserGroupID(userGroupID int32) {
	o.UserGroupID = userGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *UserGroupsGetGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	// path param userGroupId
	if err := r.SetPathParam("userGroupId", swag.FormatInt32(o.UserGroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
