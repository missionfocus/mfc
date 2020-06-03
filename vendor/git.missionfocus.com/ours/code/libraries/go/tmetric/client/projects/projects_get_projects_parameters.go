// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewProjectsGetProjectsParams creates a new ProjectsGetProjectsParams object
// with the default values initialized.
func NewProjectsGetProjectsParams() *ProjectsGetProjectsParams {
	var ()
	return &ProjectsGetProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsGetProjectsParamsWithTimeout creates a new ProjectsGetProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsGetProjectsParamsWithTimeout(timeout time.Duration) *ProjectsGetProjectsParams {
	var ()
	return &ProjectsGetProjectsParams{

		timeout: timeout,
	}
}

// NewProjectsGetProjectsParamsWithContext creates a new ProjectsGetProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsGetProjectsParamsWithContext(ctx context.Context) *ProjectsGetProjectsParams {
	var ()
	return &ProjectsGetProjectsParams{

		Context: ctx,
	}
}

// NewProjectsGetProjectsParamsWithHTTPClient creates a new ProjectsGetProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsGetProjectsParamsWithHTTPClient(client *http.Client) *ProjectsGetProjectsParams {
	var ()
	return &ProjectsGetProjectsParams{
		HTTPClient: client,
	}
}

/*ProjectsGetProjectsParams contains all the parameters to send to the API endpoint
for the projects get projects operation typically these are written to a http.Request
*/
type ProjectsGetProjectsParams struct {

	/*AccountID
	  The account identifier.

	*/
	AccountID int32
	/*FilterBilling
	  Project billing filter

	*/
	FilterBilling *int32
	/*FilterBudget
	  Project budget filter

	*/
	FilterBudget *int32
	/*FilterClientList
	  Gets or sets client list for project filter.

	*/
	FilterClientList []int32
	/*FilterStatus
	  Project status filter

	*/
	FilterStatus *int32
	/*OnlyTracked
	  If true, ruturns only projects in which user can track time.

	*/
	OnlyTracked *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects get projects params
func (o *ProjectsGetProjectsParams) WithTimeout(timeout time.Duration) *ProjectsGetProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects get projects params
func (o *ProjectsGetProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects get projects params
func (o *ProjectsGetProjectsParams) WithContext(ctx context.Context) *ProjectsGetProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects get projects params
func (o *ProjectsGetProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects get projects params
func (o *ProjectsGetProjectsParams) WithHTTPClient(client *http.Client) *ProjectsGetProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects get projects params
func (o *ProjectsGetProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the projects get projects params
func (o *ProjectsGetProjectsParams) WithAccountID(accountID int32) *ProjectsGetProjectsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the projects get projects params
func (o *ProjectsGetProjectsParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithFilterBilling adds the filterBilling to the projects get projects params
func (o *ProjectsGetProjectsParams) WithFilterBilling(filterBilling *int32) *ProjectsGetProjectsParams {
	o.SetFilterBilling(filterBilling)
	return o
}

// SetFilterBilling adds the filterBilling to the projects get projects params
func (o *ProjectsGetProjectsParams) SetFilterBilling(filterBilling *int32) {
	o.FilterBilling = filterBilling
}

// WithFilterBudget adds the filterBudget to the projects get projects params
func (o *ProjectsGetProjectsParams) WithFilterBudget(filterBudget *int32) *ProjectsGetProjectsParams {
	o.SetFilterBudget(filterBudget)
	return o
}

// SetFilterBudget adds the filterBudget to the projects get projects params
func (o *ProjectsGetProjectsParams) SetFilterBudget(filterBudget *int32) {
	o.FilterBudget = filterBudget
}

// WithFilterClientList adds the filterClientList to the projects get projects params
func (o *ProjectsGetProjectsParams) WithFilterClientList(filterClientList []int32) *ProjectsGetProjectsParams {
	o.SetFilterClientList(filterClientList)
	return o
}

// SetFilterClientList adds the filterClientList to the projects get projects params
func (o *ProjectsGetProjectsParams) SetFilterClientList(filterClientList []int32) {
	o.FilterClientList = filterClientList
}

// WithFilterStatus adds the filterStatus to the projects get projects params
func (o *ProjectsGetProjectsParams) WithFilterStatus(filterStatus *int32) *ProjectsGetProjectsParams {
	o.SetFilterStatus(filterStatus)
	return o
}

// SetFilterStatus adds the filterStatus to the projects get projects params
func (o *ProjectsGetProjectsParams) SetFilterStatus(filterStatus *int32) {
	o.FilterStatus = filterStatus
}

// WithOnlyTracked adds the onlyTracked to the projects get projects params
func (o *ProjectsGetProjectsParams) WithOnlyTracked(onlyTracked *bool) *ProjectsGetProjectsParams {
	o.SetOnlyTracked(onlyTracked)
	return o
}

// SetOnlyTracked adds the onlyTracked to the projects get projects params
func (o *ProjectsGetProjectsParams) SetOnlyTracked(onlyTracked *bool) {
	o.OnlyTracked = onlyTracked
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsGetProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	if o.FilterBilling != nil {

		// query param filter.billing
		var qrFilterBilling int32
		if o.FilterBilling != nil {
			qrFilterBilling = *o.FilterBilling
		}
		qFilterBilling := swag.FormatInt32(qrFilterBilling)
		if qFilterBilling != "" {
			if err := r.SetQueryParam("filter.billing", qFilterBilling); err != nil {
				return err
			}
		}

	}

	if o.FilterBudget != nil {

		// query param filter.budget
		var qrFilterBudget int32
		if o.FilterBudget != nil {
			qrFilterBudget = *o.FilterBudget
		}
		qFilterBudget := swag.FormatInt32(qrFilterBudget)
		if qFilterBudget != "" {
			if err := r.SetQueryParam("filter.budget", qFilterBudget); err != nil {
				return err
			}
		}

	}

	var valuesFilterClientList []string
	for _, v := range o.FilterClientList {
		valuesFilterClientList = append(valuesFilterClientList, swag.FormatInt32(v))
	}

	joinedFilterClientList := swag.JoinByFormat(valuesFilterClientList, "multi")
	// query array param filter.clientList
	if err := r.SetQueryParam("filter.clientList", joinedFilterClientList...); err != nil {
		return err
	}

	if o.FilterStatus != nil {

		// query param filter.status
		var qrFilterStatus int32
		if o.FilterStatus != nil {
			qrFilterStatus = *o.FilterStatus
		}
		qFilterStatus := swag.FormatInt32(qrFilterStatus)
		if qFilterStatus != "" {
			if err := r.SetQueryParam("filter.status", qFilterStatus); err != nil {
				return err
			}
		}

	}

	if o.OnlyTracked != nil {

		// query param onlyTracked
		var qrOnlyTracked bool
		if o.OnlyTracked != nil {
			qrOnlyTracked = *o.OnlyTracked
		}
		qOnlyTracked := swag.FormatBool(qrOnlyTracked)
		if qOnlyTracked != "" {
			if err := r.SetQueryParam("onlyTracked", qOnlyTracked); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}