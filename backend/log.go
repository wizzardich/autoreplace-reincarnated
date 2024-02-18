package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
)

func FullRequest(r *http.Request) slog.Attr {
	attrs := []slog.Attr{
		slog.String("method", r.Method),
		slog.String("uri", r.RequestURI),
	}

	if r.Body != nil {
		body, err := RequestBody(r)
		if err != nil {
			body = slog.String("body_error", err.Error())
		}
		attrs = append(attrs, body)
	}

	return Group("request", attrs...)
}

func RequestBody(r *http.Request) (slog.Attr, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		return slog.Attr{}, err
	}
	r.Body = io.NopCloser(buf)
	return slog.String("body", buf.String()), nil
}

func Group(key string, attrs ...slog.Attr) slog.Attr {
	return Attr(key, slog.GroupValue(attrs...))
}

func Attr(key string, value slog.Value) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: value,
	}
}
