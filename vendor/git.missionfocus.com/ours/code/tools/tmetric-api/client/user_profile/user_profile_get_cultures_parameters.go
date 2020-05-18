// Code generated by go-swagger; DO NOT EDIT.

package user_profile

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
)

// NewUserProfileGetCulturesParams creates a new UserProfileGetCulturesParams object
// with the default values initialized.
func NewUserProfileGetCulturesParams() *UserProfileGetCulturesParams {

	return &UserProfileGetCulturesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserProfileGetCulturesParamsWithTimeout creates a new UserProfileGetCulturesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserProfileGetCulturesParamsWithTimeout(timeout time.Duration) *UserProfileGetCulturesParams {

	return &UserProfileGetCulturesParams{

		timeout: timeout,
	}
}

// NewUserProfileGetCulturesParamsWithContext creates a new UserProfileGetCulturesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserProfileGetCulturesParamsWithContext(ctx context.Context) *UserProfileGetCulturesParams {

	return &UserProfileGetCulturesParams{

		Context: ctx,
	}
}

// NewUserProfileGetCulturesParamsWithHTTPClient creates a new UserProfileGetCulturesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserProfileGetCulturesParamsWithHTTPClient(client *http.Client) *UserProfileGetCulturesParams {

	return &UserProfileGetCulturesParams{
		HTTPClient: client,
	}
}

/*UserProfileGetCulturesParams contains all the parameters to send to the API endpoint
for the user profile get cultures operation typically these are written to a http.Request
*/
type UserProfileGetCulturesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user profile get cultures params
func (o *UserProfileGetCulturesParams) WithTimeout(timeout time.Duration) *UserProfileGetCulturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user profile get cultures params
func (o *UserProfileGetCulturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user profile get cultures params
func (o *UserProfileGetCulturesParams) WithContext(ctx context.Context) *UserProfileGetCulturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user profile get cultures params
func (o *UserProfileGetCulturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user profile get cultures params
func (o *UserProfileGetCulturesParams) WithHTTPClient(client *http.Client) *UserProfileGetCulturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user profile get cultures params
func (o *UserProfileGetCulturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UserProfileGetCulturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
