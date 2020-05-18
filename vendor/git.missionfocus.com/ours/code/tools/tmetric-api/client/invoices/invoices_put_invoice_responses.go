// Code generated by go-swagger; DO NOT EDIT.

package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InvoicesPutInvoiceReader is a Reader for the InvoicesPutInvoice structure.
type InvoicesPutInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvoicesPutInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewInvoicesPutInvoiceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInvoicesPutInvoiceNoContent creates a InvoicesPutInvoiceNoContent with default headers values
func NewInvoicesPutInvoiceNoContent() *InvoicesPutInvoiceNoContent {
	return &InvoicesPutInvoiceNoContent{}
}

/*InvoicesPutInvoiceNoContent handles this case with default header values.

No Content
*/
type InvoicesPutInvoiceNoContent struct {
}

func (o *InvoicesPutInvoiceNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/accounts/{accountId}/invoices/{invoiceId}][%d] invoicesPutInvoiceNoContent ", 204)
}

func (o *InvoicesPutInvoiceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
