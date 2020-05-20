// Code generated by go-swagger; DO NOT EDIT.

package time_off_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TimeOffPoliciesDeleteTimeOffPolicyReader is a Reader for the TimeOffPoliciesDeleteTimeOffPolicy structure.
type TimeOffPoliciesDeleteTimeOffPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimeOffPoliciesDeleteTimeOffPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTimeOffPoliciesDeleteTimeOffPolicyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimeOffPoliciesDeleteTimeOffPolicyNoContent creates a TimeOffPoliciesDeleteTimeOffPolicyNoContent with default headers values
func NewTimeOffPoliciesDeleteTimeOffPolicyNoContent() *TimeOffPoliciesDeleteTimeOffPolicyNoContent {
	return &TimeOffPoliciesDeleteTimeOffPolicyNoContent{}
}

/*TimeOffPoliciesDeleteTimeOffPolicyNoContent handles this case with default header values.

No Content
*/
type TimeOffPoliciesDeleteTimeOffPolicyNoContent struct {
}

func (o *TimeOffPoliciesDeleteTimeOffPolicyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/accounts/{accountId}/timeoff/policies/{timeOffPolicyId}][%d] timeOffPoliciesDeleteTimeOffPolicyNoContent ", 204)
}

func (o *TimeOffPoliciesDeleteTimeOffPolicyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
