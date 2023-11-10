package services

import (
	"bytes"
	"compress/zlib"
	"github.com/goccy/go-json"
	"io"
	"log"
	"net/http"
)

var cachedZlib = map[string][]byte{
	"/client/settings":      nil,
	"/client/customization": nil,
	"/client/locale/":       nil,
	"/client/items":         nil,
	"/client/globals":       nil,
	"/client/locations":     nil,
	"/client/game/config":   nil,
	"/client/languages":     nil,
	"/client/menu/locale/":  nil,
	//"/client/location/getLocalloot": {}, don't fully understand why this would be cached
}

func ZlibReply(w http.ResponseWriter, path string, data any) {
	zlibDeflate(w, path, data)
}

func ZlibJSONReply(w http.ResponseWriter, path string, data any) {
	w.Header().Set("Content-Type", "application/json")
	zlibDeflate(w, path, data)
}

func ZlibInflate(r *http.Request) *bytes.Buffer {
	buffer := new(bytes.Buffer)

	// Inflate r.Body with zlib
	reader, err := zlib.NewReader(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(reader io.ReadCloser) {
		err := reader.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(reader)

	// Read the decompressed data
	_, err = io.Copy(buffer, reader)
	if err != nil {
		log.Fatalln(err)
	}

	return buffer

}

func zlibDeflate(w http.ResponseWriter, path string, data any) {
	cached, ok := cachedZlib[path]
	if ok && cached != nil {
		w.WriteHeader(http.StatusOK)

		_, err := w.Write(cached)
		if err != nil {
			http.Error(w, "Failed to write compressed data", http.StatusInternalServerError)
			return
		}
		return
	}

	// Convert data to JSON bytes
	input, err := json.MarshalNoEscape(data)
	if err != nil {
		http.Error(w, "Failed to marshal data to JSON", http.StatusInternalServerError)
		return
	}

	// Compress the JSON bytes
	compressed := compressZlib(input)
	if ok {
		cachedZlib[path] = compressed
	}

	// Set appropriate response headers
	w.WriteHeader(http.StatusOK)

	// Write the compressed data to the response
	_, err = w.Write(compressed)
	if err != nil {
		http.Error(w, "Failed to write compressed data", http.StatusInternalServerError)
		return
	}
}

func compressZlib(data []byte) []byte {
	buffer := &bytes.Buffer{}
	writer, _ := zlib.NewWriterLevel(buffer, zlib.BestSpeed)

	defer func(writer *zlib.Writer) {
		err := writer.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(writer)

	_, err := writer.Write(data)
	if err != nil {
		log.Fatalln(err)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatalln(err)
	}

	return buffer.Bytes()
}
