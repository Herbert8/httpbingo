package middleware

import (
	"bytes"
	"log"
	"net/http"
)

type LogMiddleware struct {
	Next http.Handler
}

func (receiver *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	_ = r.Write(&buf)

	const LogTemplate = `
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
%s
=======================================
Client: %s
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

`
	log.Printf(LogTemplate, buf.String(), r.RemoteAddr)
	nextHandler := receiver.Next
	if nextHandler == nil {
		nextHandler = http.DefaultServeMux
	}
	nextHandler.ServeHTTP(w, r)
}
