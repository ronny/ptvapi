// Code generated by go-swagger; DO NOT EDIT.

package departures

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeparturesGetForStopParams creates a new DeparturesGetForStopParams object
// with the default values initialized.
func NewDeparturesGetForStopParams() *DeparturesGetForStopParams {
	var ()
	return &DeparturesGetForStopParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeparturesGetForStopParamsWithTimeout creates a new DeparturesGetForStopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeparturesGetForStopParamsWithTimeout(timeout time.Duration) *DeparturesGetForStopParams {
	var ()
	return &DeparturesGetForStopParams{

		timeout: timeout,
	}
}

// NewDeparturesGetForStopParamsWithContext creates a new DeparturesGetForStopParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeparturesGetForStopParamsWithContext(ctx context.Context) *DeparturesGetForStopParams {
	var ()
	return &DeparturesGetForStopParams{

		Context: ctx,
	}
}

// NewDeparturesGetForStopParamsWithHTTPClient creates a new DeparturesGetForStopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeparturesGetForStopParamsWithHTTPClient(client *http.Client) *DeparturesGetForStopParams {
	var ()
	return &DeparturesGetForStopParams{
		HTTPClient: client,
	}
}

/*DeparturesGetForStopParams contains all the parameters to send to the API endpoint
for the departures get for stop operation typically these are written to a http.Request
*/
type DeparturesGetForStopParams struct {

	/*DateUtc
	  Filter by the date and time of the request (ISO 8601 UTC format) (default = current date and time)

	*/
	DateUtc *strfmt.DateTime
	/*Devid
	  Your developer id

	*/
	Devid *string
	/*DirectionID
	  Filter by identifier of direction of travel; values returned by Directions API - /v3/directions/route/{route_id}

	*/
	DirectionID *int32
	/*Expand
	  List objects to be returned in full (i.e. expanded) - options include: all, stop, route, run, direction, disruption

	*/
	Expand []string
	/*Gtfs
	  Indicates that stop_id parameter will accept "GTFS stop_id" data

	*/
	Gtfs *bool
	/*IncludeCancelled
	  Indicates if cancelled services (if they exist) are returned (default = false) - metropolitan train only

	*/
	IncludeCancelled *bool
	/*LookBackwards
	  Indicates if filtering runs (and their departures) to those that arrive at destination before date_utc (default = false). Requires max_results &gt; 0.

	*/
	LookBackwards *bool
	/*MaxResults
	  Maximum number of results returned

	*/
	MaxResults *int32
	/*PlatformNumbers
	  Filter by platform number at stop

	*/
	PlatformNumbers []int32
	/*RouteType
	  Number identifying transport mode; values returned via RouteTypes API

	*/
	RouteType int32
	/*Signature
	  Authentication signature for request

	*/
	Signature *string
	/*StopID
	  Identifier of stop; values returned by Stops API

	*/
	StopID int32
	/*Token
	  Please ignore

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the departures get for stop params
func (o *DeparturesGetForStopParams) WithTimeout(timeout time.Duration) *DeparturesGetForStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the departures get for stop params
func (o *DeparturesGetForStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the departures get for stop params
func (o *DeparturesGetForStopParams) WithContext(ctx context.Context) *DeparturesGetForStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the departures get for stop params
func (o *DeparturesGetForStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the departures get for stop params
func (o *DeparturesGetForStopParams) WithHTTPClient(client *http.Client) *DeparturesGetForStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the departures get for stop params
func (o *DeparturesGetForStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateUtc adds the dateUtc to the departures get for stop params
func (o *DeparturesGetForStopParams) WithDateUtc(dateUtc *strfmt.DateTime) *DeparturesGetForStopParams {
	o.SetDateUtc(dateUtc)
	return o
}

// SetDateUtc adds the dateUtc to the departures get for stop params
func (o *DeparturesGetForStopParams) SetDateUtc(dateUtc *strfmt.DateTime) {
	o.DateUtc = dateUtc
}

// WithDevid adds the devid to the departures get for stop params
func (o *DeparturesGetForStopParams) WithDevid(devid *string) *DeparturesGetForStopParams {
	o.SetDevid(devid)
	return o
}

// SetDevid adds the devid to the departures get for stop params
func (o *DeparturesGetForStopParams) SetDevid(devid *string) {
	o.Devid = devid
}

// WithDirectionID adds the directionID to the departures get for stop params
func (o *DeparturesGetForStopParams) WithDirectionID(directionID *int32) *DeparturesGetForStopParams {
	o.SetDirectionID(directionID)
	return o
}

// SetDirectionID adds the directionId to the departures get for stop params
func (o *DeparturesGetForStopParams) SetDirectionID(directionID *int32) {
	o.DirectionID = directionID
}

// WithExpand adds the expand to the departures get for stop params
func (o *DeparturesGetForStopParams) WithExpand(expand []string) *DeparturesGetForStopParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the departures get for stop params
func (o *DeparturesGetForStopParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithGtfs adds the gtfs to the departures get for stop params
func (o *DeparturesGetForStopParams) WithGtfs(gtfs *bool) *DeparturesGetForStopParams {
	o.SetGtfs(gtfs)
	return o
}

// SetGtfs adds the gtfs to the departures get for stop params
func (o *DeparturesGetForStopParams) SetGtfs(gtfs *bool) {
	o.Gtfs = gtfs
}

// WithIncludeCancelled adds the includeCancelled to the departures get for stop params
func (o *DeparturesGetForStopParams) WithIncludeCancelled(includeCancelled *bool) *DeparturesGetForStopParams {
	o.SetIncludeCancelled(includeCancelled)
	return o
}

// SetIncludeCancelled adds the includeCancelled to the departures get for stop params
func (o *DeparturesGetForStopParams) SetIncludeCancelled(includeCancelled *bool) {
	o.IncludeCancelled = includeCancelled
}

// WithLookBackwards adds the lookBackwards to the departures get for stop params
func (o *DeparturesGetForStopParams) WithLookBackwards(lookBackwards *bool) *DeparturesGetForStopParams {
	o.SetLookBackwards(lookBackwards)
	return o
}

// SetLookBackwards adds the lookBackwards to the departures get for stop params
func (o *DeparturesGetForStopParams) SetLookBackwards(lookBackwards *bool) {
	o.LookBackwards = lookBackwards
}

// WithMaxResults adds the maxResults to the departures get for stop params
func (o *DeparturesGetForStopParams) WithMaxResults(maxResults *int32) *DeparturesGetForStopParams {
	o.SetMaxResults(maxResults)
	return o
}

// SetMaxResults adds the maxResults to the departures get for stop params
func (o *DeparturesGetForStopParams) SetMaxResults(maxResults *int32) {
	o.MaxResults = maxResults
}

// WithPlatformNumbers adds the platformNumbers to the departures get for stop params
func (o *DeparturesGetForStopParams) WithPlatformNumbers(platformNumbers []int32) *DeparturesGetForStopParams {
	o.SetPlatformNumbers(platformNumbers)
	return o
}

// SetPlatformNumbers adds the platformNumbers to the departures get for stop params
func (o *DeparturesGetForStopParams) SetPlatformNumbers(platformNumbers []int32) {
	o.PlatformNumbers = platformNumbers
}

// WithRouteType adds the routeType to the departures get for stop params
func (o *DeparturesGetForStopParams) WithRouteType(routeType int32) *DeparturesGetForStopParams {
	o.SetRouteType(routeType)
	return o
}

// SetRouteType adds the routeType to the departures get for stop params
func (o *DeparturesGetForStopParams) SetRouteType(routeType int32) {
	o.RouteType = routeType
}

// WithSignature adds the signature to the departures get for stop params
func (o *DeparturesGetForStopParams) WithSignature(signature *string) *DeparturesGetForStopParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the departures get for stop params
func (o *DeparturesGetForStopParams) SetSignature(signature *string) {
	o.Signature = signature
}

// WithStopID adds the stopID to the departures get for stop params
func (o *DeparturesGetForStopParams) WithStopID(stopID int32) *DeparturesGetForStopParams {
	o.SetStopID(stopID)
	return o
}

// SetStopID adds the stopId to the departures get for stop params
func (o *DeparturesGetForStopParams) SetStopID(stopID int32) {
	o.StopID = stopID
}

// WithToken adds the token to the departures get for stop params
func (o *DeparturesGetForStopParams) WithToken(token *string) *DeparturesGetForStopParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the departures get for stop params
func (o *DeparturesGetForStopParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *DeparturesGetForStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DateUtc != nil {

		// query param date_utc
		var qrDateUtc strfmt.DateTime
		if o.DateUtc != nil {
			qrDateUtc = *o.DateUtc
		}
		qDateUtc := qrDateUtc.String()
		if qDateUtc != "" {
			if err := r.SetQueryParam("date_utc", qDateUtc); err != nil {
				return err
			}
		}

	}

	if o.Devid != nil {

		// query param devid
		var qrDevid string
		if o.Devid != nil {
			qrDevid = *o.Devid
		}
		qDevid := qrDevid
		if qDevid != "" {
			if err := r.SetQueryParam("devid", qDevid); err != nil {
				return err
			}
		}

	}

	if o.DirectionID != nil {

		// query param direction_id
		var qrDirectionID int32
		if o.DirectionID != nil {
			qrDirectionID = *o.DirectionID
		}
		qDirectionID := swag.FormatInt32(qrDirectionID)
		if qDirectionID != "" {
			if err := r.SetQueryParam("direction_id", qDirectionID); err != nil {
				return err
			}
		}

	}

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	if o.Gtfs != nil {

		// query param gtfs
		var qrGtfs bool
		if o.Gtfs != nil {
			qrGtfs = *o.Gtfs
		}
		qGtfs := swag.FormatBool(qrGtfs)
		if qGtfs != "" {
			if err := r.SetQueryParam("gtfs", qGtfs); err != nil {
				return err
			}
		}

	}

	if o.IncludeCancelled != nil {

		// query param include_cancelled
		var qrIncludeCancelled bool
		if o.IncludeCancelled != nil {
			qrIncludeCancelled = *o.IncludeCancelled
		}
		qIncludeCancelled := swag.FormatBool(qrIncludeCancelled)
		if qIncludeCancelled != "" {
			if err := r.SetQueryParam("include_cancelled", qIncludeCancelled); err != nil {
				return err
			}
		}

	}

	if o.LookBackwards != nil {

		// query param look_backwards
		var qrLookBackwards bool
		if o.LookBackwards != nil {
			qrLookBackwards = *o.LookBackwards
		}
		qLookBackwards := swag.FormatBool(qrLookBackwards)
		if qLookBackwards != "" {
			if err := r.SetQueryParam("look_backwards", qLookBackwards); err != nil {
				return err
			}
		}

	}

	if o.MaxResults != nil {

		// query param max_results
		var qrMaxResults int32
		if o.MaxResults != nil {
			qrMaxResults = *o.MaxResults
		}
		qMaxResults := swag.FormatInt32(qrMaxResults)
		if qMaxResults != "" {
			if err := r.SetQueryParam("max_results", qMaxResults); err != nil {
				return err
			}
		}

	}

	var valuesPlatformNumbers []string
	for _, v := range o.PlatformNumbers {
		valuesPlatformNumbers = append(valuesPlatformNumbers, swag.FormatInt32(v))
	}

	joinedPlatformNumbers := swag.JoinByFormat(valuesPlatformNumbers, "multi")
	// query array param platform_numbers
	if err := r.SetQueryParam("platform_numbers", joinedPlatformNumbers...); err != nil {
		return err
	}

	// path param route_type
	if err := r.SetPathParam("route_type", swag.FormatInt32(o.RouteType)); err != nil {
		return err
	}

	if o.Signature != nil {

		// query param signature
		var qrSignature string
		if o.Signature != nil {
			qrSignature = *o.Signature
		}
		qSignature := qrSignature
		if qSignature != "" {
			if err := r.SetQueryParam("signature", qSignature); err != nil {
				return err
			}
		}

	}

	// path param stop_id
	if err := r.SetPathParam("stop_id", swag.FormatInt32(o.StopID)); err != nil {
		return err
	}

	if o.Token != nil {

		// query param token
		var qrToken string
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
