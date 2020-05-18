// Code generated by go-swagger; DO NOT EDIT.

package time_entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/tools/tmetric-api/models"
)

// TimeEntriesGetGroupTimeEntriesReader is a Reader for the TimeEntriesGetGroupTimeEntries structure.
type TimeEntriesGetGroupTimeEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimeEntriesGetGroupTimeEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimeEntriesGetGroupTimeEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimeEntriesGetGroupTimeEntriesOK creates a TimeEntriesGetGroupTimeEntriesOK with default headers values
func NewTimeEntriesGetGroupTimeEntriesOK() *TimeEntriesGetGroupTimeEntriesOK {
	return &TimeEntriesGetGroupTimeEntriesOK{}
}

/*TimeEntriesGetGroupTimeEntriesOK handles this case with default header values.

OK
*/
type TimeEntriesGetGroupTimeEntriesOK struct {
	Payload []*models.GroupTimeEntries
}

func (o *TimeEntriesGetGroupTimeEntriesOK) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{accountId}/timeentries/group][%d] timeEntriesGetGroupTimeEntriesOK  %+v", 200, o.Payload)
}

func (o *TimeEntriesGetGroupTimeEntriesOK) GetPayload() []*models.GroupTimeEntries {
	return o.Payload
}

func (o *TimeEntriesGetGroupTimeEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
