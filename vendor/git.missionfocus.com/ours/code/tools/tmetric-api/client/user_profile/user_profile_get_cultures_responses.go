// Code generated by go-swagger; DO NOT EDIT.

package user_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/tools/tmetric-api/models"
)

// UserProfileGetCulturesReader is a Reader for the UserProfileGetCultures structure.
type UserProfileGetCulturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserProfileGetCulturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserProfileGetCulturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserProfileGetCulturesOK creates a UserProfileGetCulturesOK with default headers values
func NewUserProfileGetCulturesOK() *UserProfileGetCulturesOK {
	return &UserProfileGetCulturesOK{}
}

/*UserProfileGetCulturesOK handles this case with default header values.

OK
*/
type UserProfileGetCulturesOK struct {
	Payload []*models.CultureInfoLite
}

func (o *UserProfileGetCulturesOK) Error() string {
	return fmt.Sprintf("[GET /api/userprofile/cultures][%d] userProfileGetCulturesOK  %+v", 200, o.Payload)
}

func (o *UserProfileGetCulturesOK) GetPayload() []*models.CultureInfoLite {
	return o.Payload
}

func (o *UserProfileGetCulturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
