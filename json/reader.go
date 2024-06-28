package json

import (
	"encoding/json"
	"io"
)

func ToJsonByReader(r io.ReadCloser) ([]byte, error) {
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return json.Marshal(b)
}

func ParseByReader[T any](r io.ReadCloser) (*T, error) {
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return Parse[T](b)
}

func ParseSliceByReader[T any](r io.ReadCloser) ([]T, error) {
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParseSeq[T](b)
}
