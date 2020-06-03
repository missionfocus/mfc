// Code generated by go-swagger; DO NOT EDIT.

package report_long_timers

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

// NewReportLongTimersGetLongTimersParams creates a new ReportLongTimersGetLongTimersParams object
// with the default values initialized.
func NewReportLongTimersGetLongTimersParams() *ReportLongTimersGetLongTimersParams {
	var ()
	return &ReportLongTimersGetLongTimersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReportLongTimersGetLongTimersParamsWithTimeout creates a new ReportLongTimersGetLongTimersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReportLongTimersGetLongTimersParamsWithTimeout(timeout time.Duration) *ReportLongTimersGetLongTimersParams {
	var ()
	return &ReportLongTimersGetLongTimersParams{

		timeout: timeout,
	}
}

// NewReportLongTimersGetLongTimersParamsWithContext creates a new ReportLongTimersGetLongTimersParams object
// with the default values initialized, and the ability to set a context for a request
func NewReportLongTimersGetLongTimersParamsWithContext(ctx context.Context) *ReportLongTimersGetLongTimersParams {
	var ()
	return &ReportLongTimersGetLongTimersParams{

		Context: ctx,
	}
}

// NewReportLongTimersGetLongTimersParamsWithHTTPClient creates a new ReportLongTimersGetLongTimersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReportLongTimersGetLongTimersParamsWithHTTPClient(client *http.Client) *ReportLongTimersGetLongTimersParams {
	var ()
	return &ReportLongTimersGetLongTimersParams{
		HTTPClient: client,
	}
}

/*ReportLongTimersGetLongTimersParams contains all the parameters to send to the API endpoint
for the report long timers get long timers operation typically these are written to a http.Request
*/
type ReportLongTimersGetLongTimersParams struct {

	/*ReportParamsAccountID
	  Gets or sets the account identifier.

	*/
	ReportParamsAccountID *int32
	/*ReportParamsActiveProjectsOnly
	  Gets or sets the value indicating that only active projects should be returned.

	*/
	ReportParamsActiveProjectsOnly *bool
	/*ReportParamsBillable
	  Gets or sets the value indicating which tasks should be returned: false - non-billable, true - billable, null - both.

	*/
	ReportParamsBillable *bool
	/*ReportParamsChartValue
	  Gets or sets the column that will be shown on chart.

	*/
	ReportParamsChartValue *string
	/*ReportParamsClientList
	  Gets or sets cluent list for report filter.

	*/
	ReportParamsClientList []int32
	/*ReportParamsEndDate
	  Gets or sets report end date.

	*/
	ReportParamsEndDate *strfmt.DateTime
	/*ReportParamsGroupColumnNames
	  Gets or sets the group column names.

	*/
	ReportParamsGroupColumnNames []string
	/*ReportParamsGroupList
	  Gets or sets group list for report filter.

	*/
	ReportParamsGroupList []int32
	/*ReportParamsHiddenColumns
	  Gets or sets the hidden column names.

	*/
	ReportParamsHiddenColumns []string
	/*ReportParamsInvoiced
	  Gets or sets the value indicating which tasks should be returned: false - ununvoiced, true - invoiced, null - both.

	*/
	ReportParamsInvoiced *bool
	/*ReportParamsNoRounding
	  Gets or sets the value indicating that rounding in report should be turned off.

	*/
	ReportParamsNoRounding *bool
	/*ReportParamsProfileList
	  Gets or sets profile list for report filter.

	*/
	ReportParamsProfileList []int32
	/*ReportParamsProjectList
	  Gets or sets project list for report filter.

	*/
	ReportParamsProjectList []int32
	/*ReportParamsStartDate
	  Gets or sets report start date.

	*/
	ReportParamsStartDate *strfmt.DateTime
	/*ReportParamsTagList
	  Gets or sets tag list for report filter.

	*/
	ReportParamsTagList []int32
	/*ReportParamsTimeEntryFilter
	  Gets or sets the string that will be found in time entry descriptions.

	*/
	ReportParamsTimeEntryFilter *string
	/*ReportParamsUseUtcTime
	  Gets or sets the value indicating that UTC time should be returned.

	*/
	ReportParamsUseUtcTime *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithTimeout(timeout time.Duration) *ReportLongTimersGetLongTimersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithContext(ctx context.Context) *ReportLongTimersGetLongTimersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithHTTPClient(client *http.Client) *ReportLongTimersGetLongTimersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportParamsAccountID adds the reportParamsAccountID to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsAccountID(reportParamsAccountID *int32) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsAccountID(reportParamsAccountID)
	return o
}

