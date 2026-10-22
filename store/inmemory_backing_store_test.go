package store

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
)

func TestSetAndGetsValuesFromStore(t *testing.T) {
	memoryStore := NewInMemoryBackingStore()
	assert.Equal(t, 0, len(memoryStore.Enumerate()))

	err := memoryStore.Set("name", "Michael")
	assert.Nil(t, err)

	assert.Equal(t, 1, len(memoryStore.Enumerate()))
	assert.Equal(t, memoryStore.Enumerate()["name"], "Michael")
}

func TestPreventsDuplicateKeysInStore(t *testing.T) {
	memoryStore := NewInMemoryBackingStore()
	assert.Equal(t, 0, len(memoryStore.Enumerate()))

	err := memoryStore.Set("name", "Michael")
	err = memoryStore.Set("name", "Jordan")
	assert.Nil(t, err)

	assert.Equal(t, 1, len(memoryStore.Enumerate()))
	assert.Equal(t, memoryStore.Enumerate()["name"], "Jordan")
}

func TestEnumeratesValuesChangedToNullInStore(t *testing.T) {
	memoryStore := NewInMemoryBackingStore()
	assert.Equal(t, 0, len(memoryStore.Enumerate()))

	err := memoryStore.Set("name", "Michael")
	assert.Nil(t, err)

	err = memoryStore.Set("email", "michael.jordan@live.com")
	assert.Nil(t, err)

	err = memoryStore.Set("phone", nil)
	assert.Nil(t, err)

	assert.Equal(t, 3, len(memoryStore.Enumerate()))
	assert.Equal(t, memoryStore.EnumerateKeysForValuesChangedToNil(), []string{"phone"})
}

func TestEnumeratesTypedNilCollections(t *testing.T) {
	value := "value"
	for _, test := range []struct {
		name       string
		previous   interface{}
		value      interface{}
		changedNil bool
	}{
		{"nil slice", []string{"old"}, []string(nil), true},
		{"nil object slice", []*testEntity{NewTestEntity()}, []*testEntity(nil), true},
		{"nil byte slice", []byte{1}, []byte(nil), true},
		{"nil map", map[string]string{"old": "value"}, map[string]string(nil), true},
		{"nil pointer", &value, (*string)(nil), true},
		{"nil interface", "old", nil, true},
		{"empty slice", []string{"old"}, []string{}, false},
		{"empty map", map[string]string{"old": "value"}, map[string]string{}, false},
		{"populated slice", []string{"old"}, []string{"new"}, false},
		{"populated map", map[string]string{"old": "value"}, map[string]string{"new": "value"}, false},
		{"unchanged nil slice", []string(nil), []string(nil), false},
		{"unchanged nil map", map[string]string(nil), map[string]string(nil), false},
		{"zero integer", 1, 0, false},
		{"false boolean", true, false, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			memoryStore := NewInMemoryBackingStore()
			memoryStore.SetInitializationCompleted(false)
			assert.NoError(t, memoryStore.Set("value", test.previous))
			memoryStore.SetInitializationCompleted(true)
			assert.NoError(t, memoryStore.Set("value", test.value))

			keys := memoryStore.EnumerateKeysForValuesChangedToNil()
			if test.changedNil {
				assert.Equal(t, []string{"value"}, keys)
			} else {
				assert.Empty(t, keys)
			}
		})
	}
}

func TestBackingStoreEmbeddedInModel(t *testing.T) {
	testUser := NewTestEntity()
	testUser.GetBackingStore().SetInitializationCompleted(false)
	id := "1234"
	testUser.SetId(&id)
	testUser.GetBackingStore().SetInitializationCompleted(true)

	testUser.SetPhoneNumbers([]string{"+1234", "+2345"})
	name := "Jeane"
	testUser.SetName(&name)

	testUser.GetBackingStore().SetReturnOnlyChangedValues(false)
	assert.Equal(t, "1234", *testUser.GetId())

	testUser.GetBackingStore().SetReturnOnlyChangedValues(true)
	assert.Nil(t, testUser.GetId())

	changes := testUser.GetBackingStore().Enumerate()

	assert.Equal(t, 2, len(changes))
	assert.Equal(t, []string{"+1234", "+2345"}, changes["phoneNumbers"])
	assert.Equal(t, "Jeane", *changes["name"].(*string))

	assert.Equal(t, "Jeane", *testUser.GetName())
}

