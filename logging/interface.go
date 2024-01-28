package logging

import "net/http"

// AccessLogger is an interface for logging HTTP requests to endpoints.
type AccessLogger interface {
	LogAccess(request *http.Request)
}
