package api

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

const requestTimeoutLimit = 5 * time.Second

// uploadWithPresignedUrl uploads files to remote bucket using Presigned URL.
func uploadWithPresignedUrl(preSignedUrl, objectKey string) (int, error) {
	object, err := os.ReadFile(objectKey)
	if err != nil {
		return 1, fmt.Errorf("failed to read file: %w:", err)
	}

	req, err := http.NewRequest("PUT", preSignedUrl, bytes.NewReader(object))

	client := &http.Client{
		Timeout: requestTimeoutLimit,
	}

	resp, err := client.Do(req)
	if err != nil {
		return 1, fmt.Errorf("failed to send request: %w:", err)
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}