func TestSubscribe(t *testing.T) {
	testUser := NewTestEntity()

	calls := 0
	var keys []string
	var oldValues []interface{}
	var newValues []interface{}
	subscriber := func(key string, oldVal interface{}, newVal interface{}) {
		calls++
		keys = append(keys, key)

		if oldVal != nil {
			oldValues = append(oldValues, *oldVal.(*string))
		} else {
			oldValues = append(oldValues, nil)
		}

		if newVal != nil {
			newValues = append(newValues, *newVal.(*string))
		} else {
			newValues = append(newValues, nil)
		}
	}
	testUser.GetBackingStore().Subscribe(subscriber)

	id := "1234"
	testUser.SetId(&id)
	x := "11"
	testUser.SetId(&x)

	testUser.SetId(&id)

	assert.Equal(t, 3, calls)
	assert.Equal(t, []string{"id", "id", "id"}, keys)
	assert.Equal(t, []interface{}{nil, "1234", "11"}, oldValues)
	assert.Equal(t, []interface{}{"1234", "11", "1234"}, newValues)
}

func TestSubscribeUnSubscribe(t *testing.T) {

	testUser := NewTestEntity()
	calls := 0
	subscriber := func(key string, oldVal interface{}, newVal interface{}) {
		calls++
	}
	subscriberId := (uuid.New()).String()
	err := testUser.GetBackingStore().SubscribeWithId(subscriber, subscriberId)
	assert.Nil(t, err)

	id := "1"
	testUser.SetId(&id)
	assert.Equal(t, 1, calls)

	id2 := "2"
	testUser.SetId(&id2)
	assert.Equal(t, 2, calls)

	err = testUser.GetBackingStore().Unsubscribe(subscriberId)
	assert.Nil(t, err)

	id3 := "3"
	testUser.SetId(&id3)
	assert.Equal(t, 2, calls)
}

func TestClear(t *testing.T) {

	testUser := NewTestEntity()

	id := "1"
	testUser.SetId(&id)

	assert.Equal(t, "1", *testUser.GetId())
	testUser.GetBackingStore().Clear()

	assert.Nil(t, testUser.GetId())
}

func TestReplaceSlice(t *testing.T) {
	memoryStore := NewInMemoryBackingStore()
	assert.Equal(t, 0, len(memoryStore.Enumerate()))

	err := memoryStore.Set("key", []string{"a", "b"})
	assert.Nil(t, err)

	err = memoryStore.Set("key", []string{"b", "c"})
	assert.Nil(t, err)

	assert.Equal(t, 1, len(memoryStore.Enumerate()))
	val := memoryStore.Enumerate()["key"].([]string)
	assert.Equal(t, 2, len(val))
	assert.Equal(t, "b", val[0])
	assert.Equal(t, "c", val[1])
}

func TestReplaceMap(t *testing.T) {
	memoryStore := NewInMemoryBackingStore()
	assert.Equal(t, 0, len(memoryStore.Enumerate()))

	err := memoryStore.Set("key", map[string]string{"k": "v1"})
	assert.Nil(t, err)

	err = memoryStore.Set("key", map[string]string{"k": "v2"})
	assert.Nil(t, err)

	assert.Equal(t, 1, len(memoryStore.Enumerate()))
	val := memoryStore.Enumerate()["key"].(map[string]string)
	assert.Equal(t, 1, len(val))
	assert.Equal(t, "v2", val["k"])
}

