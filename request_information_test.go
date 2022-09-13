package abstractions

import (
	"testing"
	"time"

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
