// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/libraries/go/tmetric/models"
)

// ClientsGetClientReader is a Reader for the ClientsGetClient structure.
type ClientsGetClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClientsGetClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClientsGetClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewClientsGetClientOK creates a ClientsGetClientOK with default headers values
func NewClientsGetClientOK() *ClientsGetClientOK {
	return &ClientsGetClientOK{}
}

/*ClientsGetClientOK handles this case with default header values.

OK
*/
type ClientsGetClientOK struct {
	Payload *models.Client
}

func (o *ClientsGetClientOK) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{accountId}/clients/{clientId}][%d] clientsGetClientOK  %+v", 200, o.Payload)
}

func (o *ClientsGetClientOK) GetPayload() *models.Client {
	return o.Payload
}

func (o *ClientsGetClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Client)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}