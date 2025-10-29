package handlers

import (
	"encoding/json"
	"io"
)

func tryDecodeJSON[T any](body io.ReadCloser) (*T, error) {
	defer body.Close()

	var msg T
	if err := json.NewDecoder(body).Decode(&msg); err != nil {
		return nil, err
	}

	return &msg, nil
}
