// Code generated by go-swagger; DO NOT EDIT.

package timeline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TimelinePostTimelineEntriesReader is a Reader for the TimelinePostTimelineEntries structure.
type TimelinePostTimelineEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimelinePostTimelineEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimelinePostTimelineEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimelinePostTimelineEntriesOK creates a TimelinePostTimelineEntriesOK with default headers values
func NewTimelinePostTimelineEntriesOK() *TimelinePostTimelineEntriesOK {
	return &TimelinePostTimelineEntriesOK{}
}

/*TimelinePostTimelineEntriesOK handles this case with default header values.

OK
*/
type TimelinePostTimelineEntriesOK struct {
	Payload interface{}
}

func (o *TimelinePostTimelineEntriesOK) Error() string {
	return fmt.Sprintf("[POST /api/timeline/{accountId}][%d] timelinePostTimelineEntriesOK  %+v", 200, o.Payload)
}

func (o *TimelinePostTimelineEntriesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TimelinePostTimelineEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
