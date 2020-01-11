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

// V3DirectionsResponse v3 directions response
// swagger:model V3.DirectionsResponse
type V3DirectionsResponse struct {

	// Directions of travel of route
	Directions []*V3DirectionWithDescription `json:"directions"`

	// API Status / Metadata
	Status *V3Status `json:"status,omitempty"`
}

// Validate validates this v3 directions response
func (m *V3DirectionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3DirectionsResponse) validateDirections(formats strfmt.Registry) error {

	if swag.IsZero(m.Directions) { // not required
		return nil
	}

	for i := 0; i < len(m.Directions); i++ {
		if swag.IsZero(m.Directions[i]) { // not required
			continue
		}

		if m.Directions[i] != nil {
			if err := m.Directions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("directions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3DirectionsResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3DirectionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3DirectionsResponse) UnmarshalBinary(b []byte) error {
	var res V3DirectionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}