func TestReplaceStruct(t *testing.T) {
	memoryStore := NewInMemoryBackingStore()
	assert.Equal(t, 0, len(memoryStore.Enumerate()))

	prev := struct{ slice []string }{
		slice: []string{"a", "b"},
	}
	err := memoryStore.Set("key", prev)
	assert.Nil(t, err)

	curr := struct{ slice []string }{
		slice: []string{"b", "c"},
	}
	err = memoryStore.Set("key", curr)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(memoryStore.Enumerate()))
	assert.True(t, reflect.DeepEqual(curr, memoryStore.Enumerate()["key"]))
}

func TestItConsidersSettingNullAsAChange(t *testing.T) {
	testEntityInstance := NewTestEntity()
	testEntityInstance.GetBackingStore().SetInitializationCompleted(true)
	testEntityInstance.GetBackingStore().SetReturnOnlyChangedValues(true)
	testEntityInstance.SetId(nil)
	assert.Equal(t, 1, len(testEntityInstance.GetBackingStore().Enumerate()))
	assert.Equal(t, []string{"id"}, testEntityInstance.GetBackingStore().EnumerateKeysForValuesChangedToNil())
}

type testEntity struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]interface{}
	id             *string
	name           *string
	phoneNumbers   []string
	items          []*testEntity
	backingStore   BackingStore
}

func (t *testEntity) GetAdditionalData() map[string]interface{} {
	return t.additionalData
}

func (t *testEntity) SetAdditionalData(value map[string]interface{}) {
	t.additionalData = value
}

func (t *testEntity) GetBackingStore() BackingStore {
	return t.backingStore
}

func (t *testEntity) Serialize(writer serialization.SerializationWriter) error {
	panic("implement me")
}

func (t *testEntity) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	panic("implement me")
}

func (t *testEntity) GetId() *string {
	val, _ := t.GetBackingStore().Get("id")
	if val != nil {
		return val.(*string)
	}
	return nil
}

func (t *testEntity) SetId(id *string) {
	err := t.GetBackingStore().Set("id", id)
	if err != nil {
		panic(err)
	}
}

func (t *testEntity) GetName() *string {
	val, _ := t.GetBackingStore().Get("name")
	if val != nil {
		return val.(*string)
	}
	return nil
}

func (t *testEntity) SetName(name *string) {
	err := t.GetBackingStore().Set("name", name)
	if err != nil {
		panic(err)
	}
}

func (t *testEntity) GetPhoneNumbers() []string {
	val, _ := t.GetBackingStore().Get("phoneNumbers")
	if val != nil {
		return val.([]string)
	}
	return nil
}

func (t *testEntity) SetPhoneNumbers(numbers []string) {
	err := t.GetBackingStore().Set("phoneNumbers", numbers)
	if err != nil {
		panic(err)
	}
}

func NewTestEntity() *testEntity {
	return &testEntity{
		backingStore: BackingStoreFactoryInstance(),
	}
}

func TestInMemoryBackingStoreNilReceiverSafety(t *testing.T) {
	var nilStore *InMemoryBackingStore

	assert.NotPanics(t, func() {
		val, err := nilStore.Get("test")
		assert.Nil(t, val)
		assert.NotNil(t, err)

		err = nilStore.Set("test", "val")
		assert.NotNil(t, err)

		enum := nilStore.Enumerate()
		assert.Empty(t, enum)

		keys := nilStore.EnumerateKeysForValuesChangedToNil()
		assert.Empty(t, keys)

		subId := nilStore.Subscribe(func(key string, oldVal interface{}, newVal interface{}) {})
		assert.Empty(t, subId)

		err = nilStore.SubscribeWithId(func(key string, oldVal interface{}, newVal interface{}) {}, "sub")
		assert.NotNil(t, err)

		err = nilStore.Unsubscribe("sub")
		assert.Nil(t, err)

		nilStore.Clear()
		assert.False(t, nilStore.GetInitializationCompleted())
		nilStore.SetInitializationCompleted(true)
		assert.False(t, nilStore.GetReturnOnlyChangedValues())
		nilStore.SetReturnOnlyChangedValues(true)
	})
}
