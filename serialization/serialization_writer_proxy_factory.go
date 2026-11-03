package serialization

import "errors"

// ParsableAction Encapsulates a method with a single Parsable parameter
type ParsableAction func(Parsable) error

// ParsableWriter  Encapsulates a method that receives a Parsable and SerializationWriter as parameters
type ParsableWriter func(Parsable, SerializationWriter) error

// SerializationWriterProxyFactory factory that allows the composition of before and after callbacks on existing factories.
type SerializationWriterProxyFactory struct {
	factory              SerializationWriterFactory
	onBeforeAction       ParsableAction
	onAfterAction        ParsableAction
	onSerializationStart ParsableWriter
}

// NewSerializationWriterProxyFactory constructs a new instance of SerializationWriterProxyFactory
func NewSerializationWriterProxyFactory(
	factory SerializationWriterFactory,
	onBeforeAction ParsableAction,
	onAfterAction ParsableAction,
	onSerializationStart ParsableWriter,
) *SerializationWriterProxyFactory {
	return &SerializationWriterProxyFactory{
		factory:              factory,
		onBeforeAction:       onBeforeAction,
		onAfterAction:        onAfterAction,
		onSerializationStart: onSerializationStart,
	}
}

func (s *SerializationWriterProxyFactory) GetValidContentType() (string, error) {
	if s == nil || s.factory == nil {
		return "", errors.New("factory is nil")
	}
	return s.factory.GetValidContentType()
}

func (s *SerializationWriterProxyFactory) GetSerializationWriter(contentType string) (SerializationWriter, error) {
	if s == nil || s.factory == nil {
		return nil, errors.New("factory is nil")
	}
	writer, err := s.factory.GetSerializationWriter(contentType)
	if err != nil {
		return nil, err
	}
	if writer == nil {
		return nil, errors.New("writer is nil")
	}

	originalBefore := writer.GetOnBeforeSerialization()
	err = writer.SetOnBeforeSerialization(func(parsable Parsable) error {
		if s != nil && s.onBeforeAction != nil {
			err := s.onBeforeAction(parsable)
			if err != nil {
				return err
			}
		}
		if originalBefore != nil {
			err := originalBefore(parsable)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	originalAfter := writer.GetOnAfterObjectSerialization()
	err = writer.SetOnAfterObjectSerialization(func(parsable Parsable) error {
		if s != nil && s.onAfterAction != nil {
			err := s.onAfterAction(parsable)
			if err != nil {
				return err
			}
		}
		if originalAfter != nil {
			err := originalAfter(parsable)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	originalStart := writer.GetOnStartObjectSerialization()
	err = writer.SetOnStartObjectSerialization(func(parsable Parsable, writer SerializationWriter) error {
		if s != nil && s.onSerializationStart != nil {
			err := s.onSerializationStart(parsable, writer)
			if err != nil {
				return err
			}
		}
		if originalStart != nil {
			err := originalStart(parsable, writer)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return writer, nil
}
