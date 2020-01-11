// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V3ResultStop v3 result stop
// swagger:model V3.ResultStop
type V3ResultStop struct {

	// Transport mode identifier
	RouteType int32 `json:"route_type,omitempty"`

	// Distance of stop from input location (in metres); returns 0 if no location is input
	StopDistance float32 `json:"stop_distance,omitempty"`

	// Stop identifier
	StopID int32 `json:"stop_id,omitempty"`

	// Geographic coordinate of latitude at stop
	StopLatitude float32 `json:"stop_latitude,omitempty"`

	// Geographic coordinate of longitude at stop
	StopLongitude float32 `json:"stop_longitude,omitempty"`

	// Name of stop
	StopName string `json:"stop_name,omitempty"`

	// Sequence of the stop on the route/run; return 0 when route_id or run_id not specified. Order ascendingly by this field (when non zero) to get physical order (earliest first) of stops on the route_id/run_id.
	StopSequence int32 `json:"stop_sequence,omitempty"`

	// suburb of stop
	StopSuburb string `json:"stop_suburb,omitempty"`
}

// Validate validates this v3 result stop
func (m *V3ResultStop) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V3ResultStop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3ResultStop) UnmarshalBinary(b []byte) error {
	var res V3ResultStop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
