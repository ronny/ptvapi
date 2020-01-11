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

// V3BulkDeparturesResponse v3 bulk departures response
// swagger:model V3.BulkDeparturesResponse
type V3BulkDeparturesResponse struct {

	// Directions of travel of route
	Directions []*V3Direction `json:"directions"`

	// Disruption information applicable to relevant routes or stops
	Disruptions map[string]V3Disruption `json:"disruptions,omitempty"`

	// Contains departures for the requested stop and route(s). It includes details as to the route_direction and whether it is still valid.
	Responses []*V3BulkDeparturesUpdateResponse `json:"responses"`

	// Train lines, tram routes, bus routes, regional coach routes, Night Bus routes
	Routes []*V3Route `json:"routes"`

	// Individual trips/services of a route
	Runs []*V3Run `json:"runs"`

	// API Status / Metadata
	Status *V3Status `json:"status,omitempty"`

	// A train station, tram stop, bus stop, regional coach stop or Night Bus stop
	Stops map[string]V3BulkDeparturesStopResponse `json:"stops,omitempty"`
}

// Validate validates this v3 bulk departures response
func (m *V3BulkDeparturesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisruptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponses(formats); err != nil {
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

func (m *V3BulkDeparturesResponse) validateDirections(formats strfmt.Registry) error {

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

func (m *V3BulkDeparturesResponse) validateDisruptions(formats strfmt.Registry) error {

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

func (m *V3BulkDeparturesResponse) validateResponses(formats strfmt.Registry) error {

	if swag.IsZero(m.Responses) { // not required
		return nil
	}

	for i := 0; i < len(m.Responses); i++ {
		if swag.IsZero(m.Responses[i]) { // not required
			continue
		}

		if m.Responses[i] != nil {
			if err := m.Responses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("responses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3BulkDeparturesResponse) validateRoutes(formats strfmt.Registry) error {

	if swag.IsZero(m.Routes) { // not required
		return nil
	}

	for i := 0; i < len(m.Routes); i++ {
		if swag.IsZero(m.Routes[i]) { // not required
			continue
		}

		if m.Routes[i] != nil {
			if err := m.Routes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3BulkDeparturesResponse) validateRuns(formats strfmt.Registry) error {

	if swag.IsZero(m.Runs) { // not required
		return nil
	}

	for i := 0; i < len(m.Runs); i++ {
		if swag.IsZero(m.Runs[i]) { // not required
			continue
		}

		if m.Runs[i] != nil {
			if err := m.Runs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3BulkDeparturesResponse) validateStatus(formats strfmt.Registry) error {

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

func (m *V3BulkDeparturesResponse) validateStops(formats strfmt.Registry) error {

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
func (m *V3BulkDeparturesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3BulkDeparturesResponse) UnmarshalBinary(b []byte) error {
	var res V3BulkDeparturesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
