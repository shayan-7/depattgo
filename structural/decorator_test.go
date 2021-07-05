package structural

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	r := NewBaseResponse()

	// Before setting header by base response
	h := r.GetHeader()
	k, ok := h[StatusKey]
	assert.False(t, ok)
	assert.Equal(t, "", k)

	r.SetHeader()

	// After setting header by base response
	h = r.GetHeader()
	k, ok = h[StatusKey]
	assert.True(t, ok)
	assert.Equal(t, "200 OK", k)

	// Before using the JSON decorator
	k, ok = h[ContentTypeKey]
	assert.False(t, ok)
	assert.Equal(t, "", k)

	r = NewJSONDecorator(r)
	r.SetHeader()

	// After using the JSON decorator
	k, ok = h[ContentTypeKey]
	assert.True(t, ok)
	assert.Equal(t, "application/json", k)

	// Before using the Gzip decorator
	k, ok = h[EncodingKey]
	assert.False(t, ok)
	assert.Equal(t, "", k)

	// After using the Gzip decorator
	r = NewGzipDecorator(r)
	r.SetHeader()
	k, ok = h[EncodingKey]
	assert.True(t, ok)
	assert.Equal(t, "gzip", k)
}
