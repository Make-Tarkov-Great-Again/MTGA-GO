package tools

import (
	"encoding/json"
	"errors"
	"strings"
)

// Stringify returns a string representation of the given data.
func Stringify(data interface{}, oneline bool) string {
	if oneline {
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return ""
		}
		return string(jsonBytes)
	}

	var buf strings.Builder // Use a strings.Builder instead of bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(data); err != nil {
		return ""
	}
	return buf.String()
}

// ParseJSON parses the given byte slice as JSON and returns an array or object representation of the parsed data.
func ParseJSON(data *[]byte) (interface{}, error) {
	var result interface{}
	err := json.Unmarshal(*data, &result)
	if err != nil {
		return nil, err
	}

	if arr, ok := result.([]interface{}); ok {
		return arr, nil
	}

	if obj, ok := result.(map[string]interface{}); ok {
		return obj, nil
	}

	return nil, errors.New("unexpected JSON structure")
}

// ReadParsed reads a file path and parses it into an interface.
func ReadParsed(filePath string) (interface{}, error) {
	data, err := ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	result, err := ParseJSON(&data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
