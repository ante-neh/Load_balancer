package types

import "net/http/httputil"

type Server struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}