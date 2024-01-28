package format

import "encoding/json"

type JsonFormatter struct {
}

func NewJsonFormatter() *JsonFormatter {
	return &JsonFormatter{}
}

func (f *JsonFormatter) FormatMsg(caption, text string) ([]byte, error) {
	return json.Marshal(map[string]string{
		caption: text,
	})
}

func (f *JsonFormatter) FormatData(data any) ([]byte, error) {
	return json.Marshal(data)
}

func (f *JsonFormatter) ContentType() string {
	return "application/json"
}
