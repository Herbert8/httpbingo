package service

import (
	"bytes"
	"net/http"
)

func ProcDumpRequest(responseWriter http.ResponseWriter, request *http.Request) {
	var buf bytes.Buffer
	_ = request.Write(&buf)
	_, _ = writeByteSliceToResponse(responseWriter, buf.Bytes())
}
