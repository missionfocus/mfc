// Code generated by go-swagger; DO NOT EDIT.

package time_off_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/tools/tmetric-api/models"
)

// TimeOffPoliciesPostTimeOffPolicyReader is a Reader for the TimeOffPoliciesPostTimeOffPolicy structure.
type TimeOffPoliciesPostTimeOffPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimeOffPoliciesPostTimeOffPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTimeOffPoliciesPostTimeOffPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimeOffPoliciesPostTimeOffPolicyOK creates a TimeOffPoliciesPostTimeOffPolicyOK with default headers values
func NewTimeOffPoliciesPostTimeOffPolicyOK() *TimeOffPoliciesPostTimeOffPolicyOK {
	return &TimeOffPoliciesPostTimeOffPolicyOK{}
}

/*TimeOffPoliciesPostTimeOffPolicyOK handles this case with default header values.

OK
*/
type TimeOffPoliciesPostTimeOffPolicyOK struct {
	Payload *models.TimeOffPolicy
}

func (o *TimeOffPoliciesPostTimeOffPolicyOK) Error() string {
	return fmt.Sprintf("[POST /api/accounts/{accountId}/timeoff/policies][%d] timeOffPoliciesPostTimeOffPolicyOK  %+v", 200, o.Payload)
}

func (o *TimeOffPoliciesPostTimeOffPolicyOK) GetPayload() *models.TimeOffPolicy {
	return o.Payload
}

func (o *TimeOffPoliciesPostTimeOffPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeOffPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
