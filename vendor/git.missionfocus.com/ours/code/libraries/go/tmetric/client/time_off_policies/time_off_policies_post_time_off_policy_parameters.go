// Code generated by go-swagger; DO NOT EDIT.

package time_off_policies

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

	"git.missionfocus.com/ours/code/libraries/go/tmetric/models"
)

// NewTimeOffPoliciesPostTimeOffPolicyParams creates a new TimeOffPoliciesPostTimeOffPolicyParams object
// with the default values initialized.
func NewTimeOffPoliciesPostTimeOffPolicyParams() *TimeOffPoliciesPostTimeOffPolicyParams {
	var ()
	return &TimeOffPoliciesPostTimeOffPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTimeOffPoliciesPostTimeOffPolicyParamsWithTimeout creates a new TimeOffPoliciesPostTimeOffPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimeOffPoliciesPostTimeOffPolicyParamsWithTimeout(timeout time.Duration) *TimeOffPoliciesPostTimeOffPolicyParams {
	var ()
	return &TimeOffPoliciesPostTimeOffPolicyParams{

		timeout: timeout,
	}
}

// NewTimeOffPoliciesPostTimeOffPolicyParamsWithContext creates a new TimeOffPoliciesPostTimeOffPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimeOffPoliciesPostTimeOffPolicyParamsWithContext(ctx context.Context) *TimeOffPoliciesPostTimeOffPolicyParams {
	var ()
	return &TimeOffPoliciesPostTimeOffPolicyParams{

		Context: ctx,
	}
}

// NewTimeOffPoliciesPostTimeOffPolicyParamsWithHTTPClient creates a new TimeOffPoliciesPostTimeOffPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimeOffPoliciesPostTimeOffPolicyParamsWithHTTPClient(client *http.Client) *TimeOffPoliciesPostTimeOffPolicyParams {
	var ()
	return &TimeOffPoliciesPostTimeOffPolicyParams{
		HTTPClient: client,
	}
}

/*TimeOffPoliciesPostTimeOffPolicyParams contains all the parameters to send to the API endpoint
for the time off policies post time off policy operation typically these are written to a http.Request
*/
type TimeOffPoliciesPostTimeOffPolicyParams struct {

	/*AccountID*/
	AccountID int32
	/*TimeOffPolicy*/
	TimeOffPolicy *models.TimeOffPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the time off policies post time off policy params
func (o *TimeOffPoliciesPostTimeOffPolicyParams) WithTimeout(timeout time.Duration) *TimeOffPoliciesPostTimeOffPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the time off policies post time off policy params
func (o *TimeOffPoliciesPostTimeOffPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the time off policies post time off policy params
func (o *TimeOffPoliciesPostTimeOffPolicyParams) WithContext(ctx context.Context) *TimeOffPoliciesPostTimeOffPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the time off policies post time off policy params
func (o *TimeOffPoliciesPostTimeOffPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the time off policies post time off policy params
func (o *TimeOffPoliciesPostTimeOffPolicyParams) WithHTTPClient(client *http.Client) *TimeOffPoliciesPostTimeOffPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the time off policies post time off policy params
func (o *TimeOffPoliciesPostTimeOffPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the time off policies post time off policy params
func (o *TimeOffPoliciesPostTimeOffPolicyParams) WithAccountID(accountID int32) *TimeOffPoliciesPostTimeOffPolicyParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the time off policies post time off policy params
func (o *TimeOffPoliciesPostTimeOffPolicyParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithTimeOffPolicy adds the timeOffPolicy to the time off policies post time off policy params
func (o *TimeOffPoliciesPostTimeOffPolicyParams) WithTimeOffPolicy(timeOffPolicy *models.TimeOffPolicy) *TimeOffPoliciesPostTimeOffPolicyParams {
	o.SetTimeOffPolicy(timeOffPolicy)
	return o
}

// SetTimeOffPolicy adds the timeOffPolicy to the time off policies post time off policy params
func (o *TimeOffPoliciesPostTimeOffPolicyParams) SetTimeOffPolicy(timeOffPolicy *models.TimeOffPolicy) {
	o.TimeOffPolicy = timeOffPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *TimeOffPoliciesPostTimeOffPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	if o.TimeOffPolicy != nil {
		if err := r.SetBodyParam(o.TimeOffPolicy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}