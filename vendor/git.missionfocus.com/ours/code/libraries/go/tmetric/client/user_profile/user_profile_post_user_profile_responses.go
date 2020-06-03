// Code generated by go-swagger; DO NOT EDIT.

package user_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/libraries/go/tmetric/models"
)

// UserProfilePostUserProfileReader is a Reader for the UserProfilePostUserProfile structure.
type UserProfilePostUserProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserProfilePostUserProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserProfilePostUserProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserProfilePostUserProfileOK creates a UserProfilePostUserProfileOK with default headers values
func NewUserProfilePostUserProfileOK() *UserProfilePostUserProfileOK {
	return &UserProfilePostUserProfileOK{}
}

/*UserProfilePostUserProfileOK handles this case with default header values.

OK
*/
type UserProfilePostUserProfileOK struct {
	Payload *models.UserProfile
}

func (o *UserProfilePostUserProfileOK) Error() string {
	return fmt.Sprintf("[POST /api/userprofile][%d] userProfilePostUserProfileOK  %+v", 200, o.Payload)
}

func (o *UserProfilePostUserProfileOK) GetPayload() *models.UserProfile {
	return o.Payload
}

func (o *UserProfilePostUserProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}