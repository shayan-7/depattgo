package behavioral

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainOfResponsibility(t *testing.T) {
	tokenValidator := NewReqTokenValidator(nil)
	typeValidator := NewReqTypeValidator(tokenValidator)
	lengthValidator := NewReqLengthValidator(typeValidator)
	startingValidator := lengthValidator

	cases := []struct {
		desc string // description
		req  *Request
		err  error
	}{
		{
			"Request is valid",
			NewRequest(ValidType, ValidLength, ValidAuthorization),
			nil,
		},
		{
			"Authorization is not valid",
			NewRequest(ValidType, ValidLength, "abc"),
			errors.New(ErrInvalidAuthorization),
		},
		{
			"Content length is not valid",
			NewRequest(ValidType, 1000, ValidAuthorization),
			errors.New(ErrInvalidContentLength),
		},
		{
			"Content type is not valid",
			NewRequest("text/plain", ValidLength, ValidAuthorization),
			errors.New(ErrInvalidContentType),
		},
	}
	for _, c := range cases {
		err := startingValidator.Validate(c.req)
		if c.err == nil {
			assert.NoError(t, err)
		} else {
			assert.NotNil(t, err)
			assert.EqualError(t, c.err, err.Error())
		}
	}
}
