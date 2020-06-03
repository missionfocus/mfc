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

// InvoiceRange invoice range
//
// swagger:model InvoiceRange
type InvoiceRange struct {

	// client Id
	ClientID int32 `json:"clientId,omitempty"`

	// end time
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// invoice type
	// Enum: [0 1 2 3]
	InvoiceType int32 `json:"invoiceType,omitempty"`

	// projects
	Projects []int32 `json:"projects"`

	// start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`
}

// Validate validates this invoice range
func (m *InvoiceRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceRange) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var invoiceRangeTypeInvoiceTypePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceRangeTypeInvoiceTypePropEnum = append(invoiceRangeTypeInvoiceTypePropEnum, v)
	}
}

// prop value enum
func (m *InvoiceRange) validateInvoiceTypeEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, invoiceRangeTypeInvoiceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceRange) validateInvoiceType(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoiceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateInvoiceTypeEnum("invoiceType", "body", m.InvoiceType); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceRange) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceRange) UnmarshalBinary(b []byte) error {
	var res InvoiceRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}