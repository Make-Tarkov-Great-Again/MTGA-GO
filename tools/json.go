package tools

import (
	"fmt"
	"github.com/goccy/go-json"
	"log"
)

func GetJSONRawMessage(path string) json.RawMessage {
	b, err := ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	var rawMap map[string]json.RawMessage
	if err := json.UnmarshalNoEscape(b, &rawMap); err != nil {
		return b
	}

	if data, exists := rawMap["data"]; exists {
		return data
	}

	return b
}

func CheckParsingError(b []byte, err error) error {

	var msg error
	switch t := err.(type) {
	case *json.SyntaxError:
		jsn := string(b[t.Offset-50 : t.Offset])
		jsn += "<--(Invalid Character)"
		msg = fmt.Errorf("Invalid character at offset %v\n %s", t.Offset, jsn)
	case *json.UnmarshalTypeError:
		jsn := string(b[t.Offset-50 : t.Offset])
		jsn += "<--(Invalid Type)"
		msg = fmt.Errorf("Invalid value at offset %v\n %s", t.Offset, jsn)
	default:
		msg = err
	}
	return msg
}
