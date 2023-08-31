package tools

import (
	"strings"

	"github.com/goccy/go-json"
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

type Data struct {
	Data json.RawMessage `json:"data"`
}

func GetJSONRawMessage(path string) json.RawMessage {
	byte, err := ReadFile(path)
	if err != nil {
		panic(err)
	}

	rawJson := json.RawMessage(byte)
	var data Data
	if strings.Contains(string(rawJson), "\"data\"") {
		err := json.Unmarshal(rawJson, &data)
		if err != nil {
			panic(err)
		}
		return data.Data
	}
	return rawJson
}
