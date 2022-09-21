package internal

import "github.com/microsoft/kiota-abstractions-go/serialization"

type Person struct {
	displayName    *string
	callRecord     *CallRecord
	callRecords    []*CallRecord
	status         *PersonStatus
	previousStatus []*PersonStatus
	cardNumbers    []int
}

func NewPerson() *Person {
	return &Person{}
}

// Entity
type Entity struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]interface{}
	// Read-only.
	id *string
}

type CallRecord struct {
	Entity
}

func (c *CallRecord) Serialize(writer serialization.SerializationWriter) error {
	panic("implement me")
}

func (c *CallRecord) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	panic("implement me")
}

func NewCallRecord() *CallRecord {
	return &CallRecord{}
}
func (u *Person) SetDisplayName(name *string) {
	u.displayName = name
}

func (u *Person) GetDisplayName() *string {
	return u.displayName
}

func (u *Person) SetCallRecord(record *CallRecord) {
	u.callRecord = record
}

func (u *Person) GetCallRecord() *CallRecord {
	return u.callRecord
}

func (u *Person) SetCallRecords(records []*CallRecord) {
	u.callRecords = records
}

func (u *Person) GetCallRecords() []*CallRecord {
	return u.callRecords
}

func (u *Person) SetStatus(personStatus *PersonStatus) {
	u.status = personStatus
}

func (u *Person) GetStatus() *PersonStatus {
	return u.status
}

func (u *Person) SetPreviousStatus(previousStatus []*PersonStatus) {
	u.previousStatus = previousStatus
}

func (u *Person) GetPreviousStatus() []*PersonStatus {
	return u.previousStatus
}

func (u *Person) SetCardNumbers(numbers []int) {
	u.cardNumbers = numbers
}

func (u *Person) GetCardNumbers() []int {
	return u.cardNumbers
}
