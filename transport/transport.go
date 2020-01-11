package transport

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SigningTransport wraps another transport (`Transport`) to automatically add
// the `devid` and `signature` query params to every outgoing request as
// required by the PTV Timetable API v3.
//
// The `signature` query param is calculated from the request URI and `APIKey`.
type SigningTransport struct {
	Transport runtime.ClientTransport
	AuthInfo  runtime.ClientAuthInfoWriter
}

var _ runtime.ClientTransport = (*SigningTransport)(nil)

func NewSigningTransport(transport runtime.ClientTransport, devID string, apiKey string) *SigningTransport {
	authInfo := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		err := r.SetQueryParam("devid", devID)
		if err != nil {
			return err
		}

		requestURI := r.GetPath() + "?" + r.GetQueryParams().Encode()

		sig, err := Signature(apiKey, requestURI)
		if err != nil {
			return err
		}
		return r.SetQueryParam("signature", sig)
	})

	return &SigningTransport{
		Transport: transport,
		AuthInfo:  authInfo,
	}
}

func (t *SigningTransport) Submit(op *runtime.ClientOperation) (interface{}, error) {
	op.AuthInfo = t.AuthInfo
	return t.Transport.Submit(op)
}
