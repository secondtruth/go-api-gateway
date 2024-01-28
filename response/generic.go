package response

import (
	"errors"
	"net/http"

	"github.com/secondtruth/go-api-gateway/response/format"
)

// GenericResponder is the default implementation of the HttpResponder interface
type GenericResponder struct {
	formatter format.HttpResponseFormatter
}

func NewGenericResponder(formatter format.HttpResponseFormatter) *GenericResponder {
	if formatter == nil {
		formatter = format.NewPlainFormatter()
	}
	return &GenericResponder{
		formatter: formatter,
	}
}

func (r *GenericResponder) SetFormatter(formatter format.HttpResponseFormatter) {
	r.formatter = formatter
}

func (r *GenericResponder) Send(w http.ResponseWriter, content []byte, code int) error {
	w.Header().Set("Content-Type", r.formatter.ContentType())
	w.WriteHeader(code)
	_, err := w.Write(content)
	return err
}

func (r *GenericResponder) SendMsg(w http.ResponseWriter, text, caption string, code int) error {
	b, err := r.formatter.FormatMsg(caption, text)
	if err != nil {
		return err
	}
	return r.Send(w, b, code)
}

func (r *GenericResponder) SendData(w http.ResponseWriter, data any) error {
	b, err := r.formatter.FormatData(data)
	if err != nil {
		return err
	}
	return r.Send(w, b, http.StatusOK)
}

func (r *GenericResponder) SendClientError(w http.ResponseWriter, msg string, code int) error {
	if code < 400 || code >= 500 {
		return errors.New("invalid client error code")
	}

	return r.SendMsg(w, msg, "error", code)
}

func (r *GenericResponder) SendServerError(w http.ResponseWriter, msg string, code int) error {
	if code < 500 {
		return errors.New("invalid server error code")
	}

	return r.SendMsg(w, msg, "error", code)
}
