// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AccountsDeleteAccountReader is a Reader for the AccountsDeleteAccount structure.
type AccountsDeleteAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountsDeleteAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAccountsDeleteAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAccountsDeleteAccountNoContent creates a AccountsDeleteAccountNoContent with default headers values
func NewAccountsDeleteAccountNoContent() *AccountsDeleteAccountNoContent {
	return &AccountsDeleteAccountNoContent{}
}

/*AccountsDeleteAccountNoContent handles this case with default header values.

No Content
*/
type AccountsDeleteAccountNoContent struct {
}

func (o *AccountsDeleteAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/accounts/{accountId}][%d] accountsDeleteAccountNoContent ", 204)
}

func (o *AccountsDeleteAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
