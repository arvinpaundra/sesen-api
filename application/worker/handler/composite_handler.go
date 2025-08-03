package handler

import (
	"context"

	"github.com/arvinpaundra/sesen-api/core/event"
)

type CompositeHandler struct {
	handlers  []event.DomainEventHandler
	eventType string
}

func NewCompositeHandler(eventType string) *CompositeHandler {
	return &CompositeHandler{
		eventType: eventType,
	}
}

func (h *CompositeHandler) AddHandler(handleFunc event.DomainEventHandler) *CompositeHandler {
	h.handlers = append(h.handlers, handleFunc)
	return h
}

func (h *CompositeHandler) Handle(ctx context.Context, eventData []byte) error {
	for _, handler := range h.handlers {
		if err := handler.Handle(ctx, eventData); err != nil {
			return err
		}
	}

	return nil
}
