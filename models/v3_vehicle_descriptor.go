// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V3VehicleDescriptor v3 vehicle descriptor
// swagger:model V3.VehicleDescriptor
type V3VehicleDescriptor struct {

	// Indicator if vehicle is air conditioned. May be null. Only available for some tram runs.
	AirConditioned bool `json:"air_conditioned,omitempty"`

	// Vehicle description such as "6 Car Comeng", "6 Car Xtrapolis", "3 Car Comeng", "6 Car Siemens", "3 Car Siemens". May be null/empty.
	//             Only available for some metropolitan train runs.
	Description string `json:"description,omitempty"`

	// Operator identifier of the vehicle such as "26094". May be null/empty. Only available for some tram and bus runs.
	ID string `json:"id,omitempty"`

	// Indicator if vehicle has a low floor. May be null. Only available for some tram runs.
	LowFloor bool `json:"low_floor,omitempty"`

	// Operator name of the vehicle such as "Metro Trains Melbourne", "Yarra Trams", "Ventura Bus Line", "CDC" or "Sita Bus Lines" . May be null/empty.
	//             Only available for train, tram, v/line and some bus runs.
	Operator string `json:"operator,omitempty"`

	// Supplier of vehicle descriptor data.
	// Read Only: true
	Supplier string `json:"supplier,omitempty"`
}

// Validate validates this v3 vehicle descriptor
func (m *V3VehicleDescriptor) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V3VehicleDescriptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3VehicleDescriptor) UnmarshalBinary(b []byte) error {
	var res V3VehicleDescriptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}