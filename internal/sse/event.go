package sse

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
)

type Event struct {
	c *gin.Context
}

func NewEvent(c *gin.Context) *Event {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Render(-1, sse.Event{
		Event: "open",
		Retry: 3600 * 1000,
	}) // SSE 重试1h，防止重试
	c.SSEvent("start", time.Now().String())
	c.Writer.Flush()

	e := &Event{c: c}
	return e
}

func (e *Event) Context() *gin.Context {
	return e.c
}

func (e *Event) SSEvent(name string, message any) {
	e.c.SSEvent(name, message)
}

func (e *Event) Request() *http.Request {
	return e.Context().Request
}

func (e *Event) Flush() {
	select {
	case <-e.Request().Context().Done():
	default:
		e.Context().Writer.Flush()
	}
}
