// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V3StopAccessibility v3 stop accessibility
// swagger:model V3.StopAccessibility
type V3StopAccessibility struct {

	// Indicates if there is at least one audio customer information at the stop/platform
	AudioCustomerInformation bool `json:"audio_customer_information,omitempty"`

	// Indicates if there is at least one accessible escalator at the stop/platform that complies with the Disability Standards for Accessible Public Transport under the Disability Discrimination Act (1992)
	Escalator bool `json:"escalator,omitempty"`

	// Indicates if there is a hearing loop facility at the stop/platform
	HearingLoop bool `json:"hearing_loop,omitempty"`

	// Indicates if there is an elevator at the stop/platform
	Lift bool `json:"lift,omitempty"`

	// Indicates if there is lighting at the stop
	Lighting bool `json:"lighting,omitempty"`

	// Indicates the platform number for xivic information (Platform 0 indicates general stop facilities)
	PlatformNumber int32 `json:"platform_number,omitempty"`

	// Indicates if there are stairs available in the stop
	Stairs bool `json:"stairs,omitempty"`

	// Indicates if the stop is accessible
	StopAccessible bool `json:"stop_accessible,omitempty"`

	// Indicates if there are tactile tiles (also known as tactile ground surface indicators, or TGSIs) at the stop
	TactileGroundSurfaceIndicator bool `json:"tactile_ground_surface_indicator,omitempty"`

	// Indicates if there is a general waiting area at the stop
	WaitingRoom bool `json:"waiting_room,omitempty"`

	// Facilities relating to the accessibility of the stop by wheelchair
	Wheelchair *V3StopAccessibilityWheelchair `json:"wheelchair,omitempty"`
}

// Validate validates this v3 stop accessibility
func (m *V3StopAccessibility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWheelchair(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3StopAccessibility) validateWheelchair(formats strfmt.Registry) error {

	if swag.IsZero(m.Wheelchair) { // not required
		return nil
	}

	if m.Wheelchair != nil {
		if err := m.Wheelchair.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wheelchair")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3StopAccessibility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3StopAccessibility) UnmarshalBinary(b []byte) error {
	var res V3StopAccessibility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
