package abstractions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testHandler(response interface{}, errorMappings ErrorMappings) (interface{}, error) {
	return nil, nil
}

func TestRequestHandlerOption(t *testing.T) {

	handlerOption := NewRequestHandlerOption()
	handlerOption.SetResponseHandler(testHandler)
	assert.NotNil(t, handlerOption)
	assert.NotNil(t, handlerOption.GetKey())
	assert.NotNil(t, handlerOption.GetResponseHandler())
	assert.Implements(t, (*RequestHandlerOption)(nil), handlerOption)
	assert.Implements(t, (*RequestOption)(nil), handlerOption)
}
func TestRequestHandlerOptionNilReceiver(t *testing.T) {
	var nilOption *requestHandlerOption
	assert.Nil(t, nilOption.GetResponseHandler())
	nilOption.SetResponseHandler(testHandler)
	assert.Equal(t, ResponseHandlerOptionKey, nilOption.GetKey())
}
