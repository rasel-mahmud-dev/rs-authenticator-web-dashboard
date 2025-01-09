package login

import "net/http"

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(w http.ResponseWriter, r *http.Request) bool
}

type BaseHandler struct {
	next Handler
}

//func (h *BaseHandler) SetNext(handler Handler) {
//	h.next = handler
//}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) HandleNext(w http.ResponseWriter, r *http.Request) bool {
	if h.next != nil {
		return h.next.Handle(w, r)
	}
	return true
}
