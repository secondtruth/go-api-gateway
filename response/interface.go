package response

import (
	"net/http"

	"github.com/secondtruth/go-api-gateway/response/format"
)

// HttpResponder represents the component that sends HTTP messages in response to client requests
type HttpResponder interface {
	// SetFormatter sets the formatter to be used
	SetFormatter(formatter format.HttpResponseFormatter)

	// Send writes a byte array to the response body
	Send(w http.ResponseWriter, content []byte, code int) error

	// SendMsg writes a message with an optional caption to the response body
	SendMsg(w http.ResponseWriter, text, caption string, code int) error
	
	// SendData writes a data object to the response body
	SendData(w http.ResponseWriter, data any) error

	// SendClientError writes a client error message to the response body
	SendClientError(w http.ResponseWriter, msg string, code int) error

	// SendServerError writes a server error message to the response body
	SendServerError(w http.ResponseWriter, msg string, code int) error
}
