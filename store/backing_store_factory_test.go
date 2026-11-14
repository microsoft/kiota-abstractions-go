package store

import (
	"testing"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDefaultBackingStoreInstance(t *testing.T) {
	assert.NotNil(t, BackingStoreFactoryInstance)
}

type testBackedModel struct {
	store BackingStore
}

func newTestBackedModel() *testBackedModel {
	return &testBackedModel{store: NewInMemoryBackingStore()}
}

func (m *testBackedModel) GetBackingStore() BackingStore {
	return m.store
}

func (m *testBackedModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func (m *testBackedModel) Serialize(writer serialization.SerializationWriter) error {
	return nil
}

type mockParseNodeWithCallbacks struct {
	serialization.ParseNode
	before serialization.ParsableAction
	after  serialization.ParsableAction
}

func (m *mockParseNodeWithCallbacks) GetOnBeforeAssignFieldValues() serialization.ParsableAction {
	return m.before
}

func (m *mockParseNodeWithCallbacks) SetOnBeforeAssignFieldValues(action serialization.ParsableAction) error {
	m.before = action
	return nil
}

func (m *mockParseNodeWithCallbacks) GetOnAfterAssignFieldValues() serialization.ParsableAction {
	return m.after
}

func (m *mockParseNodeWithCallbacks) SetOnAfterAssignFieldValues(action serialization.ParsableAction) error {
	m.after = action
	return nil
}

type mockParseNodeFactoryForStore struct {
	node serialization.ParseNode
}

func (m *mockParseNodeFactoryForStore) GetValidContentType() (string, error) {
	return "application/json", nil
}

func (m *mockParseNodeFactoryForStore) GetRootParseNode(contentType string, content []byte) (serialization.ParseNode, error) {
	return m.node, nil
}

type mockWriterWithCallbacksForStore struct {
	serialization.SerializationWriter
	before serialization.ParsableAction
	after  serialization.ParsableAction
	start  serialization.ParsableWriter
}

func (w *mockWriterWithCallbacksForStore) GetOnBeforeSerialization() serialization.ParsableAction {
	return w.before
}

func (w *mockWriterWithCallbacksForStore) SetOnBeforeSerialization(action serialization.ParsableAction) error {
	w.before = action
	return nil
}

func (w *mockWriterWithCallbacksForStore) GetOnAfterObjectSerialization() serialization.ParsableAction {
	return w.after
}

func (w *mockWriterWithCallbacksForStore) SetOnAfterObjectSerialization(action serialization.ParsableAction) error {
	w.after = action
	return nil
}

func (w *mockWriterWithCallbacksForStore) GetOnStartObjectSerialization() serialization.ParsableWriter {
	return w.start
}

func (w *mockWriterWithCallbacksForStore) SetOnStartObjectSerialization(action serialization.ParsableWriter) error {
	w.start = action
	return nil
}

type mockWriterFactoryForStore struct {
	writer serialization.SerializationWriter
}

func (m *mockWriterFactoryForStore) GetValidContentType() (string, error) {
	return "application/json", nil
}

func (m *mockWriterFactoryForStore) GetSerializationWriter(contentType string) (serialization.SerializationWriter, error) {
	return m.writer, nil
}

func TestBackingStoreParseNodeFactory_Lifecycle(t *testing.T) {
	node := &mockParseNodeWithCallbacks{}
	factory := NewBackingStoreParseNodeFactory(&mockParseNodeFactoryForStore{node: node})

	ct, err := factory.GetValidContentType()
	require.NoError(t, err)
	assert.Equal(t, "application/json", ct)

	rootNode, err := factory.GetRootParseNode("application/json", []byte{})
	require.NoError(t, err)
	require.NotNil(t, rootNode)

	model := newTestBackedModel()
	assert.True(t, model.GetBackingStore().GetInitializationCompleted())

	// Before assign field values: initialization completed should become false
	beforeAction := rootNode.GetOnBeforeAssignFieldValues()
	require.NotNil(t, beforeAction)
	err = beforeAction(model)
	require.NoError(t, err)
	assert.False(t, model.GetBackingStore().GetInitializationCompleted())

	// After assign field values: initialization completed MUST become true
	afterAction := rootNode.GetOnAfterAssignFieldValues()
	require.NotNil(t, afterAction)
	err = afterAction(model)
	require.NoError(t, err)
	assert.True(t, model.GetBackingStore().GetInitializationCompleted())
}

func TestBackingStoreSerializationWriterProxyFactory_Lifecycle(t *testing.T) {
	writer := &mockWriterWithCallbacksForStore{}
	factory := NewBackingStoreSerializationWriterProxyFactory(&mockWriterFactoryForStore{writer: writer})

	ct, err := factory.GetValidContentType()
	require.NoError(t, err)
	assert.Equal(t, "application/json", ct)

	writerRes, err := factory.GetSerializationWriter("application/json")
	require.NoError(t, err)
	require.NotNil(t, writerRes)

	model := newTestBackedModel()

	// Before serialization: return only changed values should become true
	beforeAction := writerRes.GetOnBeforeSerialization()
	require.NotNil(t, beforeAction)
	err = beforeAction(model)
	require.NoError(t, err)
	assert.True(t, model.GetBackingStore().GetReturnOnlyChangedValues())

	// After serialization: return only changed values should become false, and init completed true
	afterAction := writerRes.GetOnAfterObjectSerialization()
	require.NotNil(t, afterAction)
	err = afterAction(model)
	require.NoError(t, err)
	assert.False(t, model.GetBackingStore().GetReturnOnlyChangedValues())
	assert.True(t, model.GetBackingStore().GetInitializationCompleted())
}

func TestBackingStoreFactories_NilGuards(t *testing.T) {
	var nilParseFactory *BackingStoreParseNodeFactory
	ct, err := nilParseFactory.GetValidContentType()
	assert.Error(t, err)
	assert.Empty(t, ct)

	node, err := nilParseFactory.GetRootParseNode("application/json", []byte{})
	assert.Error(t, err)
	assert.Nil(t, node)

	emptyParseFactory := &BackingStoreParseNodeFactory{}
	ct, err = emptyParseFactory.GetValidContentType()
	assert.Error(t, err)
	assert.Empty(t, ct)

	node, err = emptyParseFactory.GetRootParseNode("application/json", []byte{})
	assert.Error(t, err)
	assert.Nil(t, node)

	var nilWriterFactory *BackingStoreSerializationWriterProxyFactory
	ct, err = nilWriterFactory.GetValidContentType()
	assert.Error(t, err)
	assert.Empty(t, ct)

	writer, err := nilWriterFactory.GetSerializationWriter("application/json")
	assert.Error(t, err)
	assert.Nil(t, writer)

	emptyWriterFactory := &BackingStoreSerializationWriterProxyFactory{}
	ct, err = emptyWriterFactory.GetValidContentType()
	assert.Error(t, err)
	assert.Empty(t, ct)

	writer, err = emptyWriterFactory.GetSerializationWriter("application/json")
	assert.Error(t, err)
	assert.Nil(t, writer)
}

type mockParsableWithDeserializers struct {
	testBackedModel
	fields map[string]func(serialization.ParseNode) error
}

func newMockParsableWithDeserializers(fields map[string]func(serialization.ParseNode) error) *mockParsableWithDeserializers {
	return &mockParsableWithDeserializers{
		testBackedModel: *newTestBackedModel(),
		fields:          fields,
	}
}

func (m *mockParsableWithDeserializers) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return m.fields
}

type mockRecordingSerializationWriter struct {
	mockWriterWithCallbacksForStore
	writtenNullKeys []string
}

func (m *mockRecordingSerializationWriter) WriteNullValue(key string) error {
	m.writtenNullKeys = append(m.writtenNullKeys, key)
	return nil
}

func TestBackingStoreSerializationWriterProxyFactory_KebabCaseResolution(t *testing.T) {
	writer := &mockRecordingSerializationWriter{}
	factory := NewBackingStoreSerializationWriterProxyFactory(&mockWriterFactoryForStore{writer: writer})

	writerRes, err := factory.GetSerializationWriter("application/json")
	require.NoError(t, err)
	require.NotNil(t, writerRes)

	model := newMockParsableWithDeserializers(map[string]func(serialization.ParseNode) error{
		"org-name": func(serialization.ParseNode) error { return nil },
	})
	model.GetBackingStore().SetInitializationCompleted(true)
	model.GetBackingStore().Set("orgName", nil)

	startAction := writerRes.GetOnStartObjectSerialization()
	require.NotNil(t, startAction)

	err = startAction(model, writer)
	require.NoError(t, err)

	assert.Equal(t, []string{"org-name"}, writer.writtenNullKeys)
}

func TestBackingStoreSerializationWriterProxyFactory_SnakeCaseResolution(t *testing.T) {
	writer := &mockRecordingSerializationWriter{}
	factory := NewBackingStoreSerializationWriterProxyFactory(&mockWriterFactoryForStore{writer: writer})

	writerRes, err := factory.GetSerializationWriter("application/json")
	require.NoError(t, err)
	require.NotNil(t, writerRes)

	model := newMockParsableWithDeserializers(map[string]func(serialization.ParseNode) error{
		"first_name":   func(serialization.ParseNode) error { return nil },
		"home_address": func(serialization.ParseNode) error { return nil },
	})
	model.GetBackingStore().SetInitializationCompleted(true)
	model.GetBackingStore().Set("firstName", nil)

	startAction := writerRes.GetOnStartObjectSerialization()
	require.NotNil(t, startAction)

	err = startAction(model, writer)
	require.NoError(t, err)

	assert.Equal(t, []string{"first_name"}, writer.writtenNullKeys)
}
