package sse

import (
	"fmt"
	"net/http"
)

type Client struct {
	id string
	rw http.ResponseWriter
}

func NewClient(id string, rw http.ResponseWriter) *Client {
	return &Client{
		id: id,
		rw: rw,
	}
}

func (c *Client) Send(data string) error {
	flusher, ok := c.rw.(http.Flusher)
	if !ok {
		return fmt.Errorf("failed to send data")
	}

	_, err := fmt.Fprintf(c.rw, "data: %s\n\n", data)
	if err != nil {
		return err
	}

	flusher.Flush()

	return nil
}
