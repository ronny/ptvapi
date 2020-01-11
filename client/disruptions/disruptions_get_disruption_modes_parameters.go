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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDisruptionsGetDisruptionModesParams creates a new DisruptionsGetDisruptionModesParams object
// with the default values initialized.
func NewDisruptionsGetDisruptionModesParams() *DisruptionsGetDisruptionModesParams {
	var ()
	return &DisruptionsGetDisruptionModesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDisruptionsGetDisruptionModesParamsWithTimeout creates a new DisruptionsGetDisruptionModesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDisruptionsGetDisruptionModesParamsWithTimeout(timeout time.Duration) *DisruptionsGetDisruptionModesParams {
	var ()
	return &DisruptionsGetDisruptionModesParams{

		timeout: timeout,
	}
}

// NewDisruptionsGetDisruptionModesParamsWithContext creates a new DisruptionsGetDisruptionModesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDisruptionsGetDisruptionModesParamsWithContext(ctx context.Context) *DisruptionsGetDisruptionModesParams {
	var ()
	return &DisruptionsGetDisruptionModesParams{

		Context: ctx,
	}
}

// NewDisruptionsGetDisruptionModesParamsWithHTTPClient creates a new DisruptionsGetDisruptionModesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDisruptionsGetDisruptionModesParamsWithHTTPClient(client *http.Client) *DisruptionsGetDisruptionModesParams {
	var ()
	return &DisruptionsGetDisruptionModesParams{
		HTTPClient: client,
	}
}

/*DisruptionsGetDisruptionModesParams contains all the parameters to send to the API endpoint
for the disruptions get disruption modes operation typically these are written to a http.Request
*/
type DisruptionsGetDisruptionModesParams struct {

	/*Devid
	  Your developer id

	*/
	Devid *string
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

// WithTimeout adds the timeout to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) WithTimeout(timeout time.Duration) *DisruptionsGetDisruptionModesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) WithContext(ctx context.Context) *DisruptionsGetDisruptionModesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) WithHTTPClient(client *http.Client) *DisruptionsGetDisruptionModesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevid adds the devid to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) WithDevid(devid *string) *DisruptionsGetDisruptionModesParams {
	o.SetDevid(devid)
	return o
}

// SetDevid adds the devid to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) SetDevid(devid *string) {
	o.Devid = devid
}

// WithSignature adds the signature to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) WithSignature(signature *string) *DisruptionsGetDisruptionModesParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) SetSignature(signature *string) {
	o.Signature = signature
}

// WithToken adds the token to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) WithToken(token *string) *DisruptionsGetDisruptionModesParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the disruptions get disruption modes params
func (o *DisruptionsGetDisruptionModesParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *DisruptionsGetDisruptionModesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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