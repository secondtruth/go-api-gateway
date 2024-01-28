package logging

import (
	"io"
	"log"
	"net/http"
)

// SimpleAccessLog represents a simple access log for logging HTTP requests. Implements the AccessLogger interface.
type SimpleAccessLog struct {
	handler log.Logger
}

// NewSimpleAccessLog creates a new instance of SimpleAccessLog using an io.Writer, a prefix string, and a flag as parameters.
func NewSimpleAccessLog(out io.Writer, prefix string, flag int) *SimpleAccessLog {
	return &SimpleAccessLog{
		handler: *log.New(out, prefix, flag),
	}
}

// LogAccess logs the access details of an HTTP request.
func (l *SimpleAccessLog) LogAccess(r *http.Request) {
	l.handler.Printf("[%s] %s %s", r.Host, r.Method, r.URL.Path)
}
