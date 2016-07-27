// Package jsonutil contains some json related helper functions.
package jsonutil

import (
	"bytes"
	"encoding/json"
)

func JsonPrettyPrint(in, prefix, indent string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), prefix, indent)
	if err != nil {
		return in
	}
	return out.String()
}

func PrettyPrintMap(m map[string]interface{}) string {
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return ""
	}
	return string(b)
}
