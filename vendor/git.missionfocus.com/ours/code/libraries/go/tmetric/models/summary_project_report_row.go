// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SummaryProjectReportRow summary project report row
//
// swagger:model SummaryProjectReportRow
type SummaryProjectReportRow struct {

	// billable amount
	BillableAmount []*Money `json:"billableAmount"`

	// billable duration
	BillableDuration float64 `json:"billableDuration,omitempty"`

	// budget currency
	BudgetCurrency string `json:"budgetCurrency,omitempty"`

	// client
	Client string `json:"client,omitempty"`

	// client Id
	ClientID int32 `json:"clientId,omitempty"`

	// costs
	Costs []*Money `json:"costs"`

	// duration
	Duration float64 `json:"duration,omitempty"`

	// project
	Project string `json:"project,omitempty"`

	// project code
	ProjectCode string `json:"projectCode,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// spent budget
	SpentBudget float64 `json:"spentBudget,omitempty"`

	// total budget
	TotalBudget float64 `json:"totalBudget,omitempty"`
}

// Validate validates this summary project report row
func (m *SummaryProjectReportRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillableAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCosts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SummaryProjectReportRow) validateBillableAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.BillableAmount) { // not required
		return nil
	}

	for i := 0; i < len(m.BillableAmount); i++ {
		if swag.IsZero(m.BillableAmount[i]) { // not required
			continue
		}

		if m.BillableAmount[i] != nil {
			if err := m.BillableAmount[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("billableAmount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SummaryProjectReportRow) validateCosts(formats strfmt.Registry) error {

	if swag.IsZero(m.Costs) { // not required
		return nil
	}

	for i := 0; i < len(m.Costs); i++ {
		if swag.IsZero(m.Costs[i]) { // not required
			continue
		}

		if m.Costs[i] != nil {
			if err := m.Costs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("costs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SummaryProjectReportRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryProjectReportRow) UnmarshalBinary(b []byte) error {
	var res SummaryProjectReportRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}