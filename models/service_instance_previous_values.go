// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceInstancePreviousValues service instance previous values
//
// swagger:model ServiceInstancePreviousValues
type ServiceInstancePreviousValues struct {

	// maintenance info
	MaintenanceInfo *MaintenanceInfo `json:"maintenance_info,omitempty"`

	// organization id
	OrganizationID string `json:"organization_id,omitempty"`

	// plan id
	PlanID string `json:"plan_id,omitempty"`

	// service id
	ServiceID string `json:"service_id,omitempty"`

	// space id
	SpaceID string `json:"space_id,omitempty"`
}

// Validate validates this service instance previous values
func (m *ServiceInstancePreviousValues) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaintenanceInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceInstancePreviousValues) validateMaintenanceInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.MaintenanceInfo) { // not required
		return nil
	}

	if m.MaintenanceInfo != nil {
		if err := m.MaintenanceInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance_info")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service instance previous values based on the context it is used
func (m *ServiceInstancePreviousValues) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaintenanceInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceInstancePreviousValues) contextValidateMaintenanceInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.MaintenanceInfo != nil {
		if err := m.MaintenanceInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceInstancePreviousValues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceInstancePreviousValues) UnmarshalBinary(b []byte) error {
	var res ServiceInstancePreviousValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
