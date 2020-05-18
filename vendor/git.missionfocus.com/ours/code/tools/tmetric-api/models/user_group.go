// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserGroup user group
//
// swagger:model UserGroup
type UserGroup struct {

	// account Id
	AccountID int32 `json:"accountId,omitempty"`

	// members
	Members []*UserGroupMember `json:"members"`

	// name
	// Required: true
	Name *string `json:"name"`

	// supervisors
	Supervisors []*UserGroupSupervisor `json:"supervisors"`

	// user group Id
	UserGroupID int32 `json:"userGroupId,omitempty"`
}

// Validate validates this user group
func (m *UserGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupervisors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserGroup) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UserGroup) validateSupervisors(formats strfmt.Registry) error {

	if swag.IsZero(m.Supervisors) { // not required
		return nil
	}

	for i := 0; i < len(m.Supervisors); i++ {
		if swag.IsZero(m.Supervisors[i]) { // not required
			continue
		}

		if m.Supervisors[i] != nil {
			if err := m.Supervisors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supervisors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGroup) UnmarshalBinary(b []byte) error {
	var res UserGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
