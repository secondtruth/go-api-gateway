package format

import "reflect"

type PlainFormatter struct {
}

func NewPlainFormatter() *PlainFormatter {
	return &PlainFormatter{}
}

func (f *PlainFormatter) FormatMsg(caption, text string) ([]byte, error) {
	var msg string
	if caption != "" {
		msg = caption + ": "
	}
	msg += text
	return []byte(msg), nil
}

func (f *PlainFormatter) FormatData(data any) ([]byte, error) {
	var msg string
	dtype := reflect.TypeOf(data)
	dval := reflect.ValueOf(data)
	if dtype.Kind() == reflect.Struct {
		for i := 0; i < dtype.NumField(); i++ {
			msg += dtype.Field(i).Name + ": " + dval.Field(i).String() + "\n"
		}
	} else if dtype.Kind() == reflect.Map {
		for _, key := range dval.MapKeys() {
			msg += key.String() + ": " + dval.MapIndex(key).String() + "\n"
		}
	} else {
		msg = dval.String()
	}
	return []byte(msg), nil
}

func (f *PlainFormatter) ContentType() string {
	return "text/plain"
}
