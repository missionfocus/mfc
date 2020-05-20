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

// PersonalRate personal rate
//
// swagger:model PersonalRate
type PersonalRate struct {

	// billable rate amount
	// Maximum: 1e+09
	// Minimum: 0
	BillableRateAmount *float64 `json:"billableRateAmount,omitempty"`

	// budget hours
	// Maximum: 1e+09
	// Minimum: 0
	BudgetHours *float64 `json:"budgetHours,omitempty"`

	// member fee
	// Maximum: 1e+09
	// Minimum: 0
	MemberFee *float64 `json:"memberFee,omitempty"`

	// user profile Id
	UserProfileID int32 `json:"userProfileId,omitempty"`
}

// Validate validates this personal rate
func (m *PersonalRate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillableRateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBudgetHours(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberFee(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PersonalRate) validateBillableRateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.BillableRateAmount) { // not required
		return nil
	}

	if err := validate.Minimum("billableRateAmount", "body", float64(*m.BillableRateAmount), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("billableRateAmount", "body", float64(*m.BillableRateAmount), 1e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *PersonalRate) validateBudgetHours(formats strfmt.Registry) error {

	if swag.IsZero(m.BudgetHours) { // not required
		return nil
	}

	if err := validate.Minimum("budgetHours", "body", float64(*m.BudgetHours), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("budgetHours", "body", float64(*m.BudgetHours), 1e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *PersonalRate) validateMemberFee(formats strfmt.Registry) error {

	if swag.IsZero(m.MemberFee) { // not required
		return nil
	}

	if err := validate.Minimum("memberFee", "body", float64(*m.MemberFee), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("memberFee", "body", float64(*m.MemberFee), 1e+09, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PersonalRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PersonalRate) UnmarshalBinary(b []byte) error {
	var res PersonalRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
