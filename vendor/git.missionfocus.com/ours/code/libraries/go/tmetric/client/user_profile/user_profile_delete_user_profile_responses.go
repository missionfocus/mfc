// Code generated by go-swagger; DO NOT EDIT.

package user_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UserProfileDeleteUserProfileReader is a Reader for the UserProfileDeleteUserProfile structure.
type UserProfileDeleteUserProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserProfileDeleteUserProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUserProfileDeleteUserProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserProfileDeleteUserProfileNoContent creates a UserProfileDeleteUserProfileNoContent with default headers values
func NewUserProfileDeleteUserProfileNoContent() *UserProfileDeleteUserProfileNoContent {
	return &UserProfileDeleteUserProfileNoContent{}
}

/*UserProfileDeleteUserProfileNoContent handles this case with default header values.

No Content
*/
type UserProfileDeleteUserProfileNoContent struct {
}

func (o *UserProfileDeleteUserProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/userprofile][%d] userProfileDeleteUserProfileNoContent ", 204)
}

func (o *UserProfileDeleteUserProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}