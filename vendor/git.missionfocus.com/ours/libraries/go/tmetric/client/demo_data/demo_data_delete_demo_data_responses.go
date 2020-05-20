// Code generated by go-swagger; DO NOT EDIT.

package demo_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DemoDataDeleteDemoDataReader is a Reader for the DemoDataDeleteDemoData structure.
type DemoDataDeleteDemoDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DemoDataDeleteDemoDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDemoDataDeleteDemoDataNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDemoDataDeleteDemoDataNoContent creates a DemoDataDeleteDemoDataNoContent with default headers values
func NewDemoDataDeleteDemoDataNoContent() *DemoDataDeleteDemoDataNoContent {
	return &DemoDataDeleteDemoDataNoContent{}
}

/*DemoDataDeleteDemoDataNoContent handles this case with default header values.

No Content
*/
type DemoDataDeleteDemoDataNoContent struct {
}

func (o *DemoDataDeleteDemoDataNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/demodata/{accountId}][%d] demoDataDeleteDemoDataNoContent ", 204)
}

func (o *DemoDataDeleteDemoDataNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
