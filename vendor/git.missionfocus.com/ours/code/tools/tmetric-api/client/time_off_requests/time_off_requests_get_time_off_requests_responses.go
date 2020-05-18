// Code generated by go-swagger; DO NOT EDIT.

package time_off_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/tools/tmetric-api/models"
)

// TimeOffRequestsGetTimeOffRequestsReader is a Reader for the TimeOffRequestsGetTimeOffRequests structure.
type TimeOffRequestsGetTimeOffRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimeOffRequestsGetTimeOffRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimeOffRequestsGetTimeOffRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimeOffRequestsGetTimeOffRequestsOK creates a TimeOffRequestsGetTimeOffRequestsOK with default headers values
func NewTimeOffRequestsGetTimeOffRequestsOK() *TimeOffRequestsGetTimeOffRequestsOK {
	return &TimeOffRequestsGetTimeOffRequestsOK{}
}

/*TimeOffRequestsGetTimeOffRequestsOK handles this case with default header values.

OK
*/
type TimeOffRequestsGetTimeOffRequestsOK struct {
	Payload []*models.TimeOffRequest
}

func (o *TimeOffRequestsGetTimeOffRequestsOK) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{accountId}/timeoff/requests][%d] timeOffRequestsGetTimeOffRequestsOK  %+v", 200, o.Payload)
}

func (o *TimeOffRequestsGetTimeOffRequestsOK) GetPayload() []*models.TimeOffRequest {
	return o.Payload
}

func (o *TimeOffRequestsGetTimeOffRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
