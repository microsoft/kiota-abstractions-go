package abstractions

import (
	"errors"
	"github.com/microsoft/kiota-abstractions-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSetValueWithoutError(t *testing.T) {

	person := internal.NewPerson()
	callFactory := func() (*internal.CallRecord, error) {
		return internal.NewCallRecord(), nil
	}

	err := SetValue(callFactory, person.SetCallRecord)
	assert.Nil(t, err)
	assert.NotNil(t, person.GetCallRecord())
}

func TestSetValueWithError(t *testing.T) {

	person := internal.NewPerson()
	callFactory := func() (*internal.CallRecord, error) {
		return nil, errors.New("could not get from factory")
	}

	err := SetValue(callFactory, person.SetCallRecord)
	assert.NotNil(t, err)
	assert.Nil(t, person.GetCallRecord())
}

func createCallRecordNode(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewCallRecord(), nil
}

func getObjectsValues(ctor serialization.ParsableFactory) ([]serialization.Parsable, error) {
	slice := []serialization.Parsable{internal.NewCallRecord(), internal.NewCallRecord(), internal.NewCallRecord()}
	return slice, nil
}

func getObjectsValuesWithError(ctor serialization.ParsableFactory) ([]serialization.Parsable, error) {
	return nil, errors.New("could not get from factory")
}

func TestSetCollectionValueValueWithoutError(t *testing.T) {

	person := internal.NewPerson()
	err := SetCollectionValue(getObjectsValues, createCallRecordNode, person.SetCallRecords)
	assert.Nil(t, err)
	assert.NotNil(t, person.GetCallRecords())
	assert.Equal(t, len(person.GetCallRecords()), 3)
}

func TestSetCollectionValueValueWithError(t *testing.T) {

	person := internal.NewPerson()
	err := SetCollectionValue(getObjectsValuesWithError, createCallRecordNode, person.SetCallRecords)
	assert.NotNil(t, err)
	assert.Nil(t, person.GetCallRecords())
	assert.Equal(t, len(person.GetCallRecords()), 0)
}

func TestCollectionApply(t *testing.T) {

	slice := []string{"1", "2", "3"}
	response := CollectionApply(slice, func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})

	assert.Equal(t, len(response), 3)
	assert.Equal(t, response, []int{1, 2, 3})
}

func TestCollectionCast1(t *testing.T) {
	slice := []*int{p(1), p(2), p(3)}
	response := CollectionValueCast[int](slice)

	assert.Equal(t, len(response), 3)
	assert.Equal(t, response, []int{1, 2, 3})
}

func TestCollectionCast(t *testing.T) {

	var val1 interface{}
	val1 = 1
	var val2 interface{}
	val2 = 2
	var val3 interface{}
	val3 = 3

	slice := []interface{}{val1, val2, val3}
	response := CollectionCast[int](slice)

	assert.Equal(t, len(response), 3)
	assert.Equal(t, response, []int{1, 2, 3})
}

func TestCollectionCastAsParsable(t *testing.T) {
	slice := []interface{}{&internal.CallRecord{}, &internal.CallRecord{}, &internal.CallRecord{}}
	response := CollectionCast[serialization.Parsable](slice)

	assert.Equal(t, len(response), 3)
}

func getEnumValue(parser serialization.EnumFactory) (interface{}, error) {
	status := internal.ACTIVE
	return &status, nil
}

func getEnumValueWithError(parser serialization.EnumFactory) (interface{}, error) {
	return nil, errors.New("could not get from factory")
}

func TestSetReferencedEnumValueValueValueWithoutError(t *testing.T) {
	person := internal.NewPerson()
	err := SetReferencedEnumValue(getEnumValue, internal.ParsePersonStatus, person.SetStatus)
	assert.Nil(t, err)
	assert.Equal(t, person.GetStatus().String(), internal.ACTIVE.String())
}

func TestSetReferencedEnumValueValueValueWithError(t *testing.T) {
	person := internal.NewPerson()
	err := SetReferencedEnumValue(getEnumValueWithError, internal.ParsePersonStatus, person.SetStatus)
	assert.NotNil(t, err)
	assert.Nil(t, person.GetStatus())
}

