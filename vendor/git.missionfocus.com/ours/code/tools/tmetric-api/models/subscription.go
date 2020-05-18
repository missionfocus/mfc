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

// Subscription Represents a customer's subscription for a TMetric account.
//
// swagger:model Subscription
type Subscription struct {

	// Gets or sets the identifier of an account to which the subscription is attached.
	AccountID int32 `json:"accountId,omitempty"`

	// Gets or sets the subscription's expiration date.
	// Format: date-time
	ExpirationDate strfmt.DateTime `json:"expirationDate,omitempty"`

	// Gets a value indicating whether this subscription is annual.
	IsAnnual bool `json:"isAnnual,omitempty"`

	// Gets or sets a value indicating whether the subscription has recurring billing flag.
	IsRecurring bool `json:"isRecurring,omitempty"`

	// Gets or sets a value indicating whether the subscription is created for evaluation purpose.
	IsTrial bool `json:"isTrial,omitempty"`

	// Gets or sets the maximum allowed number of users for the account
	// under this subscription. The value is defined at the time of purchase
	// by the 'Number of Users' pricing option in Avangate's system.
	MaxUsers int32 `json:"maxUsers,omitempty"`

	// Gets or sets the subscription's start date.
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// Gets or sets the subscription status.
	// Enum: [0 1 2 3 4 128]
	Status int32 `json:"status,omitempty"`

	// Gets or sets the subscription plan identifier.
	SubscriptionID int32 `json:"subscriptionId,omitempty"`

	// Gets or sets the subscription plan that defines the product features
	// avaiable for the subscritpion. This is a navigation property for SubscriptionPlanId.
	SubscriptionPlan *SubscriptionPlan `json:"subscriptionPlan,omitempty"`

	// Gets or sets the identifier of the subscription plan
	// that defines the product features avaiable for the subscritpion.
	SubscriptionPlanID int32 `json:"subscriptionPlanId,omitempty"`
}

// Validate validates this subscription
func (m *Subscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionPlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscription) validateExpirationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpirationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expirationDate", "body", "date-time", m.ExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var subscriptionTypeStatusPropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,128]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionTypeStatusPropEnum = append(subscriptionTypeStatusPropEnum, v)
	}
}

// prop value enum
func (m *Subscription) validateStatusEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, subscriptionTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateSubscriptionPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionPlan) { // not required
		return nil
	}

	if m.SubscriptionPlan != nil {
		if err := m.SubscriptionPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscriptionPlan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscription) UnmarshalBinary(b []byte) error {
	var res Subscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
