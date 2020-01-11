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

// V3StopDepartureRequest v3 stop departure request
// swagger:model V3.StopDepartureRequest
type V3StopDepartureRequest struct {

	// Indicates that stop_id parameter will accept "GTFS stop_id" data and route_directions[x].route_id parameters will accept route_gtfs_id data
	Gtfs bool `json:"gtfs,omitempty"`

	// Maximum number of results returned
	// Maximum: 2.147483647e+09
	// Minimum: 0
	MaxResults *int32 `json:"max_results,omitempty"`

	// The route directions to find departures for at this stop.
	// Required: true
	RouteDirections []*V3StopDepartureRequestRouteDirection `json:"route_directions"`

	// Number identifying transport mode; values returned via RouteTypes API
	// Enum: [0 1 2 3 4]
	RouteType int32 `json:"route_type,omitempty"`

	// Identifier of stop; values returned by Stops API
	// Maximum: 2.147483647e+09
	// Minimum: 0
	StopID *int32 `json:"stop_id,omitempty"`
}

// Validate validates this v3 stop departure request
func (m *V3StopDepartureRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteDirections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3StopDepartureRequest) validateMaxResults(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxResults) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_results", "body", int64(*m.MaxResults), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_results", "body", int64(*m.MaxResults), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *V3StopDepartureRequest) validateRouteDirections(formats strfmt.Registry) error {

	if err := validate.Required("route_directions", "body", m.RouteDirections); err != nil {
		return err
	}

	for i := 0; i < len(m.RouteDirections); i++ {
		if swag.IsZero(m.RouteDirections[i]) { // not required
			continue
		}

		if m.RouteDirections[i] != nil {
			if err := m.RouteDirections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("route_directions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var v3StopDepartureRequestTypeRouteTypePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2,3,4]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3StopDepartureRequestTypeRouteTypePropEnum = append(v3StopDepartureRequestTypeRouteTypePropEnum, v)
	}
}

// prop value enum
func (m *V3StopDepartureRequest) validateRouteTypeEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, v3StopDepartureRequestTypeRouteTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V3StopDepartureRequest) validateRouteType(formats strfmt.Registry) error {

	if swag.IsZero(m.RouteType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRouteTypeEnum("route_type", "body", m.RouteType); err != nil {
		return err
	}

	return nil
}

func (m *V3StopDepartureRequest) validateStopID(formats strfmt.Registry) error {

	if swag.IsZero(m.StopID) { // not required
		return nil
	}

	if err := validate.MinimumInt("stop_id", "body", int64(*m.StopID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("stop_id", "body", int64(*m.StopID), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3StopDepartureRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3StopDepartureRequest) UnmarshalBinary(b []byte) error {
	var res V3StopDepartureRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
