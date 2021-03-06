// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V3Departure v3 departure
// swagger:model V3.Departure
type V3Departure struct {

	// Indicates if the metropolitan train service is at the platform at the time of query; returns false for other modes
	AtPlatform bool `json:"at_platform,omitempty"`

	// Chronological sequence of the departure for the run on the route. Order ascendingly by this field to get chronological order (earliest first) of departures with the same route_id and run_id.
	DepartureSequence int32 `json:"departure_sequence,omitempty"`

	// Direction of travel identifier
	DirectionID int32 `json:"direction_id,omitempty"`

	// Disruption information identifier(s)
	DisruptionIds []int64 `json:"disruption_ids"`

	// Real-time estimate of departure time and date in ISO 8601 UTC format
	// Format: date-time
	EstimatedDepartureUtc strfmt.DateTime `json:"estimated_departure_utc,omitempty"`

	// Flag indicating special condition for run (e.g. RR Reservations Required, GC Guaranteed Connection, DOO Drop Off Only, PUO Pick Up Only, MO Mondays only, TU Tuesdays only, WE Wednesdays only, TH Thursdays only, FR Fridays only, SS School days only; ignore E flag)
	Flags string `json:"flags,omitempty"`

	// Platform number at stop (metropolitan train only; returns null for other modes)
	PlatformNumber string `json:"platform_number,omitempty"`

	// Route identifier
	RouteID int32 `json:"route_id,omitempty"`

	// Trip/service run identifier
	RunID int32 `json:"run_id,omitempty"`

	// Scheduled (i.e. timetabled) departure time and date in ISO 8601 UTC format
	// Format: date-time
	ScheduledDepartureUtc strfmt.DateTime `json:"scheduled_departure_utc,omitempty"`

	// Stop identifier
	StopID int32 `json:"stop_id,omitempty"`
}

// Validate validates this v3 departure
func (m *V3Departure) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEstimatedDepartureUtc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledDepartureUtc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3Departure) validateEstimatedDepartureUtc(formats strfmt.Registry) error {

	if swag.IsZero(m.EstimatedDepartureUtc) { // not required
		return nil
	}

	if err := validate.FormatOf("estimated_departure_utc", "body", "date-time", m.EstimatedDepartureUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3Departure) validateScheduledDepartureUtc(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledDepartureUtc) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduled_departure_utc", "body", "date-time", m.ScheduledDepartureUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3Departure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3Departure) UnmarshalBinary(b []byte) error {
	var res V3Departure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
