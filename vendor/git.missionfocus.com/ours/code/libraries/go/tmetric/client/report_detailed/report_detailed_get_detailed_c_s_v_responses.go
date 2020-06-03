// Code generated by go-swagger; DO NOT EDIT.

package report_detailed

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReportDetailedGetDetailedCSVReader is a Reader for the ReportDetailedGetDetailedCSV structure.
type ReportDetailedGetDetailedCSVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportDetailedGetDetailedCSVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportDetailedGetDetailedCSVOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReportDetailedGetDetailedCSVOK creates a ReportDetailedGetDetailedCSVOK with default headers values
func NewReportDetailedGetDetailedCSVOK() *ReportDetailedGetDetailedCSVOK {
	return &ReportDetailedGetDetailedCSVOK{}
}

/*ReportDetailedGetDetailedCSVOK handles this case with default header values.

OK
*/
type ReportDetailedGetDetailedCSVOK struct {
	Payload interface{}
}

func (o *ReportDetailedGetDetailedCSVOK) Error() string {
	return fmt.Sprintf("[GET /api/reports/detailed/csv][%d] reportDetailedGetDetailedCSVOK  %+v", 200, o.Payload)
}

func (o *ReportDetailedGetDetailedCSVOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ReportDetailedGetDetailedCSVOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}