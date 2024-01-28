package format

type HttpResponseFormatter interface {
	FormatMsg(caption, text string) ([]byte, error)
	FormatData(data any) ([]byte, error)
	ContentType() string
}