// SetReportParamsAccountID adds the reportParamsAccountId to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsAccountID(reportParamsAccountID *int32) {
	o.ReportParamsAccountID = reportParamsAccountID
}

// WithReportParamsActiveProjectsOnly adds the reportParamsActiveProjectsOnly to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsActiveProjectsOnly(reportParamsActiveProjectsOnly *bool) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsActiveProjectsOnly(reportParamsActiveProjectsOnly)
	return o
}

// SetReportParamsActiveProjectsOnly adds the reportParamsActiveProjectsOnly to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsActiveProjectsOnly(reportParamsActiveProjectsOnly *bool) {
	o.ReportParamsActiveProjectsOnly = reportParamsActiveProjectsOnly
}

// WithReportParamsBillable adds the reportParamsBillable to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsBillable(reportParamsBillable *bool) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsBillable(reportParamsBillable)
	return o
}

// SetReportParamsBillable adds the reportParamsBillable to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsBillable(reportParamsBillable *bool) {
	o.ReportParamsBillable = reportParamsBillable
}

// WithReportParamsChartValue adds the reportParamsChartValue to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsChartValue(reportParamsChartValue *string) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsChartValue(reportParamsChartValue)
	return o
}

// SetReportParamsChartValue adds the reportParamsChartValue to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsChartValue(reportParamsChartValue *string) {
	o.ReportParamsChartValue = reportParamsChartValue
}

// WithReportParamsClientList adds the reportParamsClientList to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsClientList(reportParamsClientList []int32) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsClientList(reportParamsClientList)
	return o
}

// SetReportParamsClientList adds the reportParamsClientList to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsClientList(reportParamsClientList []int32) {
	o.ReportParamsClientList = reportParamsClientList
}

// WithReportParamsEndDate adds the reportParamsEndDate to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsEndDate(reportParamsEndDate *strfmt.DateTime) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsEndDate(reportParamsEndDate)
	return o
}

// SetReportParamsEndDate adds the reportParamsEndDate to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsEndDate(reportParamsEndDate *strfmt.DateTime) {
	o.ReportParamsEndDate = reportParamsEndDate
}

// WithReportParamsGroupColumnNames adds the reportParamsGroupColumnNames to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsGroupColumnNames(reportParamsGroupColumnNames []string) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsGroupColumnNames(reportParamsGroupColumnNames)
	return o
}

// SetReportParamsGroupColumnNames adds the reportParamsGroupColumnNames to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsGroupColumnNames(reportParamsGroupColumnNames []string) {
	o.ReportParamsGroupColumnNames = reportParamsGroupColumnNames
}

// WithReportParamsGroupList adds the reportParamsGroupList to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsGroupList(reportParamsGroupList []int32) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsGroupList(reportParamsGroupList)
	return o
}

// SetReportParamsGroupList adds the reportParamsGroupList to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsGroupList(reportParamsGroupList []int32) {
	o.ReportParamsGroupList = reportParamsGroupList
}

// WithReportParamsHiddenColumns adds the reportParamsHiddenColumns to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsHiddenColumns(reportParamsHiddenColumns []string) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsHiddenColumns(reportParamsHiddenColumns)
	return o
}

// SetReportParamsHiddenColumns adds the reportParamsHiddenColumns to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsHiddenColumns(reportParamsHiddenColumns []string) {
	o.ReportParamsHiddenColumns = reportParamsHiddenColumns
}

