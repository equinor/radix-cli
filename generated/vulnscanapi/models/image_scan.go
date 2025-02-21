// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageScan ImageScan holdes information about a spcific vulnerability scan for an image
//
// swagger:model ImageScan
type ImageScan struct {

	// Flag indicating if scan succeeded or not
	// Example: true
	// Required: true
	ScanSuccess *bool `json:"scanSuccess"`

	// Date and time of scan
	// Example: 2022-05-05T14:26:45+02:00
	// Required: true
	// Format: date-time
	ScanTime *strfmt.DateTime `json:"scanTime"`

	// List of vulnerabilities
	Vulnerabilities []*Vulnerability `json:"vulnerabilities"`

	// Count of vulnerabilities grouped by severity
	VulnerabilitySummary map[string]int64 `json:"vulnerabilitySummary,omitempty"`
}

// Validate validates this image scan
func (m *ImageScan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScanSuccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageScan) validateScanSuccess(formats strfmt.Registry) error {

	if err := validate.Required("scanSuccess", "body", m.ScanSuccess); err != nil {
		return err
	}

	return nil
}

func (m *ImageScan) validateScanTime(formats strfmt.Registry) error {

	if err := validate.Required("scanTime", "body", m.ScanTime); err != nil {
		return err
	}

	if err := validate.FormatOf("scanTime", "body", "date-time", m.ScanTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageScan) validateVulnerabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Vulnerabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.Vulnerabilities); i++ {
		if swag.IsZero(m.Vulnerabilities[i]) { // not required
			continue
		}

		if m.Vulnerabilities[i] != nil {
			if err := m.Vulnerabilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vulnerabilities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vulnerabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this image scan based on the context it is used
func (m *ImageScan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVulnerabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageScan) contextValidateVulnerabilities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vulnerabilities); i++ {

		if m.Vulnerabilities[i] != nil {

			if swag.IsZero(m.Vulnerabilities[i]) { // not required
				return nil
			}

			if err := m.Vulnerabilities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vulnerabilities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vulnerabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageScan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageScan) UnmarshalBinary(b []byte) error {
	var res ImageScan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
