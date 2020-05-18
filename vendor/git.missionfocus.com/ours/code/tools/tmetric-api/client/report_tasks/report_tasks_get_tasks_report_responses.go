// Code generated by go-swagger; DO NOT EDIT.

package report_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/tools/tmetric-api/models"
)

// ReportTasksGetTasksReportReader is a Reader for the ReportTasksGetTasksReport structure.
type ReportTasksGetTasksReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportTasksGetTasksReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportTasksGetTasksReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReportTasksGetTasksReportOK creates a ReportTasksGetTasksReportOK with default headers values
func NewReportTasksGetTasksReportOK() *ReportTasksGetTasksReportOK {
	return &ReportTasksGetTasksReportOK{}
}

/*ReportTasksGetTasksReportOK handles this case with default header values.

OK
*/
type ReportTasksGetTasksReportOK struct {
	Payload []*models.TaskReportRow
}

func (o *ReportTasksGetTasksReportOK) Error() string {
	return fmt.Sprintf("[GET /api/reports/summary/tasks][%d] reportTasksGetTasksReportOK  %+v", 200, o.Payload)
}

func (o *ReportTasksGetTasksReportOK) GetPayload() []*models.TaskReportRow {
	return o.Payload
}

func (o *ReportTasksGetTasksReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
