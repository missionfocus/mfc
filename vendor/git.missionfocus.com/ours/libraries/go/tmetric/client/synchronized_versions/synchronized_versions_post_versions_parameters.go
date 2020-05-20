// Code generated by go-swagger; DO NOT EDIT.

package synchronized_versions

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

// NewSynchronizedVersionsPostVersionsParams creates a new SynchronizedVersionsPostVersionsParams object
// with the default values initialized.
func NewSynchronizedVersionsPostVersionsParams() *SynchronizedVersionsPostVersionsParams {
	var ()
	return &SynchronizedVersionsPostVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSynchronizedVersionsPostVersionsParamsWithTimeout creates a new SynchronizedVersionsPostVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSynchronizedVersionsPostVersionsParamsWithTimeout(timeout time.Duration) *SynchronizedVersionsPostVersionsParams {
	var ()
	return &SynchronizedVersionsPostVersionsParams{

		timeout: timeout,
	}
}

// NewSynchronizedVersionsPostVersionsParamsWithContext creates a new SynchronizedVersionsPostVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSynchronizedVersionsPostVersionsParamsWithContext(ctx context.Context) *SynchronizedVersionsPostVersionsParams {
	var ()
	return &SynchronizedVersionsPostVersionsParams{

		Context: ctx,
	}
}

// NewSynchronizedVersionsPostVersionsParamsWithHTTPClient creates a new SynchronizedVersionsPostVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSynchronizedVersionsPostVersionsParamsWithHTTPClient(client *http.Client) *SynchronizedVersionsPostVersionsParams {
	var ()
	return &SynchronizedVersionsPostVersionsParams{
		HTTPClient: client,
	}
}

/*SynchronizedVersionsPostVersionsParams contains all the parameters to send to the API endpoint
for the synchronized versions post versions operation typically these are written to a http.Request
*/
type SynchronizedVersionsPostVersionsParams struct {

	/*AccountID*/
	AccountID int32
	/*IntegrationID*/
	IntegrationID int32
	/*Versions*/
	Versions []*models.SynchronizedVersion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) WithTimeout(timeout time.Duration) *SynchronizedVersionsPostVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) WithContext(ctx context.Context) *SynchronizedVersionsPostVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) WithHTTPClient(client *http.Client) *SynchronizedVersionsPostVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) WithAccountID(accountID int32) *SynchronizedVersionsPostVersionsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithIntegrationID adds the integrationID to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) WithIntegrationID(integrationID int32) *SynchronizedVersionsPostVersionsParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) SetIntegrationID(integrationID int32) {
	o.IntegrationID = integrationID
}

// WithVersions adds the versions to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) WithVersions(versions []*models.SynchronizedVersion) *SynchronizedVersionsPostVersionsParams {
	o.SetVersions(versions)
	return o
}

// SetVersions adds the versions to the synchronized versions post versions params
func (o *SynchronizedVersionsPostVersionsParams) SetVersions(versions []*models.SynchronizedVersion) {
	o.Versions = versions
}

// WriteToRequest writes these params to a swagger request
func (o *SynchronizedVersionsPostVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	// path param integrationId
	if err := r.SetPathParam("integrationId", swag.FormatInt32(o.IntegrationID)); err != nil {
		return err
	}

	if o.Versions != nil {
		if err := r.SetBodyParam(o.Versions); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