func TestSetCollectionOfReferencedEnumValueWithoutError(t *testing.T) {
	person := internal.NewPerson()

	slice := []interface{}{p(internal.ACTIVE), p(internal.SUSPEND)}
	enumSource := func(parser serialization.EnumFactory) ([]interface{}, error) {
		return slice, nil
	}

	err := SetCollectionOfReferencedEnumValue(enumSource, internal.ParsePersonStatus, person.SetPreviousStatus)
	assert.Nil(t, err)
	assert.Equal(t, person.GetPreviousStatus()[0].String(), internal.ACTIVE.String())
	assert.Equal(t, person.GetPreviousStatus()[1].String(), internal.SUSPEND.String())
}

func TestSetCollectionOfReferencedEnumValueWithError(t *testing.T) {
	person := internal.NewPerson()

	enumSource := func(parser serialization.EnumFactory) ([]interface{}, error) {
		return nil, errors.New("could not get from factory")
	}

	err := SetCollectionOfReferencedEnumValue(enumSource, internal.ParsePersonStatus, person.SetPreviousStatus)
	assert.NotNil(t, err)
	assert.Nil(t, person.GetPreviousStatus())
}

func TestSetCollectionOfReferencedPrimitiveValueWithoutError(t *testing.T) {
	person := internal.NewPerson()

	slice := []interface{}{p(1), p(2), p(3)}
	dataSource := func(targetType string) ([]interface{}, error) {
		return slice, nil
	}

	err := SetCollectionOfReferencedPrimitiveValue(dataSource, "int", person.SetCardNumbers)
	assert.Nil(t, err)
	assert.Equal(t, person.GetCardNumbers()[0], 1)
	assert.Equal(t, person.GetCardNumbers()[1], 2)
	assert.Equal(t, person.GetCardNumbers()[2], 3)
}

func TestSetCollectionOfReferencedPrimitiveValueWithError(t *testing.T) {
	person := internal.NewPerson()

	dataSource := func(targetType string) ([]interface{}, error) {
		return nil, errors.New("could not get from factory")
	}

	err := SetCollectionOfReferencedPrimitiveValue(dataSource, "int", person.SetCardNumbers)
	assert.NotNil(t, err)
	assert.Nil(t, person.GetCardNumbers())
}

func TestSetCollectionOfPrimitiveValueWithoutError(t *testing.T) {
	person := internal.NewPerson()

	slice := []interface{}{1, 2, 3}
	dataSource := func(targetType string) ([]interface{}, error) {
		return slice, nil
	}

	err := SetCollectionOfPrimitiveValue(dataSource, "int", person.SetCardNumbers)
	assert.Nil(t, err)
	assert.Equal(t, person.GetCardNumbers()[0], 1)
	assert.Equal(t, person.GetCardNumbers()[1], 2)
	assert.Equal(t, person.GetCardNumbers()[1], 2)
	assert.Equal(t, person.GetCardNumbers()[2], 3)
}

func TestSetCollectionOfPrimitiveValueWithError(t *testing.T) {
	person := internal.NewPerson()

	dataSource := func(targetType string) ([]interface{}, error) {
		return nil, errors.New("could not get from factory")
	}

	err := SetCollectionOfPrimitiveValue(dataSource, "int", person.SetCardNumbers)
	assert.NotNil(t, err)
	assert.Nil(t, person.GetCardNumbers())
}

func TestGetValueReturn(t *testing.T) {
	person := internal.NewPerson()

	assert.Equal(t, GetValueOrDefault(person.GetDisplayName, "Unknown"), "Unknown")

	person.SetDisplayName(p("Jane"))
	assert.Equal(t, GetValueOrDefault(person.GetDisplayName, "Unknown"), "Jane")
}

type foo struct {
	Name string
}

type animal interface {
	Eat()
	Call() string
}

type dog interface {
	Bark()
	Call() string
}

func (f *foo) Bark()        {}
func (f *foo) Eat()         {}
func (f *foo) Call() string { return f.Name }

func TestCollectionStructCast(t *testing.T) {

	foos := []foo{{Name: "Cooper"}, {Name: "Buddy"}, {Name: "Peanut"}}
	animals := CollectionStructCast[animal](foos)
	assert.Equal(t, "Cooper", animals[0].Call())
	assert.Equal(t, "Buddy", animals[1].Call())
	assert.Equal(t, "Peanut", animals[2].Call())
}

func TestCopyMap(t *testing.T) {
	source := map[string]int{"foo": 1, "bar": 2}
	duplicate := CopyMap(source)

	assert.Equal(t, duplicate, source)
}

func TestStringCopyMap(t *testing.T) {
	source := map[string]string{"foo": "1", "bar": "2"}
	duplicate := CopyStringMap(source)

	assert.Equal(t, duplicate, source)
}
