// Code generated by go-swagger; DO NOT EDIT.

package account_members

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AccountMembersPostAccountMemberReader is a Reader for the AccountMembersPostAccountMember structure.
type AccountMembersPostAccountMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountMembersPostAccountMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAccountMembersPostAccountMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAccountMembersPostAccountMemberNoContent creates a AccountMembersPostAccountMemberNoContent with default headers values
func NewAccountMembersPostAccountMemberNoContent() *AccountMembersPostAccountMemberNoContent {
	return &AccountMembersPostAccountMemberNoContent{}
}

/*AccountMembersPostAccountMemberNoContent handles this case with default header values.

No Content
*/
type AccountMembersPostAccountMemberNoContent struct {
}

func (o *AccountMembersPostAccountMemberNoContent) Error() string {
	return fmt.Sprintf("[POST /api/accounts/{accountId}/members][%d] accountMembersPostAccountMemberNoContent ", 204)
}

func (o *AccountMembersPostAccountMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}