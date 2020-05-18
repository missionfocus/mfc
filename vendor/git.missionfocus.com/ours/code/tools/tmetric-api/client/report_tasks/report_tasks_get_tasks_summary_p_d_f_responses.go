// Code generated by go-swagger; DO NOT EDIT.

package report_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReportTasksGetTasksSummaryPDFReader is a Reader for the ReportTasksGetTasksSummaryPDF structure.
type ReportTasksGetTasksSummaryPDFReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportTasksGetTasksSummaryPDFReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportTasksGetTasksSummaryPDFOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReportTasksGetTasksSummaryPDFOK creates a ReportTasksGetTasksSummaryPDFOK with default headers values
func NewReportTasksGetTasksSummaryPDFOK() *ReportTasksGetTasksSummaryPDFOK {
	return &ReportTasksGetTasksSummaryPDFOK{}
}

/*ReportTasksGetTasksSummaryPDFOK handles this case with default header values.

OK
*/
type ReportTasksGetTasksSummaryPDFOK struct {
	Payload interface{}
}

func (o *ReportTasksGetTasksSummaryPDFOK) Error() string {
	return fmt.Sprintf("[GET /api/reports/summary/tasks/pdf][%d] reportTasksGetTasksSummaryPDFOK  %+v", 200, o.Payload)
}

func (o *ReportTasksGetTasksSummaryPDFOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ReportTasksGetTasksSummaryPDFOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
