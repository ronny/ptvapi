// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V3DisruptionsResponse v3 disruptions response
// swagger:model V3.DisruptionsResponse
type V3DisruptionsResponse struct {

	// Disruption information applicable to relevant routes or stops
	Disruptions *V3Disruptions `json:"disruptions,omitempty"`

	// API Status / Metadata
	Status *V3Status `json:"status,omitempty"`
}

// Validate validates this v3 disruptions response
func (m *V3DisruptionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisruptions(formats); err != nil {
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

func (m *V3DisruptionsResponse) validateDisruptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Disruptions) { // not required
		return nil
	}

	if m.Disruptions != nil {
		if err := m.Disruptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disruptions")
			}
			return err
		}
	}

	return nil
}

func (m *V3DisruptionsResponse) validateStatus(formats strfmt.Registry) error {

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
func (m *V3DisruptionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3DisruptionsResponse) UnmarshalBinary(b []byte) error {
	var res V3DisruptionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
