package xyaml

import (
	"github.com/leaxoy/x-go/xencoding"
	"gopkg.in/yaml.v2"
)

const Name = "yaml"

type yamlc struct{}

func (yamlc) Name() string {
	return Name
}

func (yamlc) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (yamlc) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func init() {
	xencoding.Register(yamlc{})
}
