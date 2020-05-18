// Code generated by go-swagger; DO NOT EDIT.

package time_entries

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

// NewTimeEntriesGetGroupTimeEntriesByUserParams creates a new TimeEntriesGetGroupTimeEntriesByUserParams object
// with the default values initialized.
func NewTimeEntriesGetGroupTimeEntriesByUserParams() *TimeEntriesGetGroupTimeEntriesByUserParams {
	var ()
	return &TimeEntriesGetGroupTimeEntriesByUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTimeEntriesGetGroupTimeEntriesByUserParamsWithTimeout creates a new TimeEntriesGetGroupTimeEntriesByUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTimeEntriesGetGroupTimeEntriesByUserParamsWithTimeout(timeout time.Duration) *TimeEntriesGetGroupTimeEntriesByUserParams {
	var ()
	return &TimeEntriesGetGroupTimeEntriesByUserParams{

		timeout: timeout,
	}
}

// NewTimeEntriesGetGroupTimeEntriesByUserParamsWithContext creates a new TimeEntriesGetGroupTimeEntriesByUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewTimeEntriesGetGroupTimeEntriesByUserParamsWithContext(ctx context.Context) *TimeEntriesGetGroupTimeEntriesByUserParams {
	var ()
	return &TimeEntriesGetGroupTimeEntriesByUserParams{

		Context: ctx,
	}
}

// NewTimeEntriesGetGroupTimeEntriesByUserParamsWithHTTPClient creates a new TimeEntriesGetGroupTimeEntriesByUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTimeEntriesGetGroupTimeEntriesByUserParamsWithHTTPClient(client *http.Client) *TimeEntriesGetGroupTimeEntriesByUserParams {
	var ()
	return &TimeEntriesGetGroupTimeEntriesByUserParams{
		HTTPClient: client,
	}
}

/*TimeEntriesGetGroupTimeEntriesByUserParams contains all the parameters to send to the API endpoint
for the time entries get group time entries by user operation typically these are written to a http.Request
*/
type TimeEntriesGetGroupTimeEntriesByUserParams struct {

	/*AccountID*/
	AccountID int32
	/*TimeRangeEndTime*/
	TimeRangeEndTime *strfmt.DateTime
	/*TimeRangeStartTime*/
	TimeRangeStartTime *strfmt.DateTime
	/*UseUtcTime*/
	UseUtcTime *bool
	/*UserGroupID*/
	UserGroupID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) WithTimeout(timeout time.Duration) *TimeEntriesGetGroupTimeEntriesByUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) WithContext(ctx context.Context) *TimeEntriesGetGroupTimeEntriesByUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) WithHTTPClient(client *http.Client) *TimeEntriesGetGroupTimeEntriesByUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) WithAccountID(accountID int32) *TimeEntriesGetGroupTimeEntriesByUserParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithTimeRangeEndTime adds the timeRangeEndTime to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) WithTimeRangeEndTime(timeRangeEndTime *strfmt.DateTime) *TimeEntriesGetGroupTimeEntriesByUserParams {
	o.SetTimeRangeEndTime(timeRangeEndTime)
	return o
}

// SetTimeRangeEndTime adds the timeRangeEndTime to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) SetTimeRangeEndTime(timeRangeEndTime *strfmt.DateTime) {
	o.TimeRangeEndTime = timeRangeEndTime
}

// WithTimeRangeStartTime adds the timeRangeStartTime to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) WithTimeRangeStartTime(timeRangeStartTime *strfmt.DateTime) *TimeEntriesGetGroupTimeEntriesByUserParams {
	o.SetTimeRangeStartTime(timeRangeStartTime)
	return o
}

// SetTimeRangeStartTime adds the timeRangeStartTime to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) SetTimeRangeStartTime(timeRangeStartTime *strfmt.DateTime) {
	o.TimeRangeStartTime = timeRangeStartTime
}

// WithUseUtcTime adds the useUtcTime to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) WithUseUtcTime(useUtcTime *bool) *TimeEntriesGetGroupTimeEntriesByUserParams {
	o.SetUseUtcTime(useUtcTime)
	return o
}

// SetUseUtcTime adds the useUtcTime to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) SetUseUtcTime(useUtcTime *bool) {
	o.UseUtcTime = useUtcTime
}

// WithUserGroupID adds the userGroupID to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) WithUserGroupID(userGroupID int32) *TimeEntriesGetGroupTimeEntriesByUserParams {
	o.SetUserGroupID(userGroupID)
	return o
}

// SetUserGroupID adds the userGroupId to the time entries get group time entries by user params
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) SetUserGroupID(userGroupID int32) {
	o.UserGroupID = userGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *TimeEntriesGetGroupTimeEntriesByUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	if o.TimeRangeEndTime != nil {

		// query param timeRange.endTime
		var qrTimeRangeEndTime strfmt.DateTime
		if o.TimeRangeEndTime != nil {
			qrTimeRangeEndTime = *o.TimeRangeEndTime
		}
		qTimeRangeEndTime := qrTimeRangeEndTime.String()
		if qTimeRangeEndTime != "" {
			if err := r.SetQueryParam("timeRange.endTime", qTimeRangeEndTime); err != nil {
				return err
			}
		}

	}

	if o.TimeRangeStartTime != nil {

		// query param timeRange.startTime
		var qrTimeRangeStartTime strfmt.DateTime
		if o.TimeRangeStartTime != nil {
			qrTimeRangeStartTime = *o.TimeRangeStartTime
		}
		qTimeRangeStartTime := qrTimeRangeStartTime.String()
		if qTimeRangeStartTime != "" {
			if err := r.SetQueryParam("timeRange.startTime", qTimeRangeStartTime); err != nil {
				return err
			}
		}

	}

	if o.UseUtcTime != nil {

		// query param useUtcTime
		var qrUseUtcTime bool
		if o.UseUtcTime != nil {
			qrUseUtcTime = *o.UseUtcTime
		}
		qUseUtcTime := swag.FormatBool(qrUseUtcTime)
		if qUseUtcTime != "" {
			if err := r.SetQueryParam("useUtcTime", qUseUtcTime); err != nil {
				return err
			}
		}

	}

	// path param userGroupId
	if err := r.SetPathParam("userGroupId", swag.FormatInt32(o.UserGroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
