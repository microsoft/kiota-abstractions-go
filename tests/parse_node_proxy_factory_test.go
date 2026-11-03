package tests

import (
	"errors"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockParseNodeWithCallbacks struct {
	internal.MockParseNode
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

type mockParseNodeFactoryWithNode struct {
	node serialization.ParseNode
	err  error
}

func (m *mockParseNodeFactoryWithNode) GetValidContentType() (string, error) {
	return "application/json", nil
}

func (m *mockParseNodeFactoryWithNode) GetRootParseNode(contentType string, content []byte) (serialization.ParseNode, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.node, nil
}

func TestParseNodeProxyFactory_InvokesBothCallbacks(t *testing.T) {
	var calls []string
	node := &mockParseNodeWithCallbacks{
		before: func(p serialization.Parsable) error {
			calls = append(calls, "originalBefore")
			return nil
		},
		after: func(p serialization.Parsable) error {
			calls = append(calls, "originalAfter")
			return nil
		},
	}
	mockFactory := &mockParseNodeFactoryWithNode{node: node}
	proxyFactory := serialization.NewParseNodeProxyFactory(
		mockFactory,
		func(p serialization.Parsable) error {
			calls = append(calls, "proxyBefore")
			return nil
		},
		func(p serialization.Parsable) error {
			calls = append(calls, "proxyAfter")
			return nil
		},
	)

	resultNode, err := proxyFactory.GetRootParseNode("application/json", []byte{})
	require.NoError(t, err)
	require.NotNil(t, resultNode)

	model := internal.NewPerson()

	require.NotNil(t, resultNode.GetOnBeforeAssignFieldValues())
	err = resultNode.GetOnBeforeAssignFieldValues()(model)
	require.NoError(t, err)
	assert.Equal(t, []string{"proxyBefore", "originalBefore"}, calls)

	require.NotNil(t, resultNode.GetOnAfterAssignFieldValues())
	err = resultNode.GetOnAfterAssignFieldValues()(model)
	require.NoError(t, err)
	assert.Equal(t, []string{"proxyBefore", "originalBefore", "proxyAfter", "originalAfter"}, calls)
}

func TestParseNodeProxyFactory_NilCallbacks(t *testing.T) {
	node := &mockParseNodeWithCallbacks{}
	mockFactory := &mockParseNodeFactoryWithNode{node: node}
	proxyFactory := serialization.NewParseNodeProxyFactory(mockFactory, nil, nil)

	resultNode, err := proxyFactory.GetRootParseNode("application/json", []byte{})
	require.NoError(t, err)

	model := internal.NewPerson()
	assert.NotPanics(t, func() {
		err := resultNode.GetOnBeforeAssignFieldValues()(model)
		assert.NoError(t, err)
	})
	assert.NotPanics(t, func() {
		err := resultNode.GetOnAfterAssignFieldValues()(model)
		assert.NoError(t, err)
	})
}

func TestParseNodeProxyFactory_ErrorPropagation(t *testing.T) {
	expectedErr := errors.New("callback error")

	// Error in proxy before
	node := &mockParseNodeWithCallbacks{}
	factory := serialization.NewParseNodeProxyFactory(
		&mockParseNodeFactoryWithNode{node: node},
		func(p serialization.Parsable) error { return expectedErr },
		nil,
	)
	resultNode, err := factory.GetRootParseNode("application/json", []byte{})
	require.NoError(t, err)
	assert.ErrorIs(t, resultNode.GetOnBeforeAssignFieldValues()(internal.NewPerson()), expectedErr)

	// Error in proxy after
	node2 := &mockParseNodeWithCallbacks{}
	factory2 := serialization.NewParseNodeProxyFactory(
		&mockParseNodeFactoryWithNode{node: node2},
		nil,
		func(p serialization.Parsable) error { return expectedErr },
	)
	resultNode2, err := factory2.GetRootParseNode("application/json", []byte{})
	require.NoError(t, err)
	assert.ErrorIs(t, resultNode2.GetOnAfterAssignFieldValues()(internal.NewPerson()), expectedErr)
}

func TestParseNodeProxyFactory_NilGuards(t *testing.T) {
	var nilFactory *serialization.ParseNodeProxyFactory

	ct, err := nilFactory.GetValidContentType()
	assert.Error(t, err)
	assert.Empty(t, ct)

	node, err := nilFactory.GetRootParseNode("application/json", []byte{})
	assert.Error(t, err)
	assert.Nil(t, node)

	emptyFactory := serialization.NewParseNodeProxyFactory(nil, nil, nil)

	ct, err = emptyFactory.GetValidContentType()
	assert.Error(t, err)
	assert.Empty(t, ct)

	node, err = emptyFactory.GetRootParseNode("application/json", []byte{})
	assert.Error(t, err)
	assert.Nil(t, node)

	nilNodeFactory := &mockParseNodeFactoryWithNode{node: nil}
	proxyWithNilNode := serialization.NewParseNodeProxyFactory(nilNodeFactory, nil, nil)
	node, err = proxyWithNilNode.GetRootParseNode("application/json", []byte{})
	assert.Error(t, err)
	assert.Nil(t, node)
}
