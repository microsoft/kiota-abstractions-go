package abstractions

import (
	"context"
	"testing"
	"time"

	"github.com/microsoft/kiota-abstractions-go/store"

	"github.com/microsoft/kiota-abstractions-go/internal"
	s "github.com/microsoft/kiota-abstractions-go/serialization"
	assert "github.com/stretchr/testify/assert"
)

type QueryParameters struct {
	Count          *bool
	Expand         []string
	Filter         *string
	Orderby        []string
	Search         *string
	Select_escaped []string
	Skip           *int32
	Top            *int32
}

func TestItAddsStringQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	value := "somefilter"
	queryParameters := QueryParameters{
		Filter: &value,
	}
	requestInformation.AddQueryParameters(queryParameters)

	assert.Equal(t, value, requestInformation.QueryParameters["Filter"])
}

func TestItAddsBoolQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	value := true
	queryParameters := QueryParameters{
		Count: &value,
	}
	requestInformation.AddQueryParameters(queryParameters)
	assert.Equal(t, "true", requestInformation.QueryParameters["Count"])
}

func TestItAddsIntQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	value := int32(42)
	queryParameters := QueryParameters{
		Top: &value,
	}
	requestInformation.AddQueryParameters(queryParameters)
	assert.Equal(t, "42", requestInformation.QueryParameters["Top"])
}

func TestItAddsStringArrayQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	value := []string{"somefilter", "someotherfilter"}
	queryParameters := QueryParameters{
		Expand: value,
	}
	requestInformation.AddQueryParameters(queryParameters)
	assert.Equal(t, "somefilter,someotherfilter", requestInformation.QueryParameters["Expand"])
}

func TestItSetsTheRawURL(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.PathParameters[raw_url_key] = "https://someurl.com"
	requestInformation.UrlTemplate = "https://someotherurl.com{?select}"
	requestInformation.AddQueryParameters(QueryParameters{
		Select_escaped: []string{"somefield", "somefield2"},
	})
	uri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://someurl.com", uri.String())
	assert.Equal(t, 0, len(requestInformation.QueryParameters))
}

type getQueryParameters struct {
	Count          *bool    `uriparametername:"%24count"`
	Expand         []string `uriparametername:"%24expand"`
	Select_escaped []string `uriparametername:"%24select"`
	Filter         *string  `uriparametername:"%24filter"`
	Orderby        []string `uriparametername:"%24orderby"`
	Search         *string  `uriparametername:"%24search"`
}

func TestItSetsSelectAndCountQueryParameters(t *testing.T) {
	value := true
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "http://localhost/me{?%24select,%24count}"
	requestInformation.AddQueryParameters(getQueryParameters{
		Select_escaped: []string{"id", "displayName"},
		Count:          &value,
	})
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/me?%24select=id%2CdisplayName&%24count=true", resultUri.String())
}

func TestItDoesNotSetEmptySelectQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "http://localhost/me{?%24select}"
	requestInformation.AddQueryParameters(getQueryParameters{
		Select_escaped: []string{},
	})
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/me", resultUri.String())
}

func TestItDoesNotSetEmptySearchQueryParameters(t *testing.T) {
	emptyString := ""
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "http://localhost/me{?%24search}"
	requestInformation.AddQueryParameters(getQueryParameters{
		Search: &emptyString,
	})
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/me", resultUri.String())
}

func TestItSetsPathParametersOfDateTimeOffsetType(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "http://localhost/getDirectRoutingCalls(fromDateTime='{fromDateTime}',toDateTime='{toDateTime}')"

	fromDateTime := time.Date(2022, 8, 1, 20, 34, 58, 0, time.UTC)
	toDateTime := time.Date(2022, 8, 2, 20, 34, 58, 0, time.UTC)

	requestInformation.PathParameters["fromDateTime"] = fromDateTime.Format(time.RFC3339)
	requestInformation.PathParameters["toDateTime"] = toDateTime.Format(time.RFC3339)

	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Contains(t, resultUri.String(), "fromDateTime='2022-08-01T20%3A34%3A58Z'")
	assert.Contains(t, resultUri.String(), "toDateTime='2022-08-02T20%3A34%3A58Z'")
}

func TestItErrorsWhenBaseUrlNotSet(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/users{?%24count}"

	_, err := requestInformation.GetUri()
	assert.NotNil(t, err)
}

func TestItBuildsUrlOnProvidedBaseUrl(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/users{?%24count}"

	requestInformation.PathParameters["baseurl"] = "http://localhost"

	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/users", resultUri.String())
}

