package abstractions

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/microsoft/kiota-abstractions-go/store"

	"github.com/microsoft/kiota-abstractions-go/internal"
	s "github.com/microsoft/kiota-abstractions-go/serialization"
	assert "github.com/stretchr/testify/assert"
)

type QueryParameters struct {
	Count          *bool
	Expand         []string
	ExpandAny      []any
	Filter         *string
	Orderby        []string
	Search         *string
	Select_escaped []string
	Skip           *int32
	Top            *int32
	Status         *internal.PersonStatus  `uriparametername:"status"`
	Statuses       []internal.PersonStatus `uriparametername:"statuses"`
	Id             *uuid.UUID
}

func TestItAddsStringQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	value := "somefilter"
	queryParameters := QueryParameters{
		Filter: &value,
	}
	requestInformation.AddQueryParameters(queryParameters)

	assert.Equal(t, value, requestInformation.QueryParameters["Filter"])
	assert.Nil(t, requestInformation.QueryParametersAny["Filter"])
}

func TestItAddsBoolQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	value := true
	queryParameters := QueryParameters{
		Count: &value,
	}
	requestInformation.AddQueryParameters(queryParameters)
	assert.Equal(t, "true", requestInformation.QueryParameters["Count"])
	assert.Nil(t, requestInformation.QueryParametersAny["Count"])
}

func TestItAddsIntQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	value := int32(42)
	queryParameters := QueryParameters{
		Top: &value,
	}
	requestInformation.AddQueryParameters(queryParameters)
	assert.Equal(t, "42", requestInformation.QueryParameters["Top"])
	assert.Nil(t, requestInformation.QueryParametersAny["Top"])
}

func TestItAddsStringArrayQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	value := []string{"somefilter", "someotherfilter"}
	queryParameters := QueryParameters{
		Expand: value,
	}
	requestInformation.AddQueryParameters(queryParameters)
	assert.Equal(t, "somefilter,someotherfilter", requestInformation.QueryParameters["Expand"])
	assert.Equal(t, []any{"somefilter", "someotherfilter"}, requestInformation.QueryParametersAny["Expand"])
}

func TestItAddsAnyArrayQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	value := []any{"somefilter", "someotherfilter"}
	queryParameters := QueryParameters{
		ExpandAny: value,
	}
	requestInformation.AddQueryParameters(queryParameters)
	assert.Empty(t, requestInformation.QueryParameters["ExpandAny"])
	assert.Equal(t, []any{"somefilter", "someotherfilter"}, requestInformation.QueryParametersAny["ExpandAny"])
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
	assert.Equal(t, 0, len(requestInformation.QueryParametersAny))
}

type getQueryParameters struct {
	Count          *bool      `uriparametername:"%24count"`
	Expand         []string   `uriparametername:"%24expand"`
	Select_escaped []string   `uriparametername:"%24select"`
	Filter         *string    `uriparametername:"%24filter"`
	Orderby        []string   `uriparametername:"%24orderby"`
	Search         *string    `uriparametername:"%24search"`
	Number         []int64    `uriparametername:"%24number"`
	Id             *uuid.UUID `uriparametername:"id"`
}

func TestItSetsUUIDQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "http://localhost/users{?id}"
	value := uuid.MustParse("95E943B8-52D5-4228-902D-61D65792CED7")
	requestInformation.AddQueryParameters(getQueryParameters{
		Id: &value,
	})
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/users?id=95e943b8-52d5-4228-902d-61d65792ced7", resultUri.String())
}

func TestItSetsUUIDPathParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "http://localhost/users/{id}"
	value := uuid.MustParse("95E943B8-52D5-4228-902D-61D65792CED7")
	requestInformation.PathParametersAny["id"] = &value
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/users/95e943b8-52d5-4228-902d-61d65792ced7", resultUri.String())
}

func TestItSetsNumberArrayQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "http://localhost/me{?%24number}"
	requestInformation.AddQueryParameters(getQueryParameters{
		Number: []int64{1, 2, 4},
	})
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/me?%24number=1,2,4", resultUri.String())
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
	assert.Equal(t, "http://localhost/me?%24select=id,displayName&%24count=true", resultUri.String())
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
	assert.Equal(t, "http://localhost/me?%24search=", resultUri.String())
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

func TestItSetsEnumValueInQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/users{?status}"

	requestInformation.PathParameters["baseurl"] = "http://localhost"

	status := internal.ACTIVE
	requestInformation.AddQueryParameters(QueryParameters{
		Status: &status,
	})
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/users?status=active", resultUri.String())
}

func TestItSetsEnumValuesInQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/users{?statuses}"

	requestInformation.PathParameters["baseurl"] = "http://localhost"

	statuses := make([]internal.PersonStatus, 2)
	statuses[0] = internal.ACTIVE
	statuses[1] = internal.SUSPENDED
	requestInformation.AddQueryParameters(QueryParameters{
		Statuses: statuses,
	})
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/users?statuses=active,suspended", resultUri.String())
}

func TestItSetsEnumValueInPathParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/{status}"

	status := internal.ACTIVE
	requestInformation.PathParameters["baseurl"] = "http://localhost"
	requestInformation.PathParametersAny["status"] = &status

	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/active", resultUri.String())
}

func TestItSetsEnumValuesInPathParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/{statuses}"

	statuses := make([]internal.PersonStatus, 2)
	statuses[0] = internal.ACTIVE
	statuses[1] = internal.SUSPENDED
	requestInformation.PathParameters["baseurl"] = "http://localhost"
	requestInformation.PathParametersAny["statuses"] = statuses

	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/active,suspended", resultUri.String())
}

func prepareNormalizedStdTest(arrayValues any, singleValue any, referenceArray any, referenceValue any) *RequestInformation {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "{+baseurl}/array/{arrayValues}/single/{singleValue}/referenceArray/{referenceArray}/referenceValue/{referenceValue}"

	requestInformation.PathParameters["baseurl"] = "http://localhost"
	requestInformation.PathParametersAny["arrayValues"] = arrayValues
	requestInformation.PathParametersAny["singleValue"] = singleValue
	requestInformation.PathParametersAny["referenceArray"] = referenceArray
	requestInformation.PathParametersAny["referenceValue"] = referenceValue

	return requestInformation
}

func TestItNormalizesOnStandardizedTimeParams(t *testing.T) {
	arrayValues := []time.Time{
		time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 8, 2, 0, 0, 0, 0, time.UTC),
	}

	singleValue := time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC)

	time1 := time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC)
	time2 := time.Date(2022, 8, 2, 0, 0, 0, 0, time.UTC)
	referenceArray := []*time.Time{
		&time1,
		&time2,
	}

	referenceValue := &time1

	requestInformation := prepareNormalizedStdTest(arrayValues, singleValue, referenceArray, referenceValue)
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/array/2022-08-01%2000%3A00%3A00%20%2B0000%20UTC,2022-08-02%2000%3A00%3A00%20%2B0000%20UTC/single/2022-08-01T00%3A00%3A00Z/referenceArray/2022-08-01%2000%3A00%3A00%20%2B0000%20UTC,2022-08-02%2000%3A00%3A00%20%2B0000%20UTC/referenceValue/2022-08-01T00%3A00%3A00Z", resultUri.String())

}

func TestItNormalizesOnStandardizedDurationParams(t *testing.T) {
	duration := s.NewDuration(0, 0, 1, 1, 0, 0, 0)
	duration1 := s.NewDuration(0, 0, 1, 2, 0, 0, 0)
	arrayValues := []s.ISODuration{
		*duration,
		*duration1,
	}

	value := s.NewDuration(0, 0, 2, 1, 0, 0, 0)
	singleValue := *value

	duration3 := s.NewDuration(0, 0, 1, 2, 0, 0, 0)
	referenceArray := []*s.ISODuration{
		duration3,
	}

	referenceValue := s.NewDuration(0, 0, 1, 1, 0, 0, 0)

	requestInformation := prepareNormalizedStdTest(arrayValues, singleValue, referenceArray, referenceValue)
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/array/P1DT1H,P1DT2H/single/P2DT1H/referenceArray/P1DT2H/referenceValue/P1DT1H", resultUri.String())
}

func TestItNormalizesOnStandardizedTimeOnlyParams(t *testing.T) {
	time1, err := s.ParseTimeOnly("16:20:21.000")
	assert.Nil(t, err)
	arrayValues := []s.TimeOnly{
		*time1,
	}

	value, err := s.ParseTimeOnly("16:20:21.000")
	assert.Nil(t, err)
	singleValue := *value

	time2, err := s.ParseTimeOnly("16:20:21.000")
	referenceArray := []*s.TimeOnly{
		time2,
	}

	referenceValue, err := s.ParseTimeOnly("16:20:21.000")

	requestInformation := prepareNormalizedStdTest(arrayValues, singleValue, referenceArray, referenceValue)
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/array/16%3A20%3A21.000000000/single/16%3A20%3A21.000000000/referenceArray/16%3A20%3A21.000000000/referenceValue/16%3A20%3A21.000000000", resultUri.String())
}

func TestItNormalizesOnStandardizedDateOnlyParams(t *testing.T) {
	dateOnly := s.NewDateOnly(time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC))
	arrayValues := []s.DateOnly{
		*dateOnly,
	}

	date1 := s.NewDateOnly(time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC))
	singleValue := *date1

	referenceArray := []*s.DateOnly{
		s.NewDateOnly(time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC)),
	}

	referenceValue := s.NewDateOnly(time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC))

	requestInformation := prepareNormalizedStdTest(arrayValues, singleValue, referenceArray, referenceValue)
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/array/2020-01-04/single/2020-01-04/referenceArray/2020-01-04/referenceValue/2020-01-04", resultUri.String())
}

func TestItSetsExplodedQueryParameters(t *testing.T) {
	value := true
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "http://localhost/me{?%24select*}"
	requestInformation.AddQueryParameters(getQueryParameters{
		Select_escaped: []string{"id", "displayName"},
		Count:          &value,
	})
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/me?%24select=id&%24select=displayName", resultUri.String())
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
