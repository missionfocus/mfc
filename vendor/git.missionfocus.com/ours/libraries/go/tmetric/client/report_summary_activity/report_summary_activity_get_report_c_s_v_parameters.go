// Code generated by go-swagger; DO NOT EDIT.

package report_summary_activity

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

// NewReportSummaryActivityGetReportCSVParams creates a new ReportSummaryActivityGetReportCSVParams object
// with the default values initialized.
func NewReportSummaryActivityGetReportCSVParams() *ReportSummaryActivityGetReportCSVParams {
	var ()
	return &ReportSummaryActivityGetReportCSVParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReportSummaryActivityGetReportCSVParamsWithTimeout creates a new ReportSummaryActivityGetReportCSVParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReportSummaryActivityGetReportCSVParamsWithTimeout(timeout time.Duration) *ReportSummaryActivityGetReportCSVParams {
	var ()
	return &ReportSummaryActivityGetReportCSVParams{

		timeout: timeout,
	}
}

// NewReportSummaryActivityGetReportCSVParamsWithContext creates a new ReportSummaryActivityGetReportCSVParams object
// with the default values initialized, and the ability to set a context for a request
func NewReportSummaryActivityGetReportCSVParamsWithContext(ctx context.Context) *ReportSummaryActivityGetReportCSVParams {
	var ()
	return &ReportSummaryActivityGetReportCSVParams{

		Context: ctx,
	}
}

// NewReportSummaryActivityGetReportCSVParamsWithHTTPClient creates a new ReportSummaryActivityGetReportCSVParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReportSummaryActivityGetReportCSVParamsWithHTTPClient(client *http.Client) *ReportSummaryActivityGetReportCSVParams {
	var ()
	return &ReportSummaryActivityGetReportCSVParams{
		HTTPClient: client,
	}
}

/*ReportSummaryActivityGetReportCSVParams contains all the parameters to send to the API endpoint
for the report summary activity get report c s v operation typically these are written to a http.Request
*/
type ReportSummaryActivityGetReportCSVParams struct {

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

// WithTimeout adds the timeout to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithTimeout(timeout time.Duration) *ReportSummaryActivityGetReportCSVParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithContext(ctx context.Context) *ReportSummaryActivityGetReportCSVParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithHTTPClient(client *http.Client) *ReportSummaryActivityGetReportCSVParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportParamsAccountID adds the reportParamsAccountID to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsAccountID(reportParamsAccountID *int32) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsAccountID(reportParamsAccountID)
	return o
}

// SetReportParamsAccountID adds the reportParamsAccountId to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsAccountID(reportParamsAccountID *int32) {
	o.ReportParamsAccountID = reportParamsAccountID
}

// WithReportParamsActiveProjectsOnly adds the reportParamsActiveProjectsOnly to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsActiveProjectsOnly(reportParamsActiveProjectsOnly *bool) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsActiveProjectsOnly(reportParamsActiveProjectsOnly)
	return o
}

// SetReportParamsActiveProjectsOnly adds the reportParamsActiveProjectsOnly to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsActiveProjectsOnly(reportParamsActiveProjectsOnly *bool) {
	o.ReportParamsActiveProjectsOnly = reportParamsActiveProjectsOnly
}

// WithReportParamsBillable adds the reportParamsBillable to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsBillable(reportParamsBillable *bool) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsBillable(reportParamsBillable)
	return o
}

// SetReportParamsBillable adds the reportParamsBillable to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsBillable(reportParamsBillable *bool) {
	o.ReportParamsBillable = reportParamsBillable
}

// WithReportParamsChartValue adds the reportParamsChartValue to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsChartValue(reportParamsChartValue *string) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsChartValue(reportParamsChartValue)
	return o
}

// SetReportParamsChartValue adds the reportParamsChartValue to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsChartValue(reportParamsChartValue *string) {
	o.ReportParamsChartValue = reportParamsChartValue
}

// WithReportParamsClientList adds the reportParamsClientList to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsClientList(reportParamsClientList []int32) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsClientList(reportParamsClientList)
	return o
}

// SetReportParamsClientList adds the reportParamsClientList to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsClientList(reportParamsClientList []int32) {
	o.ReportParamsClientList = reportParamsClientList
}

// WithReportParamsEndDate adds the reportParamsEndDate to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsEndDate(reportParamsEndDate *strfmt.DateTime) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsEndDate(reportParamsEndDate)
	return o
}

// SetReportParamsEndDate adds the reportParamsEndDate to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsEndDate(reportParamsEndDate *strfmt.DateTime) {
	o.ReportParamsEndDate = reportParamsEndDate
}

