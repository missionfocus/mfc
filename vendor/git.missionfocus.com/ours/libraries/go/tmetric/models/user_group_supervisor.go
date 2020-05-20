// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserGroupSupervisor user group supervisor
//
// swagger:model UserGroupSupervisor
type UserGroupSupervisor struct {

	// account Id
	AccountID int32 `json:"accountId,omitempty"`

	// user group Id
	UserGroupID int32 `json:"userGroupId,omitempty"`

	// user profile
	UserProfile *UserProfile `json:"userProfile,omitempty"`

	// user profile Id
	UserProfileID int32 `json:"userProfileId,omitempty"`
}

// Validate validates this user group supervisor
func (m *UserGroupSupervisor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserGroupSupervisor) validateUserProfile(formats strfmt.Registry) error {

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
func (m *UserGroupSupervisor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGroupSupervisor) UnmarshalBinary(b []byte) error {
	var res UserGroupSupervisor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
