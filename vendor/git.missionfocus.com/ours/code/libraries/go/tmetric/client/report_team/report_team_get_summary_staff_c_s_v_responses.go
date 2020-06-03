// Code generated by go-swagger; DO NOT EDIT.

package report_team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReportTeamGetSummaryStaffCSVReader is a Reader for the ReportTeamGetSummaryStaffCSV structure.
type ReportTeamGetSummaryStaffCSVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportTeamGetSummaryStaffCSVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportTeamGetSummaryStaffCSVOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReportTeamGetSummaryStaffCSVOK creates a ReportTeamGetSummaryStaffCSVOK with default headers values
func NewReportTeamGetSummaryStaffCSVOK() *ReportTeamGetSummaryStaffCSVOK {
	return &ReportTeamGetSummaryStaffCSVOK{}
}

/*ReportTeamGetSummaryStaffCSVOK handles this case with default header values.

OK
*/
type ReportTeamGetSummaryStaffCSVOK struct {
	Payload interface{}
}

func (o *ReportTeamGetSummaryStaffCSVOK) Error() string {
	return fmt.Sprintf("[GET /api/reports/summary/staff/csv][%d] reportTeamGetSummaryStaffCSVOK  %+v", 200, o.Payload)
}

func (o *ReportTeamGetSummaryStaffCSVOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ReportTeamGetSummaryStaffCSVOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}