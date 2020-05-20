// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IntegrationsDeleteSynchronizationReader is a Reader for the IntegrationsDeleteSynchronization structure.
type IntegrationsDeleteSynchronizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationsDeleteSynchronizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIntegrationsDeleteSynchronizationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIntegrationsDeleteSynchronizationNoContent creates a IntegrationsDeleteSynchronizationNoContent with default headers values
func NewIntegrationsDeleteSynchronizationNoContent() *IntegrationsDeleteSynchronizationNoContent {
	return &IntegrationsDeleteSynchronizationNoContent{}
}

/*IntegrationsDeleteSynchronizationNoContent handles this case with default header values.

No Content
*/
type IntegrationsDeleteSynchronizationNoContent struct {
}

func (o *IntegrationsDeleteSynchronizationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/accounts/{accountId}/integrations/sync/{integrationId}][%d] integrationsDeleteSynchronizationNoContent ", 204)
}

func (o *IntegrationsDeleteSynchronizationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
