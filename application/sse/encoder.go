package sse

import "encoding/json"

func EncodeJSON(data any) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
