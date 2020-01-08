package xjson

import (
	"encoding/json"
	"github.com/leaxoy/x-go/xencoding"
)

const Name = "json"

type jsonc struct{}

func (jsonc) Name() string { return Name }

func (jsonc) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonc) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func init() {
	xencoding.Register(jsonc{})
}
