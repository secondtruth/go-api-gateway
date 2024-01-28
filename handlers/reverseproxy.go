package handlers

import (
	"net/http"

	"github.com/secondtruth/go-api-gateway/services"
	reverseproxy "github.com/secondtruth/go-reverse-proxy"
)

// ReverseProxy creates a new reverse proxy instance
func ReverseProxy(remote string, sc services.Context) (*reverseproxy.ReverseProxyMux, error) {
	rp, err := reverseproxy.New(remote)
	if err != nil {
		return nil, err
	}
	rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		sc.Responder.SendServerError(w, err.Error(), http.StatusInternalServerError)
	}
	rp.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sc.Responder.SendClientError(w, "endpoint not found", http.StatusNotFound)
	})
	rp.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sc.Responder.SendClientError(w, "method not allowed", http.StatusMethodNotAllowed)
	})
	return rp, nil
}
