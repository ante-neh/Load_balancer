package types

import "net/http/httputil"

type Server struct {
	addr  string
	proxy httputil.ReverseProxy
}