// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Client Client related to specific account.
//
// swagger:model Client
type Client struct {

	// Gets or sets the account identifier.
	AccountID int32 `json:"accountId,omitempty"`

	// Gets or sets the active projects count.
	ActiveProjectsCount int32 `json:"activeProjectsCount,omitempty"`

	// Gets or sets the relative path to the client avatar.
	// For example, /Content/Avatars/avatar_12.svg.
	Avatar string `json:"avatar,omitempty"`

	// Gets or sets the address of the client.
	ClientAddress string `json:"clientAddress,omitempty"`

	// Gets or sets the client identifier.
	ClientID int32 `json:"clientId,omitempty"`

	// Gets or sets the name of the client.
	ClientName string `json:"clientName,omitempty"`

	// contact users
	ContactUsers []int32 `json:"contactUsers"`

	// default billable rate
	DefaultBillableRate *Rate `json:"defaultBillableRate,omitempty"`

	// Gets or sets the total projects count.
	TotalProjectsCount int32 `json:"totalProjectsCount,omitempty"`
}

// Validate validates this client
func (m *Client) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultBillableRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Client) validateDefaultBillableRate(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *Client) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Client) UnmarshalBinary(b []byte) error {
	var res Client
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
