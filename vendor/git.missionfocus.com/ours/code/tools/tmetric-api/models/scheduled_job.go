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

// ScheduledJob scheduled job
//
// swagger:model ScheduledJob
type ScheduledJob struct {

	// account Id
	AccountID int32 `json:"accountId,omitempty"`

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// integration Id
	IntegrationID int32 `json:"integrationId,omitempty"`

	// job status
	// Enum: [0 1 2]
	JobStatus int32 `json:"jobStatus,omitempty"`

	// job type
	// Enum: [0 1 2 3 4 5 6 7]
	JobType int32 `json:"jobType,omitempty"`

	// last finish time
	// Format: date-time
	LastFinishTime strfmt.DateTime `json:"lastFinishTime,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// next run time
	// Format: date-time
	NextRunTime strfmt.DateTime `json:"nextRunTime,omitempty"`

	// scheduled job Id
	ScheduledJobID int32 `json:"scheduledJobId,omitempty"`
}

// Validate validates this scheduled job
func (m *ScheduledJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastFinishTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextRunTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var scheduledJobTypeJobStatusPropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduledJobTypeJobStatusPropEnum = append(scheduledJobTypeJobStatusPropEnum, v)
	}
}

// prop value enum
func (m *ScheduledJob) validateJobStatusEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, scheduledJobTypeJobStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ScheduledJob) validateJobStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.JobStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateJobStatusEnum("jobStatus", "body", m.JobStatus); err != nil {
		return err
	}

	return nil
}

var scheduledJobTypeJobTypePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5,6,7]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduledJobTypeJobTypePropEnum = append(scheduledJobTypeJobTypePropEnum, v)
	}
}

// prop value enum
func (m *ScheduledJob) validateJobTypeEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, scheduledJobTypeJobTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ScheduledJob) validateJobType(formats strfmt.Registry) error {

	if swag.IsZero(m.JobType) { // not required
		return nil
	}

	// value enum
	if err := m.validateJobTypeEnum("jobType", "body", m.JobType); err != nil {
		return err
	}

	return nil
}

func (m *ScheduledJob) validateLastFinishTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastFinishTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastFinishTime", "body", "date-time", m.LastFinishTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ScheduledJob) validateNextRunTime(formats strfmt.Registry) error {

	if swag.IsZero(m.NextRunTime) { // not required
		return nil
	}

	if err := validate.FormatOf("nextRunTime", "body", "date-time", m.NextRunTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduledJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduledJob) UnmarshalBinary(b []byte) error {
	var res ScheduledJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
