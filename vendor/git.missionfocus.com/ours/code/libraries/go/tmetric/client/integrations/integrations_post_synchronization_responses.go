// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IntegrationsPostSynchronizationReader is a Reader for the IntegrationsPostSynchronization structure.
type IntegrationsPostSynchronizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationsPostSynchronizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIntegrationsPostSynchronizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIntegrationsPostSynchronizationOK creates a IntegrationsPostSynchronizationOK with default headers values
func NewIntegrationsPostSynchronizationOK() *IntegrationsPostSynchronizationOK {
	return &IntegrationsPostSynchronizationOK{}
}

/*IntegrationsPostSynchronizationOK handles this case with default header values.

OK
*/
type IntegrationsPostSynchronizationOK struct {
	Payload interface{}
}

func (o *IntegrationsPostSynchronizationOK) Error() string {
	return fmt.Sprintf("[POST /api/accounts/{accountId}/integrations/sync/{integrationId}][%d] integrationsPostSynchronizationOK  %+v", 200, o.Payload)
}

func (o *IntegrationsPostSynchronizationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *IntegrationsPostSynchronizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}