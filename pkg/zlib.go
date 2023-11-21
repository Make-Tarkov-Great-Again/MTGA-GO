package pkg

import (
	"bytes"
	"compress/zlib"
	"github.com/goccy/go-json"
	"io"
	"log"
	"net/http"
)

// SendJSONReply sends pre-compressed zlib JSON to client
func SendJSONReply(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write(data)
	if err != nil {
		http.Error(w, "Failed to write compressed data", http.StatusInternalServerError)
		return
	}
}

// SendZlibJSONReply compresses and sends JSON with zlib.BestSpeed
func SendZlibJSONReply(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	compressed := ZlibDeflate(data)
	_, err := w.Write(compressed)
	if err != nil {
		http.Error(w, "Failed to write compressed data", http.StatusInternalServerError)
		return
	}
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

func ZlibDeflate(data any) []byte {
	// Convert data to JSON bytes
	input, err := json.MarshalNoEscape(data)
	if err != nil {
		log.Println("Failed to marshal data to JSON")
		return nil
	}

	// Compress the JSON bytes
	return compressZlib(input, zlib.BestSpeed)
}

func compressZlib(data []byte, speed int) []byte {
	buffer := &bytes.Buffer{}
	writer, _ := zlib.NewWriterLevel(buffer, speed)

	defer func(writer *zlib.Writer) {
		err := writer.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(writer)

	_, err := writer.Write(data)
	if err != nil {
		log.Panicln(err)
	}

	err = writer.Flush()
	if err != nil {
		log.Panicln(err)
	}

	return buffer.Bytes()
}

func CreateCachedResponse(input any) *[]byte {
	dataData, err := json.MarshalNoEscape(ApplyResponseBody(input))
	if err != nil {
		log.Fatal(err)
	}

	dataZlib := compressZlib(dataData, zlib.BestCompression)

	dataSlice := make([]byte, 0, len(dataZlib))
	dataSlice = append(dataSlice, dataZlib...)

	return &dataSlice
}
