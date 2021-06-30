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

// ServiceBindingVolumeMountDevice service binding volume mount device
//
// swagger:model ServiceBindingVolumeMountDevice
type ServiceBindingVolumeMountDevice struct {

	// mount config
	MountConfig interface{} `json:"mount_config,omitempty"`

	// volume id
	// Required: true
	VolumeID *string `json:"volume_id"`
}

// Validate validates this service binding volume mount device
func (m *ServiceBindingVolumeMountDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceBindingVolumeMountDevice) validateVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("volume_id", "body", m.VolumeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service binding volume mount device based on context it is used
func (m *ServiceBindingVolumeMountDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceBindingVolumeMountDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceBindingVolumeMountDevice) UnmarshalBinary(b []byte) error {
	var res ServiceBindingVolumeMountDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
