// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/libraries/go/tmetric/models"
)

// IntegrationsPostIntegrationReader is a Reader for the IntegrationsPostIntegration structure.
type IntegrationsPostIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationsPostIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIntegrationsPostIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIntegrationsPostIntegrationOK creates a IntegrationsPostIntegrationOK with default headers values
func NewIntegrationsPostIntegrationOK() *IntegrationsPostIntegrationOK {
	return &IntegrationsPostIntegrationOK{}
}

/*IntegrationsPostIntegrationOK handles this case with default header values.

OK
*/
type IntegrationsPostIntegrationOK struct {
	Payload *models.Integration
}

func (o *IntegrationsPostIntegrationOK) Error() string {
	return fmt.Sprintf("[POST /api/accounts/{accountId}/integrations][%d] integrationsPostIntegrationOK  %+v", 200, o.Payload)
}

func (o *IntegrationsPostIntegrationOK) GetPayload() *models.Integration {
	return o.Payload
}

func (o *IntegrationsPostIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Integration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
