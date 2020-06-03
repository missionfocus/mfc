// Code generated by go-swagger; DO NOT EDIT.

package report_summary_activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReportSummaryActivityGetReportPDFReader is a Reader for the ReportSummaryActivityGetReportPDF structure.
type ReportSummaryActivityGetReportPDFReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportSummaryActivityGetReportPDFReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportSummaryActivityGetReportPDFOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReportSummaryActivityGetReportPDFOK creates a ReportSummaryActivityGetReportPDFOK with default headers values
func NewReportSummaryActivityGetReportPDFOK() *ReportSummaryActivityGetReportPDFOK {
	return &ReportSummaryActivityGetReportPDFOK{}
}

/*ReportSummaryActivityGetReportPDFOK handles this case with default header values.

OK
*/
type ReportSummaryActivityGetReportPDFOK struct {
	Payload interface{}
}

func (o *ReportSummaryActivityGetReportPDFOK) Error() string {
	return fmt.Sprintf("[GET /api/reports/summary/activity/pdf][%d] reportSummaryActivityGetReportPDFOK  %+v", 200, o.Payload)
}

func (o *ReportSummaryActivityGetReportPDFOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ReportSummaryActivityGetReportPDFOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}