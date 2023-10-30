package utils

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

func JSONFromString[T any](value string, target *T) error {
	if value == "" {
		return errors.New("empty value")
	}
	return json.NewDecoder(strings.NewReader(value)).Decode(target)
}

func JSONFromStringOpt[T any](value string, target *T) error {
	if value == "" {
		return nil
	}
	return json.NewDecoder(strings.NewReader(value)).Decode(target)
}

func Atoi(input string, target *int) error {
	if input == "" {
		return errors.New("empty value")
	}
	value, err := strconv.Atoi(input)
	if err != nil {
		return err
	}
	*target = value
	return nil
}

func AtoiOpt(input string, target *int) error {
	if input == "" {
		*target = 0
		return nil
	}
	value, err := strconv.Atoi(input)
	if err != nil {
		return err
	}
	*target = value
	return nil
}

func SafeUnwrap[T any](v *T) T {
	if v == nil {
		v = new(T)
	}
	return *v
}
