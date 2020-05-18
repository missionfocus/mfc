// Code generated by go-swagger; DO NOT EDIT.

package account_members

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/tools/tmetric-api/models"
)

// AccountMembersPostAccountMembersReader is a Reader for the AccountMembersPostAccountMembers structure.
type AccountMembersPostAccountMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountMembersPostAccountMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountMembersPostAccountMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAccountMembersPostAccountMembersOK creates a AccountMembersPostAccountMembersOK with default headers values
func NewAccountMembersPostAccountMembersOK() *AccountMembersPostAccountMembersOK {
	return &AccountMembersPostAccountMembersOK{}
}

/*AccountMembersPostAccountMembersOK handles this case with default header values.

OK
*/
type AccountMembersPostAccountMembersOK struct {
	Payload []*models.AccountMember
}

func (o *AccountMembersPostAccountMembersOK) Error() string {
	return fmt.Sprintf("[POST /api/accounts/{accountId}/members/bulk][%d] accountMembersPostAccountMembersOK  %+v", 200, o.Payload)
}

func (o *AccountMembersPostAccountMembersOK) GetPayload() []*models.AccountMember {
	return o.Payload
}

func (o *AccountMembersPostAccountMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
