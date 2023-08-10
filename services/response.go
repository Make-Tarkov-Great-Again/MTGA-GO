package services

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"net/http"
)

func ZlibReply(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	bytes := convertDataToByte(w, data)
	zlibDeflate(w, bytes)
}

func ZlibJSONReply(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	bytes := convertDataToByte(w, data)
	zlibDeflate(w, bytes)
}

func convertDataToByte(w http.ResponseWriter, data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		panic(err)
	}
	return bytes
}
func zlibDeflate(w http.ResponseWriter, data []byte) {
	var buffer bytes.Buffer
	writer := zlib.NewWriter(&buffer)
	_, err := writer.Write(data)
	if err != nil {
		writer.Close()
		http.Error(w, "Failed to write compressed data", http.StatusInternalServerError)
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		writer.Close()
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		panic(err)
	}

	writer.Close()

	_, err = w.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
}
