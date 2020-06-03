// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewProjectsGetProjectAvatarTokenParams creates a new ProjectsGetProjectAvatarTokenParams object
// with the default values initialized.
func NewProjectsGetProjectAvatarTokenParams() *ProjectsGetProjectAvatarTokenParams {
	var ()
	return &ProjectsGetProjectAvatarTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsGetProjectAvatarTokenParamsWithTimeout creates a new ProjectsGetProjectAvatarTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsGetProjectAvatarTokenParamsWithTimeout(timeout time.Duration) *ProjectsGetProjectAvatarTokenParams {
	var ()
	return &ProjectsGetProjectAvatarTokenParams{

		timeout: timeout,
	}
}

// NewProjectsGetProjectAvatarTokenParamsWithContext creates a new ProjectsGetProjectAvatarTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsGetProjectAvatarTokenParamsWithContext(ctx context.Context) *ProjectsGetProjectAvatarTokenParams {
	var ()
	return &ProjectsGetProjectAvatarTokenParams{

		Context: ctx,
	}
}

// NewProjectsGetProjectAvatarTokenParamsWithHTTPClient creates a new ProjectsGetProjectAvatarTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsGetProjectAvatarTokenParamsWithHTTPClient(client *http.Client) *ProjectsGetProjectAvatarTokenParams {
	var ()
	return &ProjectsGetProjectAvatarTokenParams{
		HTTPClient: client,
	}
}

/*ProjectsGetProjectAvatarTokenParams contains all the parameters to send to the API endpoint
for the projects get project avatar token operation typically these are written to a http.Request
*/
type ProjectsGetProjectAvatarTokenParams struct {

	/*AccountID*/
	AccountID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects get project avatar token params
func (o *ProjectsGetProjectAvatarTokenParams) WithTimeout(timeout time.Duration) *ProjectsGetProjectAvatarTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects get project avatar token params
func (o *ProjectsGetProjectAvatarTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects get project avatar token params
func (o *ProjectsGetProjectAvatarTokenParams) WithContext(ctx context.Context) *ProjectsGetProjectAvatarTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects get project avatar token params
func (o *ProjectsGetProjectAvatarTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects get project avatar token params
func (o *ProjectsGetProjectAvatarTokenParams) WithHTTPClient(client *http.Client) *ProjectsGetProjectAvatarTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects get project avatar token params
func (o *ProjectsGetProjectAvatarTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the projects get project avatar token params
func (o *ProjectsGetProjectAvatarTokenParams) WithAccountID(accountID int32) *ProjectsGetProjectAvatarTokenParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the projects get project avatar token params
func (o *ProjectsGetProjectAvatarTokenParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsGetProjectAvatarTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}