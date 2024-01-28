package handlers

import (
	"fmt"
	"net/http"

	"github.com/secondtruth/go-api-gateway/services"
	"github.com/julienschmidt/httprouter"
)

// Router creates a new HTTP router instance
func Router(sc services.Context) *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, val any) {
		sc.Responder.SendServerError(w, fmt.Sprintf("%v", val), http.StatusInternalServerError)
	}
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sc.Responder.SendClientError(w, "endpoint not found", http.StatusNotFound)
	})
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sc.Responder.SendClientError(w, "method not allowed", http.StatusMethodNotAllowed)
	})
	return router
}
