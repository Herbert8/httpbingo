package service

import (
	"net/http"
	"strconv"
)

func ProcStatus(respWriter http.ResponseWriter, req *http.Request) {

	pathParams := parsePathParams(req.URL.Path, "/status/")

	// 从 URL 中解析出用于测试的 用户名、口令
	sStatusCode := ""
	if len(pathParams) >= 1 {
		sStatusCode = pathParams[0]
	}
	nStatusCode, err := strconv.Atoi(sStatusCode)
	if err != nil {
		nStatusCode = http.StatusOK
	}
	http.Error(respWriter, http.StatusText(nStatusCode), nStatusCode)
	//respWriter.WriteHeader(nStatusCode)
}
