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

	"git.missionfocus.com/ours/libraries/go/tmetric/models"
)

// NewProjectsPostProjectsParams creates a new ProjectsPostProjectsParams object
// with the default values initialized.
func NewProjectsPostProjectsParams() *ProjectsPostProjectsParams {
	var ()
	return &ProjectsPostProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsPostProjectsParamsWithTimeout creates a new ProjectsPostProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsPostProjectsParamsWithTimeout(timeout time.Duration) *ProjectsPostProjectsParams {
	var ()
	return &ProjectsPostProjectsParams{

		timeout: timeout,
	}
}

// NewProjectsPostProjectsParamsWithContext creates a new ProjectsPostProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsPostProjectsParamsWithContext(ctx context.Context) *ProjectsPostProjectsParams {
	var ()
	return &ProjectsPostProjectsParams{

		Context: ctx,
	}
}

// NewProjectsPostProjectsParamsWithHTTPClient creates a new ProjectsPostProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsPostProjectsParamsWithHTTPClient(client *http.Client) *ProjectsPostProjectsParams {
	var ()
	return &ProjectsPostProjectsParams{
		HTTPClient: client,
	}
}

/*ProjectsPostProjectsParams contains all the parameters to send to the API endpoint
for the projects post projects operation typically these are written to a http.Request
*/
type ProjectsPostProjectsParams struct {

	/*AccountID*/
	AccountID int32
	/*Update*/
	Update *models.ProjectsUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects post projects params
func (o *ProjectsPostProjectsParams) WithTimeout(timeout time.Duration) *ProjectsPostProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects post projects params
func (o *ProjectsPostProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects post projects params
func (o *ProjectsPostProjectsParams) WithContext(ctx context.Context) *ProjectsPostProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects post projects params
func (o *ProjectsPostProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects post projects params
func (o *ProjectsPostProjectsParams) WithHTTPClient(client *http.Client) *ProjectsPostProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects post projects params
func (o *ProjectsPostProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the projects post projects params
func (o *ProjectsPostProjectsParams) WithAccountID(accountID int32) *ProjectsPostProjectsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the projects post projects params
func (o *ProjectsPostProjectsParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithUpdate adds the update to the projects post projects params
func (o *ProjectsPostProjectsParams) WithUpdate(update *models.ProjectsUpdate) *ProjectsPostProjectsParams {
	o.SetUpdate(update)
	return o
}

// SetUpdate adds the update to the projects post projects params
func (o *ProjectsPostProjectsParams) SetUpdate(update *models.ProjectsUpdate) {
	o.Update = update
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsPostProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	if o.Update != nil {
		if err := r.SetBodyParam(o.Update); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
