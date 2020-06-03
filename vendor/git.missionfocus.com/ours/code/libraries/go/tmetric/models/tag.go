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

// Tag tag
//
// swagger:model Tag
type Tag struct {

	// account Id
	AccountID int32 `json:"accountId,omitempty"`

	// default billable rate
	DefaultBillableRate *Rate `json:"defaultBillableRate,omitempty"`

	// is work type
	IsWorkType bool `json:"isWorkType,omitempty"`

	// tag Id
	TagID int32 `json:"tagId,omitempty"`

	// tag name
	// Required: true
	TagName *string `json:"tagName"`

	// work type projects
	WorkTypeProjects []int32 `json:"workTypeProjects"`
}

// Validate validates this tag
func (m *Tag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultBillableRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tag) validateDefaultBillableRate(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultBillableRate) { // not required
		return nil
	}

	if m.DefaultBillableRate != nil {
		if err := m.DefaultBillableRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultBillableRate")
			}
			return err
		}
	}

	return nil
}

func (m *Tag) validateTagName(formats strfmt.Registry) error {

	if err := validate.Required("tagName", "body", m.TagName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tag) UnmarshalBinary(b []byte) error {
	var res Tag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}