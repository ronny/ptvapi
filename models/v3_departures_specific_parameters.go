// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V3DeparturesSpecificParameters v3 departures specific parameters
// swagger:model V3.DeparturesSpecificParameters
type V3DeparturesSpecificParameters struct {

	// Filter by the date and time of the request (ISO 8601 UTC format) (default = current date and time)
	// Format: date-time
	DateUtc strfmt.DateTime `json:"date_utc,omitempty"`

	// Filter by identifier of direction of travel; values returned by Directions API - /v3/directions/route/{route_id}
	DirectionID int32 `json:"direction_id,omitempty"`

	// List objects to be returned in full (i.e. expanded) - options include: all, stop, route, run, direction, disruption
	Expand []string `json:"expand"`

	// Indicates that stop_id parameter will accept "GTFS stop_id" data
	Gtfs bool `json:"gtfs,omitempty"`

	// Indicates if cancelled services (if they exist) are returned (default = false) - metropolitan train only
	IncludeCancelled bool `json:"include_cancelled,omitempty"`

	// Indicates if filtering runs (and their departures) to those that arrive at destination before date_utc (default = false). Requires max_results &gt; 0.
	LookBackwards bool `json:"look_backwards,omitempty"`

	// Maximum number of results returned
	MaxResults int32 `json:"max_results,omitempty"`
}

// Validate validates this v3 departures specific parameters
func (m *V3DeparturesSpecificParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateUtc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpand(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3DeparturesSpecificParameters) validateDateUtc(formats strfmt.Registry) error {

	if swag.IsZero(m.DateUtc) { // not required
		return nil
	}

	if err := validate.FormatOf("date_utc", "body", "date-time", m.DateUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

var v3DeparturesSpecificParametersExpandItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["All","Stop","Route","Run","Direction","Disruption"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3DeparturesSpecificParametersExpandItemsEnum = append(v3DeparturesSpecificParametersExpandItemsEnum, v)
	}
}

func (m *V3DeparturesSpecificParameters) validateExpandItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v3DeparturesSpecificParametersExpandItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *V3DeparturesSpecificParameters) validateExpand(formats strfmt.Registry) error {

	if swag.IsZero(m.Expand) { // not required
		return nil
	}

	for i := 0; i < len(m.Expand); i++ {

		// value enum
		if err := m.validateExpandItemsEnum("expand"+"."+strconv.Itoa(i), "body", m.Expand[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3DeparturesSpecificParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3DeparturesSpecificParameters) UnmarshalBinary(b []byte) error {
	var res V3DeparturesSpecificParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}