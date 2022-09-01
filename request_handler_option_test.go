package abstractions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testHandler(response interface{}, errorMappings ErrorMappings) (interface{}, error) {
	return nil, nil
}

func TestRequestHandlerOption(t *testing.T) {

	handlerOption := NewRequestHandlerOption(testHandler)

	assert.NotNil(t, handlerOption)
	assert.NotNil(t, handlerOption.GetKey())
	assert.NotNil(t, handlerOption.GetResponseHandler())
	assert.Implements(t, (*RequestHandlerOption)(nil), handlerOption)
	assert.Implements(t, (*RequestOption)(nil), handlerOption)
}
