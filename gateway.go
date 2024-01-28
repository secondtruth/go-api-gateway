// Package gateway provides an API gateway that routes incoming requests to different handlers based on the host.
package gateway

import (
	"net"
	"net/http"
	"strings"

	"github.com/secondtruth/go-api-gateway/logging"
	"github.com/secondtruth/go-api-gateway/response"
	"github.com/secondtruth/go-api-gateway/services"
)

// Gateway represents an API gateway that routes incoming requests to different handlers based on the host.
type Gateway struct {
	entrypoints map[string]http.Handler

	Responder response.HttpResponder
	AccessLog logging.AccessLogger
}

// New creates a new instance of the API gateway.
func New() *Gateway {
	return &Gateway{
		entrypoints: make(map[string]http.Handler),
	}
}

// NewFromServiceContext creates a new instance of the API gateway using the provided service context.
func NewFromServiceContext(sc services.Context) *Gateway {
	g := New()
	g.Responder = sc.Responder
	g.AccessLog = sc.AccessLog
	return g
}

// HandleHost associates the given host with the provided handler.
func (g *Gateway) HandleHost(host string, handler http.Handler) {
	g.entrypoints[host] = handler
}

// HasHost checks if the gateway has a handler associated with the given host.
func (g *Gateway) HasHost(host string) bool {
	_, ok := g.findHost(host)
	return ok
}

// HasExactHost checks if the gateway has an exact match for the given host.
func (g *Gateway) HasExactHost(host string) bool {
	if strings.Contains(host, "*") || strings.Contains(host, "?") {
		panic("cannot use host wildcard as an exact host")
	}
	_, ok := g.entrypoints[host]
	return ok
}

// ListenAndServe starts the HTTP server and listens for incoming requests on the specified address.
func (g *Gateway) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, http.HandlerFunc(g.handle))
}

// ListenAndServeTLS starts the HTTPS server and listens for incoming requests on the specified address.
func (g *Gateway) ListenAndServeTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, http.HandlerFunc(g.handle))
}

// handle is the internal handler function that routes incoming requests to the appropriate handler based on the host.
func (g *Gateway) handle(w http.ResponseWriter, r *http.Request) {
	if g.AccessLog != nil {
		g.AccessLog.LogAccess(r)
	}

	host := r.Host
	if strings.Contains(host, ":") {
		var err error
		host, _, err = net.SplitHostPort(host)
		if err != nil {
			g.Responder.SendServerError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if handler, ok := g.findHost(host); ok {
		handler.ServeHTTP(w, r)
	} else {
		g.Responder.SendClientError(w, "entrypoint not found", http.StatusNotFound)
	}
}

// findHost finds the appropriate handler for the given host.
func (g *Gateway) findHost(host string) (http.Handler, bool) {
	for pattern, handler := range g.entrypoints {
		match := matchDomain(pattern, host)
		if match {
			return handler, true
		}
	}
	return nil, false
}
