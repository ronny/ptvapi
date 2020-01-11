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

// V3OutletResponse v3 outlet response
// swagger:model V3.OutletResponse
type V3OutletResponse struct {

	// myki ticket outlets
	Outlets []*V3Outlet `json:"outlets"`

	// API Status / Metadata
	Status *V3Status `json:"status,omitempty"`
}

// Validate validates this v3 outlet response
func (m *V3OutletResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutlets(formats); err != nil {
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

func (m *V3OutletResponse) validateOutlets(formats strfmt.Registry) error {

	if swag.IsZero(m.Outlets) { // not required
		return nil
	}

	for i := 0; i < len(m.Outlets); i++ {
		if swag.IsZero(m.Outlets[i]) { // not required
			continue
		}

		if m.Outlets[i] != nil {
			if err := m.Outlets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outlets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3OutletResponse) validateStatus(formats strfmt.Registry) error {

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
func (m *V3OutletResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3OutletResponse) UnmarshalBinary(b []byte) error {
	var res V3OutletResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
