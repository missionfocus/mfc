// Code generated by go-swagger; DO NOT EDIT.

package time_off_balances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/libraries/go/tmetric/models"
)

// TimeOffBalancesGetTimeOffBalanceMovementsReader is a Reader for the TimeOffBalancesGetTimeOffBalanceMovements structure.
type TimeOffBalancesGetTimeOffBalanceMovementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimeOffBalancesGetTimeOffBalanceMovementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimeOffBalancesGetTimeOffBalanceMovementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimeOffBalancesGetTimeOffBalanceMovementsOK creates a TimeOffBalancesGetTimeOffBalanceMovementsOK with default headers values
func NewTimeOffBalancesGetTimeOffBalanceMovementsOK() *TimeOffBalancesGetTimeOffBalanceMovementsOK {
	return &TimeOffBalancesGetTimeOffBalanceMovementsOK{}
}

/*TimeOffBalancesGetTimeOffBalanceMovementsOK handles this case with default header values.

OK
*/
type TimeOffBalancesGetTimeOffBalanceMovementsOK struct {
	Payload []*models.BalanceMovement
}

func (o *TimeOffBalancesGetTimeOffBalanceMovementsOK) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{accountId}/timeoff/balances/movements][%d] timeOffBalancesGetTimeOffBalanceMovementsOK  %+v", 200, o.Payload)
}

func (o *TimeOffBalancesGetTimeOffBalanceMovementsOK) GetPayload() []*models.BalanceMovement {
	return o.Payload
}

func (o *TimeOffBalancesGetTimeOffBalanceMovementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}