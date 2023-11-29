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
	defer func() {
		if err := reader.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Read the decompressed data
	_, err = io.Copy(buffer, reader)
	if err != nil {
		log.Panicln(err)
	}

	return buffer
}

func ZlibDeflate(data any) []byte {
	input, err := json.MarshalNoEscape(data)
	if err != nil {
		log.Println("Failed to marshal data to JSON")
		return nil
	}

	return compressZlib(input, zlib.BestSpeed)
}

func compressZlib(data []byte, speed int) []byte {
	buffer := bytes.NewBuffer(nil)
	writer, err := zlib.NewWriterLevel(buffer, speed)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := writer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	_, err = writer.Write(data)
	if err != nil {
		log.Panicln(err)
	}

	err = writer.Flush()
	if err != nil {
		log.Panicln(err)
	}

	return buffer.Bytes()
}

func CreateCachedResponse(input any) []byte {
	dataData, err := json.MarshalNoEscape(ApplyResponseBody(input))
	if err != nil {
		log.Fatal(err)
	}

	//return compressZlib(dataData, zlib.BestCompression)
	return compressZlib(dataData, zlib.BestSpeed)
}
