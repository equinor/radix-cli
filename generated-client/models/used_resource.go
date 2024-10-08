// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UsedResource UsedResource holds information about used resource
//
// swagger:model UsedResource
type UsedResource struct {

	// Avg Average resource used
	// Example: 0.00023
	Avg float64 `json:"avg,omitempty"`

	// Max resource used
	// Example: 0.00037
	Max float64 `json:"max,omitempty"`

	// Min resource used
	// Example: 0.00012
	Min float64 `json:"min,omitempty"`
}

// Validate validates this used resource
func (m *UsedResource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this used resource based on context it is used
func (m *UsedResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UsedResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsedResource) UnmarshalBinary(b []byte) error {
	var res UsedResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
