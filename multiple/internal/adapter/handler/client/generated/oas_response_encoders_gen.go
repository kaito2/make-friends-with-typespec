// Code generated by ogen, DO NOT EDIT.

package generated

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeStoresListResponse(response []Store, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := new(jx.Encoder)
	e.ArrStart()
	for _, elem := range response {
		elem.Encode(e)
	}
	e.ArrEnd()
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeStoresReadResponse(response *Store, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := new(jx.Encoder)
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}
