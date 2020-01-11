// Code generated by go-swagger; DO NOT EDIT.

package disruptions

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

// NewDisruptionsGetDisruptionsByStopParams creates a new DisruptionsGetDisruptionsByStopParams object
// with the default values initialized.
func NewDisruptionsGetDisruptionsByStopParams() *DisruptionsGetDisruptionsByStopParams {
	var ()
	return &DisruptionsGetDisruptionsByStopParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDisruptionsGetDisruptionsByStopParamsWithTimeout creates a new DisruptionsGetDisruptionsByStopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDisruptionsGetDisruptionsByStopParamsWithTimeout(timeout time.Duration) *DisruptionsGetDisruptionsByStopParams {
	var ()
	return &DisruptionsGetDisruptionsByStopParams{

		timeout: timeout,
	}
}

// NewDisruptionsGetDisruptionsByStopParamsWithContext creates a new DisruptionsGetDisruptionsByStopParams object
// with the default values initialized, and the ability to set a context for a request
func NewDisruptionsGetDisruptionsByStopParamsWithContext(ctx context.Context) *DisruptionsGetDisruptionsByStopParams {
	var ()
	return &DisruptionsGetDisruptionsByStopParams{

		Context: ctx,
	}
}

// NewDisruptionsGetDisruptionsByStopParamsWithHTTPClient creates a new DisruptionsGetDisruptionsByStopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDisruptionsGetDisruptionsByStopParamsWithHTTPClient(client *http.Client) *DisruptionsGetDisruptionsByStopParams {
	var ()
	return &DisruptionsGetDisruptionsByStopParams{
		HTTPClient: client,
	}
}

/*DisruptionsGetDisruptionsByStopParams contains all the parameters to send to the API endpoint
for the disruptions get disruptions by stop operation typically these are written to a http.Request
*/
type DisruptionsGetDisruptionsByStopParams struct {

	/*Devid
	  Your developer id

	*/
	Devid *string
	/*DisruptionStatus
	  Filter by status of disruption

	*/
	DisruptionStatus *string
	/*Signature
	  Authentication signature for request

	*/
	Signature *string
	/*StopID
	  Identifier of stop; values returned by Stops API - v3/stops

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

// WithTimeout adds the timeout to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) WithTimeout(timeout time.Duration) *DisruptionsGetDisruptionsByStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) WithContext(ctx context.Context) *DisruptionsGetDisruptionsByStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) WithHTTPClient(client *http.Client) *DisruptionsGetDisruptionsByStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevid adds the devid to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) WithDevid(devid *string) *DisruptionsGetDisruptionsByStopParams {
	o.SetDevid(devid)
	return o
}

// SetDevid adds the devid to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) SetDevid(devid *string) {
	o.Devid = devid
}

// WithDisruptionStatus adds the disruptionStatus to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) WithDisruptionStatus(disruptionStatus *string) *DisruptionsGetDisruptionsByStopParams {
	o.SetDisruptionStatus(disruptionStatus)
	return o
}

// SetDisruptionStatus adds the disruptionStatus to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) SetDisruptionStatus(disruptionStatus *string) {
	o.DisruptionStatus = disruptionStatus
}

// WithSignature adds the signature to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) WithSignature(signature *string) *DisruptionsGetDisruptionsByStopParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) SetSignature(signature *string) {
	o.Signature = signature
}

// WithStopID adds the stopID to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) WithStopID(stopID int32) *DisruptionsGetDisruptionsByStopParams {
	o.SetStopID(stopID)
	return o
}

// SetStopID adds the stopId to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) SetStopID(stopID int32) {
	o.StopID = stopID
}

// WithToken adds the token to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) WithToken(token *string) *DisruptionsGetDisruptionsByStopParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the disruptions get disruptions by stop params
func (o *DisruptionsGetDisruptionsByStopParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *DisruptionsGetDisruptionsByStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.DisruptionStatus != nil {

		// query param disruption_status
		var qrDisruptionStatus string
		if o.DisruptionStatus != nil {
			qrDisruptionStatus = *o.DisruptionStatus
		}
		qDisruptionStatus := qrDisruptionStatus
		if qDisruptionStatus != "" {
			if err := r.SetQueryParam("disruption_status", qDisruptionStatus); err != nil {
				return err
			}
		}

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