// WithReportParamsInvoiced adds the reportParamsInvoiced to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsInvoiced(reportParamsInvoiced *bool) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsInvoiced(reportParamsInvoiced)
	return o
}

// SetReportParamsInvoiced adds the reportParamsInvoiced to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsInvoiced(reportParamsInvoiced *bool) {
	o.ReportParamsInvoiced = reportParamsInvoiced
}

// WithReportParamsNoRounding adds the reportParamsNoRounding to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsNoRounding(reportParamsNoRounding *bool) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsNoRounding(reportParamsNoRounding)
	return o
}

// SetReportParamsNoRounding adds the reportParamsNoRounding to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsNoRounding(reportParamsNoRounding *bool) {
	o.ReportParamsNoRounding = reportParamsNoRounding
}

// WithReportParamsProfileList adds the reportParamsProfileList to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsProfileList(reportParamsProfileList []int32) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsProfileList(reportParamsProfileList)
	return o
}

// SetReportParamsProfileList adds the reportParamsProfileList to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsProfileList(reportParamsProfileList []int32) {
	o.ReportParamsProfileList = reportParamsProfileList
}

// WithReportParamsProjectList adds the reportParamsProjectList to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsProjectList(reportParamsProjectList []int32) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsProjectList(reportParamsProjectList)
	return o
}

// SetReportParamsProjectList adds the reportParamsProjectList to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsProjectList(reportParamsProjectList []int32) {
	o.ReportParamsProjectList = reportParamsProjectList
}

// WithReportParamsStartDate adds the reportParamsStartDate to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsStartDate(reportParamsStartDate *strfmt.DateTime) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsStartDate(reportParamsStartDate)
	return o
}

// SetReportParamsStartDate adds the reportParamsStartDate to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsStartDate(reportParamsStartDate *strfmt.DateTime) {
	o.ReportParamsStartDate = reportParamsStartDate
}

// WithReportParamsTagList adds the reportParamsTagList to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsTagList(reportParamsTagList []int32) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsTagList(reportParamsTagList)
	return o
}

// SetReportParamsTagList adds the reportParamsTagList to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsTagList(reportParamsTagList []int32) {
	o.ReportParamsTagList = reportParamsTagList
}

// WithReportParamsTimeEntryFilter adds the reportParamsTimeEntryFilter to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsTimeEntryFilter(reportParamsTimeEntryFilter *string) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsTimeEntryFilter(reportParamsTimeEntryFilter)
	return o
}

// SetReportParamsTimeEntryFilter adds the reportParamsTimeEntryFilter to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsTimeEntryFilter(reportParamsTimeEntryFilter *string) {
	o.ReportParamsTimeEntryFilter = reportParamsTimeEntryFilter
}

// WithReportParamsUseUtcTime adds the reportParamsUseUtcTime to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) WithReportParamsUseUtcTime(reportParamsUseUtcTime *bool) *ReportLongTimersGetLongTimersParams {
	o.SetReportParamsUseUtcTime(reportParamsUseUtcTime)
	return o
}

// SetReportParamsUseUtcTime adds the reportParamsUseUtcTime to the report long timers get long timers params
func (o *ReportLongTimersGetLongTimersParams) SetReportParamsUseUtcTime(reportParamsUseUtcTime *bool) {
	o.ReportParamsUseUtcTime = reportParamsUseUtcTime
}

