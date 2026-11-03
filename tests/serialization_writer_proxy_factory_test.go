package tests

import (
	"errors"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerializationWriterProxyPreservesStartCallback(t *testing.T) {
	callbackError := errors.New("serialization callback failed")
	for _, test := range []struct {
		name          string
		before        bool
		start         bool
		proxyError    error
		originalError error
		wantCalls     []string
		wantError     error
	}{
		{name: "neither callback", wantCalls: []string{"proxy"}},
		{name: "start only", start: true, wantCalls: []string{"proxy", "original"}},
		{name: "before only", before: true, wantCalls: []string{"proxy"}},
		{name: "both callbacks", before: true, start: true, wantCalls: []string{"proxy", "original"}},
		{name: "original error", start: true, originalError: callbackError, wantCalls: []string{"proxy", "original"}, wantError: callbackError},
		{name: "proxy error", start: true, proxyError: callbackError, wantCalls: []string{"proxy"}, wantError: callbackError},
	} {
		t.Run(test.name, func(t *testing.T) {
			model := internal.NewPerson()
			writer := &startCallbackWriter{}
			var calls []string
			if test.before {
				writer.before = func(serialization.Parsable) error { return nil }
			}
			if test.start {
				writer.start = func(item serialization.Parsable, actualWriter serialization.SerializationWriter) error {
					assert.Same(t, model, item)
					assert.Same(t, writer, actualWriter)
					calls = append(calls, "original")
					return test.originalError
				}
			}
			noop := func(serialization.Parsable) error { return nil }
			factory := serialization.NewSerializationWriterProxyFactory(
				&internal.MockSerializerFactory{SerializationWriter: writer}, noop, noop,
				func(item serialization.Parsable, actualWriter serialization.SerializationWriter) error {
					assert.Same(t, model, item)
					assert.Same(t, writer, actualWriter)
					calls = append(calls, "proxy")
					return test.proxyError
				},
			)
			result, err := factory.GetSerializationWriter("application/json")
			require.NoError(t, err)
			require.NotNil(t, result.GetOnStartObjectSerialization())
			require.NotPanics(t, func() {
				err = result.GetOnStartObjectSerialization()(model, result)
			})
			if test.wantError != nil {
				assert.ErrorIs(t, err, test.wantError)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.wantCalls, calls)
		})
	}
}

type startCallbackWriter struct {
	internal.MockSerializer
	before serialization.ParsableAction
	start  serialization.ParsableWriter
}

func (w *startCallbackWriter) GetOnBeforeSerialization() serialization.ParsableAction {
	return w.before
}

func (w *startCallbackWriter) GetOnStartObjectSerialization() serialization.ParsableWriter {
	return w.start
}

func (w *startCallbackWriter) SetOnStartObjectSerialization(action serialization.ParsableWriter) error {
	w.start = action
	return nil
}

type mockWriterWithAllCallbacks struct {
	internal.MockSerializer
	before serialization.ParsableAction
	after  serialization.ParsableAction
	start  serialization.ParsableWriter
}

func (w *mockWriterWithAllCallbacks) GetOnBeforeSerialization() serialization.ParsableAction {
	return w.before
}

func (w *mockWriterWithAllCallbacks) SetOnBeforeSerialization(action serialization.ParsableAction) error {
	w.before = action
	return nil
}

func (w *mockWriterWithAllCallbacks) GetOnAfterObjectSerialization() serialization.ParsableAction {
	return w.after
}

func (w *mockWriterWithAllCallbacks) SetOnAfterObjectSerialization(action serialization.ParsableAction) error {
	w.after = action
	return nil
}

func (w *mockWriterWithAllCallbacks) GetOnStartObjectSerialization() serialization.ParsableWriter {
	return w.start
}

func (w *mockWriterWithAllCallbacks) SetOnStartObjectSerialization(action serialization.ParsableWriter) error {
	w.start = action
	return nil
}

type mockWriterFactory struct {
	writer serialization.SerializationWriter
	err    error
}

func (m *mockWriterFactory) GetValidContentType() (string, error) {
	return "application/json", nil
}

func (m *mockWriterFactory) GetSerializationWriter(contentType string) (serialization.SerializationWriter, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.writer, nil
}

func TestSerializationWriterProxyNilGuards(t *testing.T) {
	var nilFactory *serialization.SerializationWriterProxyFactory

	ct, err := nilFactory.GetValidContentType()
	assert.Error(t, err)
	assert.Empty(t, ct)

	writer, err := nilFactory.GetSerializationWriter("application/json")
	assert.Error(t, err)
	assert.Nil(t, writer)

	emptyFactory := serialization.NewSerializationWriterProxyFactory(nil, nil, nil, nil)

	ct, err = emptyFactory.GetValidContentType()
	assert.Error(t, err)
	assert.Empty(t, ct)

	writer, err = emptyFactory.GetSerializationWriter("application/json")
	assert.Error(t, err)
	assert.Nil(t, writer)

	nilWriterFactory := &mockWriterFactory{writer: nil}
	proxyWithNilWriter := serialization.NewSerializationWriterProxyFactory(nilWriterFactory, nil, nil, nil)
	writer, err = proxyWithNilWriter.GetSerializationWriter("application/json")
	assert.Error(t, err)
	assert.Nil(t, writer)
}

func TestSerializationWriterProxyNilCallbacks(t *testing.T) {
	writer := &mockWriterWithAllCallbacks{}
	factory := serialization.NewSerializationWriterProxyFactory(
		&internal.MockSerializerFactory{SerializationWriter: writer},
		nil,
		nil,
		nil,
	)

	result, err := factory.GetSerializationWriter("application/json")
	require.NoError(t, err)
	require.NotNil(t, result)

	model := internal.NewPerson()

	assert.NotPanics(t, func() {
		before := result.GetOnBeforeSerialization()
		if before != nil {
			err := before(model)
			assert.NoError(t, err)
		}
	})

	assert.NotPanics(t, func() {
		after := result.GetOnAfterObjectSerialization()
		if after != nil {
			err := after(model)
			assert.NoError(t, err)
		}
	})

	assert.NotPanics(t, func() {
		start := result.GetOnStartObjectSerialization()
		if start != nil {
			err := start(model, result)
			assert.NoError(t, err)
		}
	})
}
