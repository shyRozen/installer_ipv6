// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServiceBindingResource service binding resource
// swagger:model ServiceBindingResource
type ServiceBindingResource struct {

	// credentials
	Credentials Object `json:"credentials,omitempty"`

	// parameters
	Parameters Object `json:"parameters,omitempty"`

	// route service url
	RouteServiceURL string `json:"route_service_url,omitempty"`

	// syslog drain url
	SyslogDrainURL string `json:"syslog_drain_url,omitempty"`

	// volume mounts
	VolumeMounts []*ServiceBindingVolumeMount `json:"volume_mounts"`
}

// Validate validates this service binding resource
func (m *ServiceBindingResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVolumeMounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceBindingResource) validateVolumeMounts(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeMounts) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeMounts); i++ {
		if swag.IsZero(m.VolumeMounts[i]) { // not required
			continue
		}

		if m.VolumeMounts[i] != nil {
			if err := m.VolumeMounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volume_mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceBindingResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceBindingResource) UnmarshalBinary(b []byte) error {
	var res ServiceBindingResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