// WriteToRequest writes these params to a swagger request
func (o *ReportLongTimersGetLongTimersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReportParamsAccountID != nil {

		// query param reportParams.accountId
		var qrReportParamsAccountID int32
		if o.ReportParamsAccountID != nil {
			qrReportParamsAccountID = *o.ReportParamsAccountID
		}
		qReportParamsAccountID := swag.FormatInt32(qrReportParamsAccountID)
		if qReportParamsAccountID != "" {
			if err := r.SetQueryParam("reportParams.accountId", qReportParamsAccountID); err != nil {
				return err
			}
		}

	}

	if o.ReportParamsActiveProjectsOnly != nil {

		// query param reportParams.activeProjectsOnly
		var qrReportParamsActiveProjectsOnly bool
		if o.ReportParamsActiveProjectsOnly != nil {
			qrReportParamsActiveProjectsOnly = *o.ReportParamsActiveProjectsOnly
		}
		qReportParamsActiveProjectsOnly := swag.FormatBool(qrReportParamsActiveProjectsOnly)
		if qReportParamsActiveProjectsOnly != "" {
			if err := r.SetQueryParam("reportParams.activeProjectsOnly", qReportParamsActiveProjectsOnly); err != nil {
				return err
			}
		}

	}

	if o.ReportParamsBillable != nil {

		// query param reportParams.billable
		var qrReportParamsBillable bool
		if o.ReportParamsBillable != nil {
			qrReportParamsBillable = *o.ReportParamsBillable
		}
		qReportParamsBillable := swag.FormatBool(qrReportParamsBillable)
		if qReportParamsBillable != "" {
			if err := r.SetQueryParam("reportParams.billable", qReportParamsBillable); err != nil {
				return err
			}
		}

	}

	if o.ReportParamsChartValue != nil {

		// query param reportParams.chartValue
		var qrReportParamsChartValue string
		if o.ReportParamsChartValue != nil {
			qrReportParamsChartValue = *o.ReportParamsChartValue
		}
		qReportParamsChartValue := qrReportParamsChartValue
		if qReportParamsChartValue != "" {
			if err := r.SetQueryParam("reportParams.chartValue", qReportParamsChartValue); err != nil {
				return err
			}
		}

	}

	var valuesReportParamsClientList []string
	for _, v := range o.ReportParamsClientList {
		valuesReportParamsClientList = append(valuesReportParamsClientList, swag.FormatInt32(v))
	}

	joinedReportParamsClientList := swag.JoinByFormat(valuesReportParamsClientList, "multi")
	// query array param reportParams.clientList
	if err := r.SetQueryParam("reportParams.clientList", joinedReportParamsClientList...); err != nil {
		return err
	}

	if o.ReportParamsEndDate != nil {

		// query param reportParams.endDate
		var qrReportParamsEndDate strfmt.DateTime
		if o.ReportParamsEndDate != nil {
			qrReportParamsEndDate = *o.ReportParamsEndDate
		}
		qReportParamsEndDate := qrReportParamsEndDate.String()
		if qReportParamsEndDate != "" {
			if err := r.SetQueryParam("reportParams.endDate", qReportParamsEndDate); err != nil {
				return err
			}
		}

	}

	valuesReportParamsGroupColumnNames := o.ReportParamsGroupColumnNames

	joinedReportParamsGroupColumnNames := swag.JoinByFormat(valuesReportParamsGroupColumnNames, "multi")
	// query array param reportParams.groupColumnNames
	if err := r.SetQueryParam("reportParams.groupColumnNames", joinedReportParamsGroupColumnNames...); err != nil {
		return err
	}

	var valuesReportParamsGroupList []string
	for _, v := range o.ReportParamsGroupList {
		valuesReportParamsGroupList = append(valuesReportParamsGroupList, swag.FormatInt32(v))
	}

	joinedReportParamsGroupList := swag.JoinByFormat(valuesReportParamsGroupList, "multi")
	// query array param reportParams.groupList
	if err := r.SetQueryParam("reportParams.groupList", joinedReportParamsGroupList...); err != nil {
		return err
	}

	valuesReportParamsHiddenColumns := o.ReportParamsHiddenColumns

	joinedReportParamsHiddenColumns := swag.JoinByFormat(valuesReportParamsHiddenColumns, "multi")
	// query array param reportParams.hiddenColumns
	if err := r.SetQueryParam("reportParams.hiddenColumns", joinedReportParamsHiddenColumns...); err != nil {
		return err
	}

	if o.ReportParamsInvoiced != nil {

		// query param reportParams.invoiced
		var qrReportParamsInvoiced bool
		if o.ReportParamsInvoiced != nil {
			qrReportParamsInvoiced = *o.ReportParamsInvoiced
		}
		qReportParamsInvoiced := swag.FormatBool(qrReportParamsInvoiced)
		if qReportParamsInvoiced != "" {
			if err := r.SetQueryParam("reportParams.invoiced", qReportParamsInvoiced); err != nil {
				return err
			}
		}

	}

	if o.ReportParamsNoRounding != nil {

		// query param reportParams.noRounding
		var qrReportParamsNoRounding bool
		if o.ReportParamsNoRounding != nil {
			qrReportParamsNoRounding = *o.ReportParamsNoRounding
		}
		qReportParamsNoRounding := swag.FormatBool(qrReportParamsNoRounding)
		if qReportParamsNoRounding != "" {
			if err := r.SetQueryParam("reportParams.noRounding", qReportParamsNoRounding); err != nil {
				return err
			}
		}

	}

	var valuesReportParamsProfileList []string
	for _, v := range o.ReportParamsProfileList {
		valuesReportParamsProfileList = append(valuesReportParamsProfileList, swag.FormatInt32(v))
	}

	joinedReportParamsProfileList := swag.JoinByFormat(valuesReportParamsProfileList, "multi")
	// query array param reportParams.profileList
	if err := r.SetQueryParam("reportParams.profileList", joinedReportParamsProfileList...); err != nil {
		return err
	}

	var valuesReportParamsProjectList []string
	for _, v := range o.ReportParamsProjectList {
		valuesReportParamsProjectList = append(valuesReportParamsProjectList, swag.FormatInt32(v))
	}

	joinedReportParamsProjectList := swag.JoinByFormat(valuesReportParamsProjectList, "multi")
	// query array param reportParams.projectList
	if err := r.SetQueryParam("reportParams.projectList", joinedReportParamsProjectList...); err != nil {
		return err
	}

	if o.ReportParamsStartDate != nil {

		// query param reportParams.startDate
		var qrReportParamsStartDate strfmt.DateTime
		if o.ReportParamsStartDate != nil {
			qrReportParamsStartDate = *o.ReportParamsStartDate
		}
		qReportParamsStartDate := qrReportParamsStartDate.String()
		if qReportParamsStartDate != "" {
			if err := r.SetQueryParam("reportParams.startDate", qReportParamsStartDate); err != nil {
				return err
			}
		}

	}

	var valuesReportParamsTagList []string
	for _, v := range o.ReportParamsTagList {
		valuesReportParamsTagList = append(valuesReportParamsTagList, swag.FormatInt32(v))
	}

	joinedReportParamsTagList := swag.JoinByFormat(valuesReportParamsTagList, "multi")
	// query array param reportParams.tagList
	if err := r.SetQueryParam("reportParams.tagList", joinedReportParamsTagList...); err != nil {
		return err
	}

	if o.ReportParamsTimeEntryFilter != nil {

		// query param reportParams.timeEntryFilter
		var qrReportParamsTimeEntryFilter string
		if o.ReportParamsTimeEntryFilter != nil {
			qrReportParamsTimeEntryFilter = *o.ReportParamsTimeEntryFilter
		}
		qReportParamsTimeEntryFilter := qrReportParamsTimeEntryFilter
		if qReportParamsTimeEntryFilter != "" {
			if err := r.SetQueryParam("reportParams.timeEntryFilter", qReportParamsTimeEntryFilter); err != nil {
				return err
			}
		}

	}

	if o.ReportParamsUseUtcTime != nil {

		// query param reportParams.useUtcTime
		var qrReportParamsUseUtcTime bool
		if o.ReportParamsUseUtcTime != nil {
			qrReportParamsUseUtcTime = *o.ReportParamsUseUtcTime
		}
		qReportParamsUseUtcTime := swag.FormatBool(qrReportParamsUseUtcTime)
		if qReportParamsUseUtcTime != "" {
			if err := r.SetQueryParam("reportParams.useUtcTime", qReportParamsUseUtcTime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}