// Code generated by go-swagger; DO NOT EDIT.

package time_entries

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

	"git.missionfocus.com/ours/code/tools/tmetric-api/models"
)

// NewTimeEntriesPostExternalIssueTimeEntriesParams creates a new TimeEntriesPostExternalIssueTimeEntriesParams object
// with the default values initialized.
func NewTimeEntriesPostExternalIssueTimeEntriesParams() *TimeEntriesPostExternalIssueTimeEntriesParams {
	var ()
	return &TimeEntriesPostExternalIssueTimeEntriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTimeEntriesPostExternalIssueTimeEntriesParamsWithTimeout creates a new TimeEntriesPostExternalIssueTimeEntriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimeEntriesPostExternalIssueTimeEntriesParamsWithTimeout(timeout time.Duration) *TimeEntriesPostExternalIssueTimeEntriesParams {
	var ()
	return &TimeEntriesPostExternalIssueTimeEntriesParams{

		timeout: timeout,
	}
}

// NewTimeEntriesPostExternalIssueTimeEntriesParamsWithContext creates a new TimeEntriesPostExternalIssueTimeEntriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimeEntriesPostExternalIssueTimeEntriesParamsWithContext(ctx context.Context) *TimeEntriesPostExternalIssueTimeEntriesParams {
	var ()
	return &TimeEntriesPostExternalIssueTimeEntriesParams{

		Context: ctx,
	}
}

// NewTimeEntriesPostExternalIssueTimeEntriesParamsWithHTTPClient creates a new TimeEntriesPostExternalIssueTimeEntriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimeEntriesPostExternalIssueTimeEntriesParamsWithHTTPClient(client *http.Client) *TimeEntriesPostExternalIssueTimeEntriesParams {
	var ()
	return &TimeEntriesPostExternalIssueTimeEntriesParams{
		HTTPClient: client,
	}
}

/*TimeEntriesPostExternalIssueTimeEntriesParams contains all the parameters to send to the API endpoint
for the time entries post external issue time entries operation typically these are written to a http.Request
*/
type TimeEntriesPostExternalIssueTimeEntriesParams struct {

	/*AccountID*/
	AccountID int32
	/*Identifiers*/
	Identifiers []*models.WebToolIssueIdentifier

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the time entries post external issue time entries params
func (o *TimeEntriesPostExternalIssueTimeEntriesParams) WithTimeout(timeout time.Duration) *TimeEntriesPostExternalIssueTimeEntriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the time entries post external issue time entries params
func (o *TimeEntriesPostExternalIssueTimeEntriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the time entries post external issue time entries params
func (o *TimeEntriesPostExternalIssueTimeEntriesParams) WithContext(ctx context.Context) *TimeEntriesPostExternalIssueTimeEntriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the time entries post external issue time entries params
func (o *TimeEntriesPostExternalIssueTimeEntriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the time entries post external issue time entries params
func (o *TimeEntriesPostExternalIssueTimeEntriesParams) WithHTTPClient(client *http.Client) *TimeEntriesPostExternalIssueTimeEntriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the time entries post external issue time entries params
func (o *TimeEntriesPostExternalIssueTimeEntriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the time entries post external issue time entries params
func (o *TimeEntriesPostExternalIssueTimeEntriesParams) WithAccountID(accountID int32) *TimeEntriesPostExternalIssueTimeEntriesParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the time entries post external issue time entries params
func (o *TimeEntriesPostExternalIssueTimeEntriesParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithIdentifiers adds the identifiers to the time entries post external issue time entries params
func (o *TimeEntriesPostExternalIssueTimeEntriesParams) WithIdentifiers(identifiers []*models.WebToolIssueIdentifier) *TimeEntriesPostExternalIssueTimeEntriesParams {
	o.SetIdentifiers(identifiers)
	return o
}

// SetIdentifiers adds the identifiers to the time entries post external issue time entries params
func (o *TimeEntriesPostExternalIssueTimeEntriesParams) SetIdentifiers(identifiers []*models.WebToolIssueIdentifier) {
	o.Identifiers = identifiers
}

// WriteToRequest writes these params to a swagger request
func (o *TimeEntriesPostExternalIssueTimeEntriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	if o.Identifiers != nil {
		if err := r.SetBodyParam(o.Identifiers); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
