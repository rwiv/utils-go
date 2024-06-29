package ser

import (
	"encoding/json"
)

func Parse[T any](b []byte) (*T, error) {
	v := new(T)
	err := json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ParseSeq[T any](b []byte) ([]T, error) {
	v := make([]T, 0)
	err := json.Unmarshal(b, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ToPrettyJson(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
