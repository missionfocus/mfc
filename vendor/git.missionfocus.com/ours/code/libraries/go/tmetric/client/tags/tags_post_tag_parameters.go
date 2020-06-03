// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// NewTagsPostTagParams creates a new TagsPostTagParams object
// with the default values initialized.
func NewTagsPostTagParams() *TagsPostTagParams {
	var ()
	return &TagsPostTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTagsPostTagParamsWithTimeout creates a new TagsPostTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTagsPostTagParamsWithTimeout(timeout time.Duration) *TagsPostTagParams {
	var ()
	return &TagsPostTagParams{

		timeout: timeout,
	}
}

// NewTagsPostTagParamsWithContext creates a new TagsPostTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewTagsPostTagParamsWithContext(ctx context.Context) *TagsPostTagParams {
	var ()
	return &TagsPostTagParams{

		Context: ctx,
	}
}

// NewTagsPostTagParamsWithHTTPClient creates a new TagsPostTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTagsPostTagParamsWithHTTPClient(client *http.Client) *TagsPostTagParams {
	var ()
	return &TagsPostTagParams{
		HTTPClient: client,
	}
}

/*TagsPostTagParams contains all the parameters to send to the API endpoint
for the tags post tag operation typically these are written to a http.Request
*/
type TagsPostTagParams struct {

	/*AccountID*/
	AccountID int32
	/*ClientTag*/
	ClientTag *models.Tag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tags post tag params
func (o *TagsPostTagParams) WithTimeout(timeout time.Duration) *TagsPostTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tags post tag params
func (o *TagsPostTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tags post tag params
func (o *TagsPostTagParams) WithContext(ctx context.Context) *TagsPostTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tags post tag params
func (o *TagsPostTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tags post tag params
func (o *TagsPostTagParams) WithHTTPClient(client *http.Client) *TagsPostTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tags post tag params
func (o *TagsPostTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the tags post tag params
func (o *TagsPostTagParams) WithAccountID(accountID int32) *TagsPostTagParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the tags post tag params
func (o *TagsPostTagParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithClientTag adds the clientTag to the tags post tag params
func (o *TagsPostTagParams) WithClientTag(clientTag *models.Tag) *TagsPostTagParams {
	o.SetClientTag(clientTag)
	return o
}

// SetClientTag adds the clientTag to the tags post tag params
func (o *TagsPostTagParams) SetClientTag(clientTag *models.Tag) {
	o.ClientTag = clientTag
}

// WriteToRequest writes these params to a swagger request
func (o *TagsPostTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	if o.ClientTag != nil {
		if err := r.SetBodyParam(o.ClientTag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}