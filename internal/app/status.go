package service

import (
	"httpbingo/internal/httputils"
	"net/http"
	"strconv"
)

func ProcStatus(respWriter http.ResponseWriter, req *http.Request) {

	pathParams := parsePathParams(req.URL.Path, 1)

	// 从 URL 中解析出用于测试的 用户名、口令
	sStatusCode := ""
	if len(pathParams) >= 1 {
		sStatusCode = pathParams[0]
	}
	nStatusCode, err := strconv.Atoi(sStatusCode)
	if err != nil {
		nStatusCode = http.StatusBadRequest
		httputils.Error(respWriter, "Invalid status code", nStatusCode)
		return
	}
	if nStatusCode >= 300 && nStatusCode <= 399 {
		respWriter.Header().Set("Location", "/anything")
	}
	httputils.ErrorWithDefaultStatusText(respWriter, nStatusCode)
}
