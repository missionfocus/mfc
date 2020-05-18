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

	"git.missionfocus.com/ours/code/tools/tmetric-api/models"
)

// NewTimeOffRequestsPostTimeOffRequestParams creates a new TimeOffRequestsPostTimeOffRequestParams object
// with the default values initialized.
func NewTimeOffRequestsPostTimeOffRequestParams() *TimeOffRequestsPostTimeOffRequestParams {
	var ()
	return &TimeOffRequestsPostTimeOffRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTimeOffRequestsPostTimeOffRequestParamsWithTimeout creates a new TimeOffRequestsPostTimeOffRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimeOffRequestsPostTimeOffRequestParamsWithTimeout(timeout time.Duration) *TimeOffRequestsPostTimeOffRequestParams {
	var ()
	return &TimeOffRequestsPostTimeOffRequestParams{

		timeout: timeout,
	}
}

// NewTimeOffRequestsPostTimeOffRequestParamsWithContext creates a new TimeOffRequestsPostTimeOffRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimeOffRequestsPostTimeOffRequestParamsWithContext(ctx context.Context) *TimeOffRequestsPostTimeOffRequestParams {
	var ()
	return &TimeOffRequestsPostTimeOffRequestParams{

		Context: ctx,
	}
}

// NewTimeOffRequestsPostTimeOffRequestParamsWithHTTPClient creates a new TimeOffRequestsPostTimeOffRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimeOffRequestsPostTimeOffRequestParamsWithHTTPClient(client *http.Client) *TimeOffRequestsPostTimeOffRequestParams {
	var ()
	return &TimeOffRequestsPostTimeOffRequestParams{
		HTTPClient: client,
	}
}

/*TimeOffRequestsPostTimeOffRequestParams contains all the parameters to send to the API endpoint
for the time off requests post time off request operation typically these are written to a http.Request
*/
type TimeOffRequestsPostTimeOffRequestParams struct {

	/*AccountID*/
	AccountID int32
	/*Request*/
	Request *models.TimeOffRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the time off requests post time off request params
func (o *TimeOffRequestsPostTimeOffRequestParams) WithTimeout(timeout time.Duration) *TimeOffRequestsPostTimeOffRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the time off requests post time off request params
func (o *TimeOffRequestsPostTimeOffRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the time off requests post time off request params
func (o *TimeOffRequestsPostTimeOffRequestParams) WithContext(ctx context.Context) *TimeOffRequestsPostTimeOffRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the time off requests post time off request params
func (o *TimeOffRequestsPostTimeOffRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the time off requests post time off request params
func (o *TimeOffRequestsPostTimeOffRequestParams) WithHTTPClient(client *http.Client) *TimeOffRequestsPostTimeOffRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the time off requests post time off request params
func (o *TimeOffRequestsPostTimeOffRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the time off requests post time off request params
func (o *TimeOffRequestsPostTimeOffRequestParams) WithAccountID(accountID int32) *TimeOffRequestsPostTimeOffRequestParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the time off requests post time off request params
func (o *TimeOffRequestsPostTimeOffRequestParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithRequest adds the request to the time off requests post time off request params
func (o *TimeOffRequestsPostTimeOffRequestParams) WithRequest(request *models.TimeOffRequest) *TimeOffRequestsPostTimeOffRequestParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the time off requests post time off request params
func (o *TimeOffRequestsPostTimeOffRequestParams) SetRequest(request *models.TimeOffRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *TimeOffRequestsPostTimeOffRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
