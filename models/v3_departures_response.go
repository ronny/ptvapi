// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V3DeparturesResponse v3 departures response
// swagger:model V3.DeparturesResponse
type V3DeparturesResponse struct {

	// Timetabled and real-time service departures
	Departures []*V3Departure `json:"departures"`

	// Directions of travel of route
	Directions map[string]V3Direction `json:"directions,omitempty"`

	// Disruption information applicable to relevant routes or stops
	Disruptions map[string]V3Disruption `json:"disruptions,omitempty"`

	// Train lines, tram routes, bus routes, regional coach routes, Night Bus routes
	Routes map[string]V3Route `json:"routes,omitempty"`

	// Individual trips/services of a route
	Runs map[string]V3Run `json:"runs,omitempty"`

	// API Status / Metadata
	Status *V3Status `json:"status,omitempty"`

	// A train station, tram stop, bus stop, regional coach stop or Night Bus stop
	Stops map[string]V3ResultStop `json:"stops,omitempty"`
}

// Validate validates this v3 departures response
func (m *V3DeparturesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDepartures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisruptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStops(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3DeparturesResponse) validateDepartures(formats strfmt.Registry) error {

	if swag.IsZero(m.Departures) { // not required
		return nil
	}

	for i := 0; i < len(m.Departures); i++ {
		if swag.IsZero(m.Departures[i]) { // not required
			continue
		}

		if m.Departures[i] != nil {
			if err := m.Departures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("departures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3DeparturesResponse) validateDirections(formats strfmt.Registry) error {

	if swag.IsZero(m.Directions) { // not required
		return nil
	}

	for k := range m.Directions {

		if err := validate.Required("directions"+"."+k, "body", m.Directions[k]); err != nil {
			return err
		}
		if val, ok := m.Directions[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V3DeparturesResponse) validateDisruptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Disruptions) { // not required
		return nil
	}

	for k := range m.Disruptions {

		if err := validate.Required("disruptions"+"."+k, "body", m.Disruptions[k]); err != nil {
			return err
		}
		if val, ok := m.Disruptions[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V3DeparturesResponse) validateRoutes(formats strfmt.Registry) error {

	if swag.IsZero(m.Routes) { // not required
		return nil
	}

	for k := range m.Routes {

		if err := validate.Required("routes"+"."+k, "body", m.Routes[k]); err != nil {
			return err
		}
		if val, ok := m.Routes[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V3DeparturesResponse) validateRuns(formats strfmt.Registry) error {

	if swag.IsZero(m.Runs) { // not required
		return nil
	}

	for k := range m.Runs {

		if err := validate.Required("runs"+"."+k, "body", m.Runs[k]); err != nil {
			return err
		}
		if val, ok := m.Runs[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V3DeparturesResponse) validateStatus(formats strfmt.Registry) error {

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

func (m *V3DeparturesResponse) validateStops(formats strfmt.Registry) error {

	if swag.IsZero(m.Stops) { // not required
		return nil
	}

	for k := range m.Stops {

		if err := validate.Required("stops"+"."+k, "body", m.Stops[k]); err != nil {
			return err
		}
		if val, ok := m.Stops[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3DeparturesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3DeparturesResponse) UnmarshalBinary(b []byte) error {
	var res V3DeparturesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
