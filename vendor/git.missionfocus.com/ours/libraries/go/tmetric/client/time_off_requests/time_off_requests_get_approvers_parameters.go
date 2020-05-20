// Code generated by go-swagger; DO NOT EDIT.

package time_off_requests

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

// NewTimeOffRequestsGetApproversParams creates a new TimeOffRequestsGetApproversParams object
// with the default values initialized.
func NewTimeOffRequestsGetApproversParams() *TimeOffRequestsGetApproversParams {
	var ()
	return &TimeOffRequestsGetApproversParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTimeOffRequestsGetApproversParamsWithTimeout creates a new TimeOffRequestsGetApproversParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimeOffRequestsGetApproversParamsWithTimeout(timeout time.Duration) *TimeOffRequestsGetApproversParams {
	var ()
	return &TimeOffRequestsGetApproversParams{

		timeout: timeout,
	}
}

// NewTimeOffRequestsGetApproversParamsWithContext creates a new TimeOffRequestsGetApproversParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimeOffRequestsGetApproversParamsWithContext(ctx context.Context) *TimeOffRequestsGetApproversParams {
	var ()
	return &TimeOffRequestsGetApproversParams{

		Context: ctx,
	}
}

// NewTimeOffRequestsGetApproversParamsWithHTTPClient creates a new TimeOffRequestsGetApproversParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimeOffRequestsGetApproversParamsWithHTTPClient(client *http.Client) *TimeOffRequestsGetApproversParams {
	var ()
	return &TimeOffRequestsGetApproversParams{
		HTTPClient: client,
	}
}

/*TimeOffRequestsGetApproversParams contains all the parameters to send to the API endpoint
for the time off requests get approvers operation typically these are written to a http.Request
*/
type TimeOffRequestsGetApproversParams struct {

	/*AccountID*/
	AccountID int32
	/*RequesterID*/
	RequesterID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the time off requests get approvers params
func (o *TimeOffRequestsGetApproversParams) WithTimeout(timeout time.Duration) *TimeOffRequestsGetApproversParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the time off requests get approvers params
func (o *TimeOffRequestsGetApproversParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the time off requests get approvers params
func (o *TimeOffRequestsGetApproversParams) WithContext(ctx context.Context) *TimeOffRequestsGetApproversParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the time off requests get approvers params
func (o *TimeOffRequestsGetApproversParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the time off requests get approvers params
func (o *TimeOffRequestsGetApproversParams) WithHTTPClient(client *http.Client) *TimeOffRequestsGetApproversParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the time off requests get approvers params
func (o *TimeOffRequestsGetApproversParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the time off requests get approvers params
func (o *TimeOffRequestsGetApproversParams) WithAccountID(accountID int32) *TimeOffRequestsGetApproversParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the time off requests get approvers params
func (o *TimeOffRequestsGetApproversParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithRequesterID adds the requesterID to the time off requests get approvers params
func (o *TimeOffRequestsGetApproversParams) WithRequesterID(requesterID int32) *TimeOffRequestsGetApproversParams {
	o.SetRequesterID(requesterID)
	return o
}

// SetRequesterID adds the requesterId to the time off requests get approvers params
func (o *TimeOffRequestsGetApproversParams) SetRequesterID(requesterID int32) {
	o.RequesterID = requesterID
}

// WriteToRequest writes these params to a swagger request
func (o *TimeOffRequestsGetApproversParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	// query param requesterId
	qrRequesterID := o.RequesterID
	qRequesterID := swag.FormatInt32(qrRequesterID)
	if qRequesterID != "" {
		if err := r.SetQueryParam("requesterId", qRequesterID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
