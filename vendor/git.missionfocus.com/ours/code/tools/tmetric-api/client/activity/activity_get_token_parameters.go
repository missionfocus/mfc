// Code generated by go-swagger; DO NOT EDIT.

package activity

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

// NewActivityGetTokenParams creates a new ActivityGetTokenParams object
// with the default values initialized.
func NewActivityGetTokenParams() *ActivityGetTokenParams {
	var ()
	return &ActivityGetTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActivityGetTokenParamsWithTimeout creates a new ActivityGetTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActivityGetTokenParamsWithTimeout(timeout time.Duration) *ActivityGetTokenParams {
	var ()
	return &ActivityGetTokenParams{

		timeout: timeout,
	}
}

// NewActivityGetTokenParamsWithContext creates a new ActivityGetTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewActivityGetTokenParamsWithContext(ctx context.Context) *ActivityGetTokenParams {
	var ()
	return &ActivityGetTokenParams{

		Context: ctx,
	}
}

// NewActivityGetTokenParamsWithHTTPClient creates a new ActivityGetTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActivityGetTokenParamsWithHTTPClient(client *http.Client) *ActivityGetTokenParams {
	var ()
	return &ActivityGetTokenParams{
		HTTPClient: client,
	}
}

/*ActivityGetTokenParams contains all the parameters to send to the API endpoint
for the activity get token operation typically these are written to a http.Request
*/
type ActivityGetTokenParams struct {

	/*AccountID*/
	AccountID int32
	/*Timestamp*/
	Timestamp strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the activity get token params
func (o *ActivityGetTokenParams) WithTimeout(timeout time.Duration) *ActivityGetTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the activity get token params
func (o *ActivityGetTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the activity get token params
func (o *ActivityGetTokenParams) WithContext(ctx context.Context) *ActivityGetTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the activity get token params
func (o *ActivityGetTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the activity get token params
func (o *ActivityGetTokenParams) WithHTTPClient(client *http.Client) *ActivityGetTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the activity get token params
func (o *ActivityGetTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the activity get token params
func (o *ActivityGetTokenParams) WithAccountID(accountID int32) *ActivityGetTokenParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the activity get token params
func (o *ActivityGetTokenParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithTimestamp adds the timestamp to the activity get token params
func (o *ActivityGetTokenParams) WithTimestamp(timestamp strfmt.DateTime) *ActivityGetTokenParams {
	o.SetTimestamp(timestamp)
	return o
}

// SetTimestamp adds the timestamp to the activity get token params
func (o *ActivityGetTokenParams) SetTimestamp(timestamp strfmt.DateTime) {
	o.Timestamp = timestamp
}

// WriteToRequest writes these params to a swagger request
func (o *ActivityGetTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	// query param timestamp
	qrTimestamp := o.Timestamp
	qTimestamp := qrTimestamp.String()
	if qTimestamp != "" {
		if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
