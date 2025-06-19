package proxy

import (
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(target string) (*httputil.ReverseProxy, error) {
	targetUrl, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(targetUrl), nil
}
