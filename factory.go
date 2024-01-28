package gateway

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/secondtruth/go-api-gateway/handlers"
	"github.com/secondtruth/go-api-gateway/services"
	reverseproxy "github.com/secondtruth/go-reverse-proxy"
)

// HandlerFactory creates HTTP handlers preconfigured using the service context and other configuration options.
type HandlerFactory struct {
	Transport      http.RoundTripper
	RequestHeader  http.Header
	ModifyResponse func(*http.Response) error
	Services       services.Context
}

// NewHandlerFactory creates a new instance of HandlerFactory with the given service context.
func NewHandlerFactory(sc services.Context) *HandlerFactory {
	return &HandlerFactory{
		Services: sc,
	}
}

// MakeRouter creates a preconfigured httprouter.Router instance.
func (f *HandlerFactory) MakeRouter() *httprouter.Router {
	return handlers.Router(f.Services)
}

// MakeReverseProxy creates a preconfigured reverseproxy.ReverseProxyMux instance for the given remote host.
func (f *HandlerFactory) MakeReverseProxy(remote string) (*reverseproxy.ReverseProxyMux, error) {
	rp, err := handlers.ReverseProxy(remote, f.Services)
	rp.Transport = f.Transport
	rp.RequestHeader = f.RequestHeader
	rp.ModifyResponse = f.ModifyResponse
	return rp, err
}
