// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimeZoneInfoLite time zone info lite
//
// swagger:model TimeZoneInfoLite
type TimeZoneInfoLite struct {

	// current offset
	CurrentOffset float64 `json:"currentOffset,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// summer offset
	SummerOffset float64 `json:"summerOffset,omitempty"`

	// winter offset
	WinterOffset float64 `json:"winterOffset,omitempty"`
}

// Validate validates this time zone info lite
func (m *TimeZoneInfoLite) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeZoneInfoLite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeZoneInfoLite) UnmarshalBinary(b []byte) error {
	var res TimeZoneInfoLite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
