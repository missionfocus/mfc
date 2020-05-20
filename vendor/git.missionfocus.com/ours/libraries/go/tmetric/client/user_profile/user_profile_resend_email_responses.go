// Code generated by go-swagger; DO NOT EDIT.

package user_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UserProfileResendEmailReader is a Reader for the UserProfileResendEmail structure.
type UserProfileResendEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserProfileResendEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserProfileResendEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserProfileResendEmailOK creates a UserProfileResendEmailOK with default headers values
func NewUserProfileResendEmailOK() *UserProfileResendEmailOK {
	return &UserProfileResendEmailOK{}
}

/*UserProfileResendEmailOK handles this case with default header values.

OK
*/
type UserProfileResendEmailOK struct {
	Payload interface{}
}

func (o *UserProfileResendEmailOK) Error() string {
	return fmt.Sprintf("[POST /api/userprofile/resend][%d] userProfileResendEmailOK  %+v", 200, o.Payload)
}

func (o *UserProfileResendEmailOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UserProfileResendEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
