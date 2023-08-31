package services

import (
	"bytes"
	"compress/zlib"
	"io"
	"net/http"
	"strings"

	"github.com/goccy/go-json"
)

func ZlibReply(w http.ResponseWriter, data interface{}) {
	zlibDeflate(w, data)
}

func ZlibJSONReply(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	zlibDeflate(w, data)
}

func ZlibInflate(r *http.Request) *bytes.Buffer {

	// Check if the request header includes "Unity"
	if strings.Contains(r.Header.Get("User-Agent"), "Unity") {
		buffer := &bytes.Buffer{}

		// Inflate r.Body with zlib
		reader, err := zlib.NewReader(r.Body)
		if err != nil {
			panic(err)
		}
		defer reader.Close()

		// Read the decompressed data
		_, err = io.Copy(buffer, reader)
		if err != nil {
			panic(err)
		}

		return buffer
	}
	return nil
}

func zlibDeflate(w http.ResponseWriter, data interface{}) {

	// Convert data to JSON bytes
	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal data to JSON", http.StatusInternalServerError)
		return
	}

	// Compress the JSON bytes
	compressedBytes := compressZlib(bytes)

	// Set appropriate response headers
	w.WriteHeader(http.StatusOK)

	// Write the compressed data to the response
	_, err = w.Write(compressedBytes)
	if err != nil {
		http.Error(w, "Failed to write compressed data", http.StatusInternalServerError)
		return
	}
}

func compressZlib(data []byte) []byte {
	buffer := &bytes.Buffer{}
	writer := zlib.NewWriter(buffer)

	defer writer.Close()

	_, err := writer.Write(data)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}
