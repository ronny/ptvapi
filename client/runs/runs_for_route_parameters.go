// Code generated by go-swagger; DO NOT EDIT.

package runs

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

// NewRunsForRouteParams creates a new RunsForRouteParams object
// with the default values initialized.
func NewRunsForRouteParams() *RunsForRouteParams {
	var ()
	return &RunsForRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunsForRouteParamsWithTimeout creates a new RunsForRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunsForRouteParamsWithTimeout(timeout time.Duration) *RunsForRouteParams {
	var ()
	return &RunsForRouteParams{

		timeout: timeout,
	}
}

// NewRunsForRouteParamsWithContext creates a new RunsForRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunsForRouteParamsWithContext(ctx context.Context) *RunsForRouteParams {
	var ()
	return &RunsForRouteParams{

		Context: ctx,
	}
}

// NewRunsForRouteParamsWithHTTPClient creates a new RunsForRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunsForRouteParamsWithHTTPClient(client *http.Client) *RunsForRouteParams {
	var ()
	return &RunsForRouteParams{
		HTTPClient: client,
	}
}

/*RunsForRouteParams contains all the parameters to send to the API endpoint
for the runs for route operation typically these are written to a http.Request
*/
type RunsForRouteParams struct {

	/*Devid
	  Your developer id

	*/
	Devid *string
	/*RouteID
	  Identifier of route; values returned by Routes API - v3/routes.

	*/
	RouteID int32
	/*Signature
	  Authentication signature for request

	*/
	Signature *string
	/*Token
	  Please ignore

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the runs for route params
func (o *RunsForRouteParams) WithTimeout(timeout time.Duration) *RunsForRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the runs for route params
func (o *RunsForRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the runs for route params
func (o *RunsForRouteParams) WithContext(ctx context.Context) *RunsForRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the runs for route params
func (o *RunsForRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the runs for route params
func (o *RunsForRouteParams) WithHTTPClient(client *http.Client) *RunsForRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the runs for route params
func (o *RunsForRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevid adds the devid to the runs for route params
func (o *RunsForRouteParams) WithDevid(devid *string) *RunsForRouteParams {
	o.SetDevid(devid)
	return o
}

// SetDevid adds the devid to the runs for route params
func (o *RunsForRouteParams) SetDevid(devid *string) {
	o.Devid = devid
}

// WithRouteID adds the routeID to the runs for route params
func (o *RunsForRouteParams) WithRouteID(routeID int32) *RunsForRouteParams {
	o.SetRouteID(routeID)
	return o
}

// SetRouteID adds the routeId to the runs for route params
func (o *RunsForRouteParams) SetRouteID(routeID int32) {
	o.RouteID = routeID
}

// WithSignature adds the signature to the runs for route params
func (o *RunsForRouteParams) WithSignature(signature *string) *RunsForRouteParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the runs for route params
func (o *RunsForRouteParams) SetSignature(signature *string) {
	o.Signature = signature
}

// WithToken adds the token to the runs for route params
func (o *RunsForRouteParams) WithToken(token *string) *RunsForRouteParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the runs for route params
func (o *RunsForRouteParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *RunsForRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param route_id
	if err := r.SetPathParam("route_id", swag.FormatInt32(o.RouteID)); err != nil {
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
