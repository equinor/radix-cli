// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TLSAutomation TLSAutomation describes the current condition of TLS automation
//
// swagger:model TLSAutomation
type TLSAutomation struct {

	// Message is a human readable description of the reason for the status
	Message string `json:"message,omitempty"`

	// Status of certificate automation request
	// Pending TLSAutomationPending  Certificate automation request pending
	// Success TLSAutomationSuccess  Certificate automation request succeeded
	// Failed TLSAutomationFailed  Certificate automation request failed
	// Example: Pending
	// Required: true
	// Enum: [Pending Success Failed]
	Status *string `json:"status"`
}

// Validate validates this TLS automation
func (m *TLSAutomation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tlsAutomationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Success","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tlsAutomationTypeStatusPropEnum = append(tlsAutomationTypeStatusPropEnum, v)
	}
}

const (

	// TLSAutomationStatusPending captures enum value "Pending"
	TLSAutomationStatusPending string = "Pending"

	// TLSAutomationStatusSuccess captures enum value "Success"
	TLSAutomationStatusSuccess string = "Success"

	// TLSAutomationStatusFailed captures enum value "Failed"
	TLSAutomationStatusFailed string = "Failed"
)

// prop value enum
func (m *TLSAutomation) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tlsAutomationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TLSAutomation) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this TLS automation based on context it is used
func (m *TLSAutomation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TLSAutomation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TLSAutomation) UnmarshalBinary(b []byte) error {
	var res TLSAutomation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
