// Code generated by go-swagger; DO NOT EDIT.

package patterns

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

// NewPatternsGetPatternByRunParams creates a new PatternsGetPatternByRunParams object
// with the default values initialized.
func NewPatternsGetPatternByRunParams() *PatternsGetPatternByRunParams {
	var ()
	return &PatternsGetPatternByRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatternsGetPatternByRunParamsWithTimeout creates a new PatternsGetPatternByRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatternsGetPatternByRunParamsWithTimeout(timeout time.Duration) *PatternsGetPatternByRunParams {
	var ()
	return &PatternsGetPatternByRunParams{

		timeout: timeout,
	}
}

// NewPatternsGetPatternByRunParamsWithContext creates a new PatternsGetPatternByRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatternsGetPatternByRunParamsWithContext(ctx context.Context) *PatternsGetPatternByRunParams {
	var ()
	return &PatternsGetPatternByRunParams{

		Context: ctx,
	}
}

// NewPatternsGetPatternByRunParamsWithHTTPClient creates a new PatternsGetPatternByRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatternsGetPatternByRunParamsWithHTTPClient(client *http.Client) *PatternsGetPatternByRunParams {
	var ()
	return &PatternsGetPatternByRunParams{
		HTTPClient: client,
	}
}

/*PatternsGetPatternByRunParams contains all the parameters to send to the API endpoint
for the patterns get pattern by run operation typically these are written to a http.Request
*/
type PatternsGetPatternByRunParams struct {

	/*DateUtc
	  Filter by the date and time of the request (ISO 8601 UTC format)

	*/
	DateUtc *strfmt.DateTime
	/*Devid
	  Your developer id

	*/
	Devid *string
	/*Expand
	  Objects to be returned in full (i.e. expanded) - options include: all, stop, route, run, direction, disruption. By default disruptions are expanded.

	*/
	Expand []string
	/*RouteType
	  Number identifying transport mode; values returned via RouteTypes API

	*/
	RouteType int32
	/*RunID
	  Identifier of a trip/service run; values returned by Runs API - /v3/route/{route_id} and Departures API

	*/
	RunID int32
	/*Signature
	  Authentication signature for request

	*/
	Signature *string
	/*StopID
	  Filter by stop_id; values returned by Stops API

	*/
	StopID *int32
	/*Token
	  Please ignore

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) WithTimeout(timeout time.Duration) *PatternsGetPatternByRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) WithContext(ctx context.Context) *PatternsGetPatternByRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) WithHTTPClient(client *http.Client) *PatternsGetPatternByRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateUtc adds the dateUtc to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) WithDateUtc(dateUtc *strfmt.DateTime) *PatternsGetPatternByRunParams {
	o.SetDateUtc(dateUtc)
	return o
}

// SetDateUtc adds the dateUtc to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) SetDateUtc(dateUtc *strfmt.DateTime) {
	o.DateUtc = dateUtc
}

// WithDevid adds the devid to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) WithDevid(devid *string) *PatternsGetPatternByRunParams {
	o.SetDevid(devid)
	return o
}

// SetDevid adds the devid to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) SetDevid(devid *string) {
	o.Devid = devid
}

// WithExpand adds the expand to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) WithExpand(expand []string) *PatternsGetPatternByRunParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithRouteType adds the routeType to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) WithRouteType(routeType int32) *PatternsGetPatternByRunParams {
	o.SetRouteType(routeType)
	return o
}

// SetRouteType adds the routeType to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) SetRouteType(routeType int32) {
	o.RouteType = routeType
}

// WithRunID adds the runID to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) WithRunID(runID int32) *PatternsGetPatternByRunParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) SetRunID(runID int32) {
	o.RunID = runID
}

// WithSignature adds the signature to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) WithSignature(signature *string) *PatternsGetPatternByRunParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) SetSignature(signature *string) {
	o.Signature = signature
}

// WithStopID adds the stopID to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) WithStopID(stopID *int32) *PatternsGetPatternByRunParams {
	o.SetStopID(stopID)
	return o
}

// SetStopID adds the stopId to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) SetStopID(stopID *int32) {
	o.StopID = stopID
}

// WithToken adds the token to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) WithToken(token *string) *PatternsGetPatternByRunParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the patterns get pattern by run params
func (o *PatternsGetPatternByRunParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PatternsGetPatternByRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	// path param route_type
	if err := r.SetPathParam("route_type", swag.FormatInt32(o.RouteType)); err != nil {
		return err
	}

	// path param run_id
	if err := r.SetPathParam("run_id", swag.FormatInt32(o.RunID)); err != nil {
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

	if o.StopID != nil {

		// query param stop_id
		var qrStopID int32
		if o.StopID != nil {
			qrStopID = *o.StopID
		}
		qStopID := swag.FormatInt32(qrStopID)
		if qStopID != "" {
			if err := r.SetQueryParam("stop_id", qStopID); err != nil {
				return err
			}
		}

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
