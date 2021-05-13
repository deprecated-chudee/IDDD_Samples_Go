package common

import (
	"errors"
	"strings"
)

type AssertionConcern struct {}

func (AssertionConcern) AssertArgumentEquals(struct1, struct2 interface{}, message string) error {
	if &struct1 != &struct2 {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertArgumentNotEquals(struct1, struct2 interface{}, message string) error {
	if &struct1 == &struct2 {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertFalse(boolean bool, message string) error {
	if boolean {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertTrue(boolean bool, message string) error {
	if !boolean {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertArgumentMaxLength(str string, maximum int, message string) error {
	length := len(strings.TrimSpace(str))
	if length > maximum {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertArgumentMinMaxLength(str string, minimum, maximum int, message string) error {
	length := len(strings.TrimSpace(str))
	if length < minimum || length > maximum {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertArgumentNotEmpty(str, message string) error {
	if strings.TrimSpace(str) == "" {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertArgumentNotNil(object interface{}, message string) error {
	if object == nil {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertArgumentRangeFloat64(value, min, max float64, message string) error {
	if value < min || value > max {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertArgumentRangeFloat32(value, min, max float32, message string) error {
	if value < min || value > max {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertArgumentRangeInt64(value, min, max int64, message string) error {
	if value < min || value > max {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertArgumentRangeInt32(value, min, max int32, message string) error {
	if value < min || value > max {
		return errors.New(message)
	}

	return nil
}

func (AssertionConcern) AssertArgumentRangeInt(value, min, max int, message string) error {
	if value < min || value > max {
		return errors.New(message)
	}

	return nil
}