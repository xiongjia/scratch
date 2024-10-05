package util

import (
	"bytes"
	"encoding/json"
	"io"
)

func JsonReaderToObject[T any](input io.Reader) (*T, error) {
	if input == nil {
		return nil, nil
	}
	var target T
	err := json.NewDecoder(input).Decode(&target)
	if err != nil {
		return nil, err
	}
	return &target, nil
}

func JsonStringToObject[T any](input string) (*T, error) {
	if len(input) <= 0 {
		return nil, nil
	}
	return JsonReaderToObject[T](bytes.NewBufferString(input))
}

func JsonStringFromObject[T any](input T) (string, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func JsonWriterFromObject[T any](writer io.Writer, input T) error {
	return json.NewEncoder(writer).Encode(input)
}
