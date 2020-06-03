// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Account account
//
// swagger:model Account
type Account struct {

	// account Id
	AccountID int32 `json:"accountId,omitempty"`

	// account logo
	AccountLogo string `json:"accountLogo,omitempty"`

	// account name
	AccountName string `json:"accountName,omitempty"`

	// activity capture settings
	ActivityCaptureSettings *ActivityCaptureSettingsMap `json:"activityCaptureSettings,omitempty"`

	// blur screenshots
	BlurScreenshots bool `json:"blurScreenshots,omitempty"`

	// can members create tags
	CanMembersCreateTags bool `json:"canMembersCreateTags,omitempty"`

	// can members manage public projects
	CanMembersManagePublicProjects bool `json:"canMembersManagePublicProjects,omitempty"`

	// company address
	CompanyAddress string `json:"companyAddress,omitempty"`

	// dates
	Dates []*AccountDate `json:"dates"`

	// default billable rate
	DefaultBillableRate *Rate `json:"defaultBillableRate,omitempty"`

	// 0 - only today is editable, -1 - all days are editable
	EditableDays int32 `json:"editableDays,omitempty"`

	// external account Id
	// Format: uuid
	ExternalAccountID strfmt.UUID `json:"externalAccountId,omitempty"`

	// 0 - sunday, 1 - monday, 6 - saturday
	FirstWeekDay int32 `json:"firstWeekDay,omitempty"`

	// has demo data
	HasDemoData bool `json:"hasDemoData,omitempty"`

	// inactivity stop minutes
	InactivityStopMinutes int32 `json:"inactivityStopMinutes,omitempty"`

	// Returns members time tracking permissions
	Members []*MemberInfo `json:"members"`

	// permissions
	Permissions *AccountPermissionsMap `json:"permissions,omitempty"`

	// report detailed time enabled
	ReportDetailedTimeEnabled bool `json:"reportDetailedTimeEnabled,omitempty"`

	// report time format
	ReportTimeFormat string `json:"reportTimeFormat,omitempty"`

	// report time rounding minutes
	ReportTimeRoundingMinutes int32 `json:"reportTimeRoundingMinutes,omitempty"`

	// report time rounding mode
	// Enum: [0 1 2 3]
	ReportTimeRoundingMode int32 `json:"reportTimeRoundingMode,omitempty"`

	// required fields
	RequiredFields *RequiredFieldsMap `json:"requiredFields,omitempty"`

	// workdays
	Workdays []bool `json:"workdays"`

	// working hours
	// Maximum: 24
	// Minimum: 0
	WorkingHours *float64 `json:"workingHours,omitempty"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityCaptureSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultBillableRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportTimeRoundingMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkingHours(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) validateActivityCaptureSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityCaptureSettings) { // not required
		return nil
	}

	if m.ActivityCaptureSettings != nil {
		if err := m.ActivityCaptureSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activityCaptureSettings")
			}
			return err
		}
	}

	return nil
}

func (m *Account) validateDates(formats strfmt.Registry) error {

	if swag.IsZero(m.Dates) { // not required
		return nil
	}

	for i := 0; i < len(m.Dates); i++ {
		if swag.IsZero(m.Dates[i]) { // not required
			continue
		}

		if m.Dates[i] != nil {
			if err := m.Dates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Account) validateDefaultBillableRate(formats strfmt.Registry) error {

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

func (m *Account) validateExternalAccountID(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalAccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("externalAccountId", "body", "uuid", m.ExternalAccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateMembers(formats strfmt.Registry) error {

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

func (m *Account) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

var accountTypeReportTimeRoundingModePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountTypeReportTimeRoundingModePropEnum = append(accountTypeReportTimeRoundingModePropEnum, v)
	}
}

// prop value enum
func (m *Account) validateReportTimeRoundingModeEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, accountTypeReportTimeRoundingModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Account) validateReportTimeRoundingMode(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportTimeRoundingMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateReportTimeRoundingModeEnum("reportTimeRoundingMode", "body", m.ReportTimeRoundingMode); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateRequiredFields(formats strfmt.Registry) error {

	if swag.IsZero(m.RequiredFields) { // not required
		return nil
	}

	if m.RequiredFields != nil {
		if err := m.RequiredFields.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requiredFields")
			}
			return err
		}
	}

	return nil
}

func (m *Account) validateWorkingHours(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkingHours) { // not required
		return nil
	}

	if err := validate.Minimum("workingHours", "body", float64(*m.WorkingHours), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("workingHours", "body", float64(*m.WorkingHours), 24, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}