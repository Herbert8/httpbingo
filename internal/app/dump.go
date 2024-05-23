package service

import (
	"bytes"
	"log"
	"net/http"
)

func ProcDumpRequest(responseWriter http.ResponseWriter, request *http.Request) {
	var buf bytes.Buffer
	err := request.Write(&buf)
	if err != nil {
		log.Println("ProcDumpRequest Request.Write Buffer Error:", err)
	}

	_, _ = writeByteSliceToResponse(responseWriter, []byte(buf.String()))
}
