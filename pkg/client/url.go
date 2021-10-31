// Date: 2021/10/24

// Package client
package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"caty/pkg/v"
)

type URLHandler interface {
	Url(ctx context.Context, path string) string
	UrlWithQuery(ctx context.Context, path string, value url.Values) string
	Header(ctx context.Context) http.Header
}

func NewURLHandler() URLHandler {
	return DefaultIP{}
}

type DefaultIP struct {
}

func (d DefaultIP) Header(ctx context.Context) http.Header {
	header := make(http.Header)
	traceID := v.GetTraceID(ctx)
	if traceID != "" {
		header.Add(v.XTraceID, traceID)
	}
	return header
}

func (d DefaultIP) UrlWithQuery(ctx context.Context, path string, value url.Values) string {
	if len(value) == 0 {
		return d.Url(ctx, path)
	}
	return d.Url(ctx, path) + "?" + value.Encode()
}

func (d DefaultIP) Url(ctx context.Context, path string) string {
	host := v.GetHost(ctx)
	if host == "" {
		host = "127.0.0.1:8120"
	}
	return fmt.Sprintf("http://%s%s", host, path)
}
