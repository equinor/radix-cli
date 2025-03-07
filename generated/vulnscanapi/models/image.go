// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Image Image holds name of image the base image
//
// swagger:model Image
type Image struct {

	// Base image
	// Example: alpine:3.13.2
	BaseImage string `json:"baseImage,omitempty"`

	// Name of image
	// Example: quay.io/oauth2-proxy/oauth2-proxy:v7.1.3
	// Required: true
	ImageName *string `json:"image"`
}

// Validate validates this image
func (m *Image) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Image) validateImageName(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.ImageName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this image based on context it is used
func (m *Image) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Image) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Image) UnmarshalBinary(b []byte) error {
	var res Image
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
