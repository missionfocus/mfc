// Code generated by go-swagger; DO NOT EDIT.

package time_off_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/libraries/go/tmetric/models"
)

// TimeOffPoliciesGetTimeOffPolicyReader is a Reader for the TimeOffPoliciesGetTimeOffPolicy structure.
type TimeOffPoliciesGetTimeOffPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimeOffPoliciesGetTimeOffPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimeOffPoliciesGetTimeOffPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimeOffPoliciesGetTimeOffPolicyOK creates a TimeOffPoliciesGetTimeOffPolicyOK with default headers values
func NewTimeOffPoliciesGetTimeOffPolicyOK() *TimeOffPoliciesGetTimeOffPolicyOK {
	return &TimeOffPoliciesGetTimeOffPolicyOK{}
}

/*TimeOffPoliciesGetTimeOffPolicyOK handles this case with default header values.

OK
*/
type TimeOffPoliciesGetTimeOffPolicyOK struct {
	Payload *models.TimeOffPolicy
}

func (o *TimeOffPoliciesGetTimeOffPolicyOK) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{accountId}/timeoff/policies/{timeOffPolicyId}][%d] timeOffPoliciesGetTimeOffPolicyOK  %+v", 200, o.Payload)
}

func (o *TimeOffPoliciesGetTimeOffPolicyOK) GetPayload() *models.TimeOffPolicy {
	return o.Payload
}

func (o *TimeOffPoliciesGetTimeOffPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeOffPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
