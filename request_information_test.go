package abstractions

import (
	"testing"

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
	Count          *bool    `uriparamtername:"%24count"`
	Expand         []string `uriparamtername:"%24expand"`
	Select_escaped []string `uriparamtername:"%24select"`
	Filter         *string  `uriparamtername:"%24filter"`
	Orderby        []string `uriparamtername:"%24orderby"`
	Search         *string  `uriparamtername:"%24search"`
}

func TestItSetsSelectQueryParameters(t *testing.T) {
	requestInformation := NewRequestInformation()
	requestInformation.UrlTemplate = "http://localhost/me{?%24select}"
	requestInformation.AddQueryParameters(getQueryParameters{
		Select_escaped: []string{"id", "displayName"},
	})
	resultUri, err := requestInformation.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost/me?%24select=id%2CdisplayName", resultUri.String())
}
