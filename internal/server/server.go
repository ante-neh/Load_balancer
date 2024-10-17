package server

import (
	"net/http/httputil"
	"net/url"

	"github.com/ante-neh/Load_balancer/types"
	"github.com/ante-neh/Load_balancer/util"
)

func NewServer(addr string) types.Server{
	serverUrl, err := url.Parse(addr)
	util.HandleError(err) 

	return types.Server{
		Addr: addr,
		Proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}