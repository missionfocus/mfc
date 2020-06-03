// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SyncTimeEntry sync time entry
//
// swagger:model SyncTimeEntry
type SyncTimeEntry struct {

	// client name
	ClientName string `json:"clientName,omitempty"`

	// date
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// details
	Details *TimeEntryDetail `json:"details,omitempty"`

	// is billable
	IsBillable bool `json:"isBillable,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// rate
	Rate *Rate `json:"rate,omitempty"`

	// seconds
	Seconds int32 `json:"seconds,omitempty"`

	// work type
	WorkType string `json:"workType,omitempty"`
}

// Validate validates this sync time entry
func (m *SyncTimeEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncTimeEntry) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SyncTimeEntry) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *SyncTimeEntry) validateRate(formats strfmt.Registry) error {

	if swag.IsZero(m.Rate) { // not required
		return nil
	}

	if m.Rate != nil {
		if err := m.Rate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SyncTimeEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncTimeEntry) UnmarshalBinary(b []byte) error {
	var res SyncTimeEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}