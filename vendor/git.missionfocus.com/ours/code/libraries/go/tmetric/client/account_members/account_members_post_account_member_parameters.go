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

	"git.missionfocus.com/ours/code/libraries/go/tmetric/models"
)

// NewAccountMembersPostAccountMemberParams creates a new AccountMembersPostAccountMemberParams object
// with the default values initialized.
func NewAccountMembersPostAccountMemberParams() *AccountMembersPostAccountMemberParams {
	var ()
	return &AccountMembersPostAccountMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAccountMembersPostAccountMemberParamsWithTimeout creates a new AccountMembersPostAccountMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountMembersPostAccountMemberParamsWithTimeout(timeout time.Duration) *AccountMembersPostAccountMemberParams {
	var ()
	return &AccountMembersPostAccountMemberParams{

		timeout: timeout,
	}
}

// NewAccountMembersPostAccountMemberParamsWithContext creates a new AccountMembersPostAccountMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountMembersPostAccountMemberParamsWithContext(ctx context.Context) *AccountMembersPostAccountMemberParams {
	var ()
	return &AccountMembersPostAccountMemberParams{

		Context: ctx,
	}
}

// NewAccountMembersPostAccountMemberParamsWithHTTPClient creates a new AccountMembersPostAccountMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAccountMembersPostAccountMemberParamsWithHTTPClient(client *http.Client) *AccountMembersPostAccountMemberParams {
	var ()
	return &AccountMembersPostAccountMemberParams{
		HTTPClient: client,
	}
}

/*AccountMembersPostAccountMemberParams contains all the parameters to send to the API endpoint
for the account members post account member operation typically these are written to a http.Request
*/
type AccountMembersPostAccountMemberParams struct {

	/*AccountID*/
	AccountID string
	/*InvitedMember*/
	InvitedMember *models.AccountMember
	/*Recaptcha*/
	Recaptcha *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) WithTimeout(timeout time.Duration) *AccountMembersPostAccountMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) WithContext(ctx context.Context) *AccountMembersPostAccountMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) WithHTTPClient(client *http.Client) *AccountMembersPostAccountMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) WithAccountID(accountID string) *AccountMembersPostAccountMemberParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithInvitedMember adds the invitedMember to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) WithInvitedMember(invitedMember *models.AccountMember) *AccountMembersPostAccountMemberParams {
	o.SetInvitedMember(invitedMember)
	return o
}

// SetInvitedMember adds the invitedMember to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) SetInvitedMember(invitedMember *models.AccountMember) {
	o.InvitedMember = invitedMember
}

// WithRecaptcha adds the recaptcha to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) WithRecaptcha(recaptcha *string) *AccountMembersPostAccountMemberParams {
	o.SetRecaptcha(recaptcha)
	return o
}

// SetRecaptcha adds the recaptcha to the account members post account member params
func (o *AccountMembersPostAccountMemberParams) SetRecaptcha(recaptcha *string) {
	o.Recaptcha = recaptcha
}

// WriteToRequest writes these params to a swagger request
func (o *AccountMembersPostAccountMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	if o.InvitedMember != nil {
		if err := r.SetBodyParam(o.InvitedMember); err != nil {
			return err
		}
	}

	if o.Recaptcha != nil {

		// query param recaptcha
		var qrRecaptcha string
		if o.Recaptcha != nil {
			qrRecaptcha = *o.Recaptcha
		}
		qRecaptcha := qrRecaptcha
		if qRecaptcha != "" {
			if err := r.SetQueryParam("recaptcha", qRecaptcha); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}