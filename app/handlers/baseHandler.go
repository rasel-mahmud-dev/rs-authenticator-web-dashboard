package handlers

import (
	"rs/auth/app/context"
)

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(c *context.BaseContext) bool
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) HandleNext(c *context.BaseContext) bool {
	if h.next != nil {
		return h.next.Handle(c)
	}
	return true
}
