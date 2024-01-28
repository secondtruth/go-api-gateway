package services

import (
	"github.com/secondtruth/go-api-gateway/logging"
	"github.com/secondtruth/go-api-gateway/response"
)

// Context represents the service context, which is used to pass around common service dependencies.
type Context struct {
	Responder response.HttpResponder
	AccessLog logging.AccessLogger
}

// NewContextWithDefaults creates a new service context with default values.
func NewContextWithDefaults() Context {
	return Context{
		Responder: response.NewGenericResponder(nil),
	}
}