// WithReportParamsGroupColumnNames adds the reportParamsGroupColumnNames to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsGroupColumnNames(reportParamsGroupColumnNames []string) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsGroupColumnNames(reportParamsGroupColumnNames)
	return o
}

// SetReportParamsGroupColumnNames adds the reportParamsGroupColumnNames to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsGroupColumnNames(reportParamsGroupColumnNames []string) {
	o.ReportParamsGroupColumnNames = reportParamsGroupColumnNames
}

// WithReportParamsGroupList adds the reportParamsGroupList to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsGroupList(reportParamsGroupList []int32) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsGroupList(reportParamsGroupList)
	return o
}

// SetReportParamsGroupList adds the reportParamsGroupList to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsGroupList(reportParamsGroupList []int32) {
	o.ReportParamsGroupList = reportParamsGroupList
}

// WithReportParamsHiddenColumns adds the reportParamsHiddenColumns to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsHiddenColumns(reportParamsHiddenColumns []string) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsHiddenColumns(reportParamsHiddenColumns)
	return o
}

// SetReportParamsHiddenColumns adds the reportParamsHiddenColumns to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsHiddenColumns(reportParamsHiddenColumns []string) {
	o.ReportParamsHiddenColumns = reportParamsHiddenColumns
}

// WithReportParamsInvoiced adds the reportParamsInvoiced to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsInvoiced(reportParamsInvoiced *bool) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsInvoiced(reportParamsInvoiced)
	return o
}

// SetReportParamsInvoiced adds the reportParamsInvoiced to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsInvoiced(reportParamsInvoiced *bool) {
	o.ReportParamsInvoiced = reportParamsInvoiced
}

// WithReportParamsNoRounding adds the reportParamsNoRounding to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsNoRounding(reportParamsNoRounding *bool) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsNoRounding(reportParamsNoRounding)
	return o
}

// SetReportParamsNoRounding adds the reportParamsNoRounding to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsNoRounding(reportParamsNoRounding *bool) {
	o.ReportParamsNoRounding = reportParamsNoRounding
}

// WithReportParamsProfileList adds the reportParamsProfileList to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsProfileList(reportParamsProfileList []int32) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsProfileList(reportParamsProfileList)
	return o
}

// SetReportParamsProfileList adds the reportParamsProfileList to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsProfileList(reportParamsProfileList []int32) {
	o.ReportParamsProfileList = reportParamsProfileList
}

// WithReportParamsProjectList adds the reportParamsProjectList to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsProjectList(reportParamsProjectList []int32) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsProjectList(reportParamsProjectList)
	return o
}

// SetReportParamsProjectList adds the reportParamsProjectList to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsProjectList(reportParamsProjectList []int32) {
	o.ReportParamsProjectList = reportParamsProjectList
}

// WithReportParamsStartDate adds the reportParamsStartDate to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsStartDate(reportParamsStartDate *strfmt.DateTime) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsStartDate(reportParamsStartDate)
	return o
}

// SetReportParamsStartDate adds the reportParamsStartDate to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsStartDate(reportParamsStartDate *strfmt.DateTime) {
	o.ReportParamsStartDate = reportParamsStartDate
}

// WithReportParamsTagList adds the reportParamsTagList to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsTagList(reportParamsTagList []int32) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsTagList(reportParamsTagList)
	return o
}

// SetReportParamsTagList adds the reportParamsTagList to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsTagList(reportParamsTagList []int32) {
	o.ReportParamsTagList = reportParamsTagList
}

// WithReportParamsTimeEntryFilter adds the reportParamsTimeEntryFilter to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsTimeEntryFilter(reportParamsTimeEntryFilter *string) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsTimeEntryFilter(reportParamsTimeEntryFilter)
	return o
}

// SetReportParamsTimeEntryFilter adds the reportParamsTimeEntryFilter to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsTimeEntryFilter(reportParamsTimeEntryFilter *string) {
	o.ReportParamsTimeEntryFilter = reportParamsTimeEntryFilter
}

// WithReportParamsUseUtcTime adds the reportParamsUseUtcTime to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) WithReportParamsUseUtcTime(reportParamsUseUtcTime *bool) *ReportSummaryActivityGetReportCSVParams {
	o.SetReportParamsUseUtcTime(reportParamsUseUtcTime)
	return o
}

// SetReportParamsUseUtcTime adds the reportParamsUseUtcTime to the report summary activity get report c s v params
func (o *ReportSummaryActivityGetReportCSVParams) SetReportParamsUseUtcTime(reportParamsUseUtcTime *bool) {
	o.ReportParamsUseUtcTime = reportParamsUseUtcTime
}

// WriteToRequest writes these params to a swagger request
func (o *ReportSummaryActivityGetReportCSVParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
