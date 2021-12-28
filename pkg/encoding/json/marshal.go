package json

import (
	"github.com/json-iterator/go"
)

var iterator = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) ([]byte, error) {
	return iterator.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return iterator.Unmarshal(data, v)
}

func Convert(src, dest interface{}) error {
	b, err := Marshal(src)
	if err != nil {
		return err
	}
	b = FormatCamelKey(b)
	b = FormatTime(b)
	if err := Unmarshal(b, dest); err != nil {
		return err
	}
	return nil
}
