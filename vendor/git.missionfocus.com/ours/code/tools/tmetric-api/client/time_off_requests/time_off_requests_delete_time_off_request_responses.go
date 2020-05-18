// Code generated by go-swagger; DO NOT EDIT.

package time_off_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TimeOffRequestsDeleteTimeOffRequestReader is a Reader for the TimeOffRequestsDeleteTimeOffRequest structure.
type TimeOffRequestsDeleteTimeOffRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimeOffRequestsDeleteTimeOffRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTimeOffRequestsDeleteTimeOffRequestNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimeOffRequestsDeleteTimeOffRequestNoContent creates a TimeOffRequestsDeleteTimeOffRequestNoContent with default headers values
func NewTimeOffRequestsDeleteTimeOffRequestNoContent() *TimeOffRequestsDeleteTimeOffRequestNoContent {
	return &TimeOffRequestsDeleteTimeOffRequestNoContent{}
}

/*TimeOffRequestsDeleteTimeOffRequestNoContent handles this case with default header values.

No Content
*/
type TimeOffRequestsDeleteTimeOffRequestNoContent struct {
}

func (o *TimeOffRequestsDeleteTimeOffRequestNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/accounts/{accountId}/timeoff/requests/{timeOffRequestId}][%d] timeOffRequestsDeleteTimeOffRequestNoContent ", 204)
}

func (o *TimeOffRequestsDeleteTimeOffRequestNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
