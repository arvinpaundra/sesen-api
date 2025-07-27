package handler

import (
	"time"

	"github.com/arvinpaundra/sesen-api/application/sse"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ShowDonationMessageHandler(c *gin.Context) {
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	key := c.Query("key")

	heartbeat := time.NewTicker(30 * time.Second)
	defer heartbeat.Stop()

	client := sse.NewClient(key, c.Writer)
	sse.Clients.Add(client)

	for {
		select {
		case <-c.Done():
			return
		case <-heartbeat.C:
			data := ":heartbeat\n"
			sse.Clients.Send(key, data)
		}
	}
}
