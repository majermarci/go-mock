package main

import (
	"bytes"
	"encoding/json"
)

func isJSON(data string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(data), &js) == nil
}

func formatJSON(data string) (string, error) {
	var buf bytes.Buffer
	err := json.Compact(&buf, []byte(data))
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
