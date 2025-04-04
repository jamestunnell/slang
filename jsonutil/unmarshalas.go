package jsonutil

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func UnmarshalAs[T any](d []byte) (*T, error) {
	var t T

	if err := json.Unmarshal(d, &t); err != nil {
		typStr := reflect.TypeOf(t).String()

		return nil, fmt.Errorf("failed to unmarshal as %s: %w", typStr, err)
	}

	return &t, nil
}

func UnmarshalAsWithPost[T any](d []byte, post func(t *T)) (*T, error) {
	var t T

	if err := json.Unmarshal(d, &t); err != nil {
		typStr := reflect.TypeOf(t).String()

		return nil, fmt.Errorf("failed to unmarshal as %s: %w", typStr, err)
	}

	return &t, nil
}
