package utils

import jsoniter "github.com/json-iterator/go"

func Unmarshal(data []byte, v interface{}) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, v)
}

func Marshal(data any) ([]byte, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
}
