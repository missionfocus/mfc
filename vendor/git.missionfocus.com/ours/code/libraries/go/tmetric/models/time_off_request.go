// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TimeOffRequest time off request
//
// swagger:model TimeOffRequest
type TimeOffRequest struct {

	// account Id
	AccountID int32 `json:"accountId,omitempty"`

	// approver
	// Read Only: true
	Approver *UserProfileLite `json:"approver,omitempty"`

	// approver Id
	ApproverID int32 `json:"approverId,omitempty"`

	// denial reason
	DenialReason string `json:"denialReason,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// end date
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// hours
	Hours float64 `json:"hours,omitempty"`

	// request date
	// Format: date-time
	RequestDate strfmt.DateTime `json:"requestDate,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// status
	// Enum: [0 1 2 3]
	Status int32 `json:"status,omitempty"`

	// time off policy Id
	TimeOffPolicyID int32 `json:"timeOffPolicyId,omitempty"`

	// time off request Id
	TimeOffRequestID int32 `json:"timeOffRequestId,omitempty"`

	// user profile
	// Read Only: true
	UserProfile *UserProfileLite `json:"userProfile,omitempty"`

	// user profile Id
	UserProfileID int32 `json:"userProfileId,omitempty"`

	// workdays
	// Read Only: true
	Workdays int32 `json:"workdays,omitempty"`
}

// Validate validates this time off request
func (m *TimeOffRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApprover(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeOffRequest) validateApprover(formats strfmt.Registry) error {

	if swag.IsZero(m.Approver) { // not required
		return nil
	}

	if m.Approver != nil {
		if err := m.Approver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("approver")
			}
			return err
		}
	}

	return nil
}

func (m *TimeOffRequest) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TimeOffRequest) validateRequestDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestDate", "body", "date-time", m.RequestDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TimeOffRequest) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var timeOffRequestTypeStatusPropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		timeOffRequestTypeStatusPropEnum = append(timeOffRequestTypeStatusPropEnum, v)
	}
}

// prop value enum
func (m *TimeOffRequest) validateStatusEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, timeOffRequestTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TimeOffRequest) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *TimeOffRequest) validateUserProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.UserProfile) { // not required
		return nil
	}

	if m.UserProfile != nil {
		if err := m.UserProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userProfile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimeOffRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeOffRequest) UnmarshalBinary(b []byte) error {
	var res TimeOffRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}