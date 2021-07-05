package structural

type HeaderKey string

const (
	StatusKey      HeaderKey = "Status"
	ContentTypeKey HeaderKey = "Content-Type"
	EncodingKey    HeaderKey = "Encoding"
)

type Header map[HeaderKey]string

type response interface {
	SetHeader()
	GetHeader() Header
}

type BaseResponse struct {
	Header
}

func NewBaseResponse() response {
	return &BaseResponse{Header: make(Header)}
}

func (br *BaseResponse) SetHeader() {
	br.Header["Status"] = "200 OK"
}

func (br *BaseResponse) GetHeader() Header {
	return br.Header
}

type JSONDecorator struct {
	wrappee response
}

func NewJSONDecorator(wrappee response) response {
	return &JSONDecorator{wrappee}
}

func (jd *JSONDecorator) SetHeader() {
	h := jd.wrappee.GetHeader()
	h[ContentTypeKey] = "application/json"
}

func (jd *JSONDecorator) GetHeader() Header {
	return jd.wrappee.GetHeader()
}

type GzipDecorator struct {
	wrappee response
}

func NewGzipDecorator(wrappee response) response {
	return &GzipDecorator{wrappee}
}

func (gd *GzipDecorator) SetHeader() {
	h := gd.wrappee.GetHeader()
	h[EncodingKey] = "gzip"
}

func (gd *GzipDecorator) GetHeader() Header {
	return gd.wrappee.GetHeader()
}
