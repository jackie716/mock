package main

import (
	"encoding/json"

	"github.com/go-yaml/yaml"
)

func yamlize(o interface{}) string {
	b, _ := yaml.Marshal(o)
	return string(b)
}

func jsonlize(o interface{}) string {
	b, _ := json.Marshal(o)
	return string(b)
}
