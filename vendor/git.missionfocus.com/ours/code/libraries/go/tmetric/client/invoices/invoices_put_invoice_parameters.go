// Code generated by go-swagger; DO NOT EDIT.

package invoices

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

// NewInvoicesPutInvoiceParams creates a new InvoicesPutInvoiceParams object
// with the default values initialized.
func NewInvoicesPutInvoiceParams() *InvoicesPutInvoiceParams {
	var ()
	return &InvoicesPutInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInvoicesPutInvoiceParamsWithTimeout creates a new InvoicesPutInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInvoicesPutInvoiceParamsWithTimeout(timeout time.Duration) *InvoicesPutInvoiceParams {
	var ()
	return &InvoicesPutInvoiceParams{

		timeout: timeout,
	}
}

// NewInvoicesPutInvoiceParamsWithContext creates a new InvoicesPutInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewInvoicesPutInvoiceParamsWithContext(ctx context.Context) *InvoicesPutInvoiceParams {
	var ()
	return &InvoicesPutInvoiceParams{

		Context: ctx,
	}
}

// NewInvoicesPutInvoiceParamsWithHTTPClient creates a new InvoicesPutInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInvoicesPutInvoiceParamsWithHTTPClient(client *http.Client) *InvoicesPutInvoiceParams {
	var ()
	return &InvoicesPutInvoiceParams{
		HTTPClient: client,
	}
}

/*InvoicesPutInvoiceParams contains all the parameters to send to the API endpoint
for the invoices put invoice operation typically these are written to a http.Request
*/
type InvoicesPutInvoiceParams struct {

	/*AccountID*/
	AccountID int32
	/*Invoice*/
	Invoice *models.Invoice
	/*InvoiceID*/
	InvoiceID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) WithTimeout(timeout time.Duration) *InvoicesPutInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) WithContext(ctx context.Context) *InvoicesPutInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) WithHTTPClient(client *http.Client) *InvoicesPutInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) WithAccountID(accountID int32) *InvoicesPutInvoiceParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) SetAccountID(accountID int32) {
	o.AccountID = accountID
}

// WithInvoice adds the invoice to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) WithInvoice(invoice *models.Invoice) *InvoicesPutInvoiceParams {
	o.SetInvoice(invoice)
	return o
}

// SetInvoice adds the invoice to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) SetInvoice(invoice *models.Invoice) {
	o.Invoice = invoice
}

// WithInvoiceID adds the invoiceID to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) WithInvoiceID(invoiceID int32) *InvoicesPutInvoiceParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the invoices put invoice params
func (o *InvoicesPutInvoiceParams) SetInvoiceID(invoiceID int32) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *InvoicesPutInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", swag.FormatInt32(o.AccountID)); err != nil {
		return err
	}

	if o.Invoice != nil {
		if err := r.SetBodyParam(o.Invoice); err != nil {
			return err
		}
	}

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", swag.FormatInt32(o.InvoiceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}