func TestItSetsContentFromParsable(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/users{?%24count}"
	requestInformation.Method = POST

	callsCounter := make(map[string]int)
	requestAdapter := &MockRequestAdapter{
		SerializationWriterFactory: &internal.MockSerializerFactory{
			SerializationWriter: &internal.MockSerializer{
				CallsCounter: callsCounter,
			},
		},
	}

	requestInformation.PathParameters["baseurl"] = "http://localhost"

	record := internal.CallRecord{}
	err := requestInformation.SetContentFromParsable(context.Background(), requestAdapter, "application/json", &record)
	assert.Nil(t, err)
	assert.Equal(t, 1, callsCounter["WriteObjectValue"])
	assert.Equal(t, 0, callsCounter["WriteCollectionOfObjectValues"])
}
func TestItSetsContentFromParsableCollection(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/users{?%24count}"
	requestInformation.Method = POST

	callsCounter := make(map[string]int)
	requestAdapter := &MockRequestAdapter{
		SerializationWriterFactory: &internal.MockSerializerFactory{
			SerializationWriter: &internal.MockSerializer{
				CallsCounter: callsCounter,
			},
		},
	}

	requestInformation.PathParameters["baseurl"] = "http://localhost"

	record := internal.CallRecord{}
	err := requestInformation.SetContentFromParsableCollection(context.Background(), requestAdapter, "application/json", []s.Parsable{&record})
	assert.Nil(t, err)
	assert.Equal(t, 0, callsCounter["WriteObjectValue"])
	assert.Equal(t, 1, callsCounter["WriteCollectionOfObjectValues"])
}
func TestItSetsContentFromScalarCollection(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/users{?%24count}"
	requestInformation.Method = POST

	callsCounter := make(map[string]int)
	requestAdapter := &MockRequestAdapter{
		SerializationWriterFactory: &internal.MockSerializerFactory{
			SerializationWriter: &internal.MockSerializer{
				CallsCounter: callsCounter,
			},
		},
	}

	requestInformation.PathParameters["baseurl"] = "http://localhost"

	vals := []any{"foo"}
	err := requestInformation.SetContentFromScalarCollection(context.Background(), requestAdapter, "application/json", vals)
	assert.Nil(t, err)
	assert.Equal(t, 0, callsCounter["WriteStringValue"])
	assert.Equal(t, 1, callsCounter["WriteCollectionOfStringValues"])
}

func TestItSetsContentFromScalar(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/users{?%24count}"
	requestInformation.Method = POST

	callsCounter := make(map[string]int)
	requestAdapter := &MockRequestAdapter{
		SerializationWriterFactory: &internal.MockSerializerFactory{
			SerializationWriter: &internal.MockSerializer{
				CallsCounter: callsCounter,
			},
		},
	}

	requestInformation.PathParameters["baseurl"] = "http://localhost"

	value := "foo"
	err := requestInformation.SetContentFromScalar(context.Background(), requestAdapter, "application/json", &value)
	assert.Nil(t, err)
	assert.Equal(t, 1, callsCounter["WriteStringValue"])
	assert.Equal(t, 0, callsCounter["WriteCollectionOfStringValues"])
}

func TestItSetsTheBoundaryOnMultipartBody(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/users{?%24count}"
	requestInformation.Method = POST

	callsCounter := make(map[string]int)
	requestAdapter := &MockRequestAdapter{
		SerializationWriterFactory: &internal.MockSerializerFactory{
			SerializationWriter: &internal.MockSerializer{
				CallsCounter: callsCounter,
			},
		},
	}

	requestInformation.PathParameters["baseurl"] = "http://localhost"

	multipartBody := NewMultipartBody()
	err := requestInformation.SetContentFromParsable(context.Background(), requestAdapter, "multipart/form-data", multipartBody)
	assert.Nil(t, err)
	contentTypeHeader := requestInformation.Headers.Get("Content-Type")
	assert.NotNil(t, contentTypeHeader)
	contentTypeHeaderValue := contentTypeHeader[0]
	assert.Equal(t, "multipart/form-data; boundary="+multipartBody.GetBoundary(), contentTypeHeaderValue)
}

type MockRequestAdapter struct {
	SerializationWriterFactory s.SerializationWriterFactory
}

func (r *MockRequestAdapter) Send(context context.Context, requestInfo *RequestInformation, constructor s.ParsableFactory, errorMappings ErrorMappings) (s.Parsable, error) {
	return nil, nil
}
func (r *MockRequestAdapter) SendEnum(context context.Context, requestInfo *RequestInformation, parser s.EnumFactory, errorMappings ErrorMappings) (any, error) {
	return nil, nil
}
func (r *MockRequestAdapter) SendCollection(context context.Context, requestInfo *RequestInformation, constructor s.ParsableFactory, errorMappings ErrorMappings) ([]s.Parsable, error) {
	return nil, nil
}
func (r *MockRequestAdapter) SendEnumCollection(context context.Context, requestInfo *RequestInformation, parser s.EnumFactory, errorMappings ErrorMappings) ([]any, error) {
	return nil, nil
}
func (r *MockRequestAdapter) SendPrimitive(context context.Context, requestInfo *RequestInformation, typeName string, errorMappings ErrorMappings) (any, error) {
	return nil, nil
}
func (r *MockRequestAdapter) SendPrimitiveCollection(context context.Context, requestInfo *RequestInformation, typeName string, errorMappings ErrorMappings) ([]any, error) {
	return nil, nil
}
func (r *MockRequestAdapter) SendNoContent(context context.Context, requestInfo *RequestInformation, errorMappings ErrorMappings) error {
	return nil
}
func (r *MockRequestAdapter) ConvertToNativeRequest(context context.Context, requestInfo *RequestInformation) (any, error) {
	return nil, nil
}
func (r *MockRequestAdapter) GetSerializationWriterFactory() s.SerializationWriterFactory {
	return r.SerializationWriterFactory
}
func (r *MockRequestAdapter) EnableBackingStore(factory store.BackingStoreFactory) {
}
func (r *MockRequestAdapter) SetBaseUrl(baseUrl string) {
}
func (r *MockRequestAdapter) GetBaseUrl() string {
	return ""
}
