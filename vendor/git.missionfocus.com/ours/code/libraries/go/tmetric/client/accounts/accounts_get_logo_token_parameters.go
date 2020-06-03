// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewAccountsGetLogoTokenParams creates a new AccountsGetLogoTokenParams object
// with the default values initialized.
func NewAccountsGetLogoTokenParams() *AccountsGetLogoTokenParams {
	var ()
	return &AccountsGetLogoTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAccountsGetLogoTokenParamsWithTimeout creates a new AccountsGetLogoTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountsGetLogoTokenParamsWithTimeout(timeout time.Duration) *AccountsGetLogoTokenParams {
	var ()
	return &AccountsGetLogoTokenParams{

		timeout: timeout,
	}
}

// NewAccountsGetLogoTokenParamsWithContext creates a new AccountsGetLogoTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountsGetLogoTokenParamsWithContext(ctx context.Context) *AccountsGetLogoTokenParams {
	var ()
	return &AccountsGetLogoTokenParams{

		Context: ctx,
	}
}

// NewAccountsGetLogoTokenParamsWithHTTPClient creates a new AccountsGetLogoTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAccountsGetLogoTokenParamsWithHTTPClient(client *http.Client) *AccountsGetLogoTokenParams {
	var ()
	return &AccountsGetLogoTokenParams{
		HTTPClient: client,
	}
}

/*AccountsGetLogoTokenParams contains all the parameters to send to the API endpoint
for the accounts get logo token operation typically these are written to a http.Request
*/
type AccountsGetLogoTokenParams struct {

	/*AccountID*/
	AccountID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the accounts get logo token params
func (o *AccountsGetLogoTokenParams) WithTimeout(timeout time.Duration) *AccountsGetLogoTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accounts get logo token params
func (o *AccountsGetLogoTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accounts get logo token params
func (o *AccountsGetLogoTokenParams) WithContext(ctx context.Context) *AccountsGetLogoTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accounts get logo token params
func (o *AccountsGetLogoTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accounts get logo token params
func (o *AccountsGetLogoTokenParams) WithHTTPClient(client *http.Client) *AccountsGetLogoTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accounts get logo token params
func (o *AccountsGetLogoTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the accounts get logo token params
func (o *AccountsGetLogoTokenParams) WithAccountID(accountID int32) *AccountsGetLogoTokenParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the accounts get logo token params
func (o *AccountsGetLogoTokenParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *AccountsGetLogoTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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