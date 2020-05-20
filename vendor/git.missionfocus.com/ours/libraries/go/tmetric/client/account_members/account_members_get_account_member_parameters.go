// Code generated by go-swagger; DO NOT EDIT.

package account_members

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

// NewAccountMembersGetAccountMemberParams creates a new AccountMembersGetAccountMemberParams object
// with the default values initialized.
func NewAccountMembersGetAccountMemberParams() *AccountMembersGetAccountMemberParams {
	var ()
	return &AccountMembersGetAccountMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAccountMembersGetAccountMemberParamsWithTimeout creates a new AccountMembersGetAccountMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountMembersGetAccountMemberParamsWithTimeout(timeout time.Duration) *AccountMembersGetAccountMemberParams {
	var ()
	return &AccountMembersGetAccountMemberParams{

		timeout: timeout,
	}
}

// NewAccountMembersGetAccountMemberParamsWithContext creates a new AccountMembersGetAccountMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountMembersGetAccountMemberParamsWithContext(ctx context.Context) *AccountMembersGetAccountMemberParams {
	var ()
	return &AccountMembersGetAccountMemberParams{

		Context: ctx,
	}
}

// NewAccountMembersGetAccountMemberParamsWithHTTPClient creates a new AccountMembersGetAccountMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAccountMembersGetAccountMemberParamsWithHTTPClient(client *http.Client) *AccountMembersGetAccountMemberParams {
	var ()
	return &AccountMembersGetAccountMemberParams{
		HTTPClient: client,
	}
}

/*AccountMembersGetAccountMemberParams contains all the parameters to send to the API endpoint
for the account members get account member operation typically these are written to a http.Request
*/
type AccountMembersGetAccountMemberParams struct {

	/*AccountID*/
	AccountID int32
	/*AccountMemberID*/
	AccountMemberID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the account members get account member params
func (o *AccountMembersGetAccountMemberParams) WithTimeout(timeout time.Duration) *AccountMembersGetAccountMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account members get account member params
func (o *AccountMembersGetAccountMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account members get account member params
func (o *AccountMembersGetAccountMemberParams) WithContext(ctx context.Context) *AccountMembersGetAccountMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account members get account member params
func (o *AccountMembersGetAccountMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account members get account member params
func (o *AccountMembersGetAccountMemberParams) WithHTTPClient(client *http.Client) *AccountMembersGetAccountMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account members get account member params
func (o *AccountMembersGetAccountMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the account members get account member params
func (o *AccountMembersGetAccountMemberParams) WithAccountID(accountID int32) *AccountMembersGetAccountMemberParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the account members get account member params
func (o *AccountMembersGetAccountMemberParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithAccountMemberID adds the accountMemberID to the account members get account member params
func (o *AccountMembersGetAccountMemberParams) WithAccountMemberID(accountMemberID int32) *AccountMembersGetAccountMemberParams {
	o.SetAccountMemberID(accountMemberID)
	return o
}

// SetAccountMemberID adds the accountMemberId to the account members get account member params
func (o *AccountMembersGetAccountMemberParams) SetAccountMemberID(accountMemberID int32) {
	o.AccountMemberID = accountMemberID
}

// WriteToRequest writes these params to a swagger request
func (o *AccountMembersGetAccountMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	// path param accountMemberId
	if err := r.SetPathParam("accountMemberId", swag.FormatInt32(o.AccountMemberID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
