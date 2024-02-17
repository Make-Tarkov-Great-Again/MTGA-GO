package pkg

import (
	"bytes"
	"compress/zlib"
	"fmt"
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

	compressed, err := ZlibDeflate(data)
	if err != nil {
		log.Printf("Failed to Deflate compressed data: %s", err)
		return
	}
	_, err = w.Write(compressed)
	if err != nil {
		http.Error(w, "Failed to write compressed data", http.StatusInternalServerError)
		return
	}
}

func ZlibInflate(r *http.Request) ([]byte, error) {
	// Inflate r.Body with zlib
	reader, err := zlib.NewReader(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create zlib reader: %w", err)
	}
	defer func() {
		if err := reader.Close(); err != nil {
			log.Println("error closing zlib reader:", err)
		}
	}()

	// Read the decompressed data
	buffer, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress data: %w", err)
	}

	return buffer, nil
}

func ZlibDeflate(data any) ([]byte, error) {
	input, err := json.MarshalNoEscape(data)
	if err != nil {
		log.Println("Failed to marshal data to JSON")
		return nil, err
	}

	return compressZlib(input, zlib.BestSpeed)
}

func compressZlib(data []byte, speed int) ([]byte, error) {
	buffer := bytes.NewBuffer(make([]byte, 0, len(data)))
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

	return buffer.Bytes(), nil
}

func CreateCachedResponse(input any) ([]byte, error) {
	dataData, err := json.MarshalNoEscape(ApplyResponseBody(input))
	if err != nil {
		log.Fatal(err)
	}

	//return compressZlib(dataData, zlib.BestCompression)
	return compressZlib(dataData, zlib.BestSpeed)
}
