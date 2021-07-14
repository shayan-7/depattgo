package behavioral

import "errors"

const (
	ValidType          = "application/json"
	ValidLength        = 100
	ValidAuthorization = "a.b.c"
	ValidVersion       = 2.0
)

const (
	ErrInvalidContentType   = "invalid content type"
	ErrInvalidContentLength = "invalid content length"
	ErrInvalidAuthorization = "invalid content authorization token"
)

type Request struct {
	ContentType   string
	ContentLength int
	Authorization string
}

func NewRequest(typ string, length int, auth string) *Request {
	return &Request{typ, length, auth}
}

type Validator interface {
	NextHandler() Validator
	Validate(*Request) error
}

// ReqTypeValidator checks if the Content-Type of request is the valid type
type ReqTypeValidator struct {
	v Validator
}

func NewReqTypeValidator(v Validator) *ReqTypeValidator {
	return &ReqTypeValidator{v}
}

func (rtv *ReqTypeValidator) Validate(r *Request) error {
	var err error
	if r.ContentType != ValidType {
		err = errors.New(ErrInvalidContentType)
	} else if rtv.NextHandler() != nil {
		err = rtv.NextHandler().Validate(r)
	}
	return err
}

func (rtv *ReqTypeValidator) NextHandler() Validator {
	return rtv.v
}

// ReqTypeValidator checks if the Content-Length of request is the valid type
type ReqLengthValidator struct {
	v Validator
}

func NewReqLengthValidator(v Validator) *ReqLengthValidator {
	return &ReqLengthValidator{v}
}

func (rlv *ReqLengthValidator) Validate(r *Request) error {
	var err error
	if r.ContentLength > ValidLength {
		err = errors.New(ErrInvalidContentLength)
	} else if rlv.NextHandler() != nil {
		err = rlv.NextHandler().Validate(r)
	}
	return err
}

func (rlv *ReqLengthValidator) NextHandler() Validator {
	return rlv.v
}

// ReqTypeValidator checks if the Authorization exists in request
type ReqTokenValidator struct {
	v Validator
}

func NewReqTokenValidator(v Validator) *ReqTokenValidator {
	return &ReqTokenValidator{v}
}

func (rtv *ReqTokenValidator) Validate(r *Request) error {
	var err error
	if r.Authorization != ValidAuthorization {
		err = errors.New(ErrInvalidAuthorization)
	} else if rtv.NextHandler() != nil {
		err = rtv.NextHandler().Validate(r)
	}
	return err
}

func (rtv *ReqTokenValidator) NextHandler() Validator {
	return rtv.v
}
