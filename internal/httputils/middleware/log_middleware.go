package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type LogMiddleware struct {
	Next http.Handler
}

func (receiver *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const LogTemplate = `
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
%s
=======================================
Client: %s
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

`
	nextHandler := receiver.Next
	if nextHandler == nil {
		nextHandler = http.DefaultServeMux
	}

	// 由于 http.Request 的 Body 只能读取一次，这里记录之后会导致后续访问异常
	// 所以这里做一些特殊处理：
	// 先读取到 Buffer，然后封装后放回 Body 中，通过 Buffer 中的内容进行日志记录
	// 读取 Body 的数据，留存，备用
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ServeHTTP ReadAll RequestBody Error:", err)
	}
	// 将读取到的 Body 重新包装后写回 Request，以便以后的流程读取
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	// 读取完整的 Request 信息，放入 Buffer
	var reqBuf bytes.Buffer
	err = r.Write(&reqBuf)
	if err != nil {
		log.Println("ServeHTTP Write Request Buffer Error:", err)
	}
	// 使用 Buffer 中记录的 Request 数据记录日志
	log.Printf(LogTemplate, reqBuf.String(), r.RemoteAddr)
	// 使用备份的 Body 数据，包装后重新放回 Request，以便后续流程使用
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	// 执行后续调用
	nextHandler.ServeHTTP(w, r)
}
