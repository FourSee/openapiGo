// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PairingRequestInput pairing request input
// swagger:model PairingRequestInput
type PairingRequestInput struct {

	// device info
	DeviceInfo *PairingRequestInputDeviceInfo `json:"device_info,omitempty"`

	// device name
	DeviceName string `json:"device_name,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`
}

// Validate validates this pairing request input
func (m *PairingRequestInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PairingRequestInput) validateDeviceInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceInfo) { // not required
		return nil
	}

	if m.DeviceInfo != nil {
		if err := m.DeviceInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PairingRequestInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PairingRequestInput) UnmarshalBinary(b []byte) error {
	var res PairingRequestInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}