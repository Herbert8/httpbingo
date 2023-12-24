package service

import (
	"httpbingo/internal/httputils"
	"net/http"
	"strings"
)

func ProcBasicAuth(respWriter http.ResponseWriter, req *http.Request) {
	basicAuthProcessor(respWriter, req, false, "/basic-auth/")
}

func ProcHiddenBasicAuth(respWriter http.ResponseWriter, req *http.Request) {
	basicAuthProcessor(respWriter, req, true, "/hidden-basic-auth/")
}

func basicAuthProcessor(respWriter http.ResponseWriter, req *http.Request, hiddenMode bool, prefixString string) {

	type AuthInfo struct {
		Authenticated bool   `json:"authenticated"`
		User          string `json:"user"`
	}

	pathParams := parsePathParams(req.URL.Path, 1)

	// 从 URL 中解析出用于测试的 用户名、口令
	// 如果 长度不为 2，则不是 /basic-auth/username/password 模式
	// 返回 404
	if len(pathParams) != 2 {
		httputils.ErrorNotFound(respWriter)
		return
	}
	authUsername := pathParams[0]
	authPassword := pathParams[1]

	// 通过 Header 获取的 AuthInfo
	usernameInHeader, passwordInHeader, _ := req.BasicAuth()

	// 通过 URL 获取的 AuthInfo
	usernameInUrl := req.URL.User.Username()
	passwordInUrl, _ := req.URL.User.Password()

	// 判断认证是否可以通过
	if authUsername == usernameInHeader && authPassword == passwordInHeader ||
		authUsername == usernameInUrl && authPassword == passwordInUrl {
		authInfo := AuthInfo{
			Authenticated: true,
			User:          authUsername,
		}
		writeJSONResponse(authInfo, respWriter)
	} else if !hiddenMode { // 如果不是隐式模式，则 返回 WWW-Authenticate 头
		respWriter.Header().Set("WWW-Authenticate", "Basic realm=\"Fake Realm\"")
		httputils.ErrorWithStatusCode(respWriter, http.StatusUnauthorized)
	} else { // 如果是 隐式模式，则 返回 404
		httputils.ErrorWithStatusCode(respWriter, http.StatusNotFound)
	}
}

func extractBearer(r *http.Request) (bearer string, ok bool) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", false
	}
	return parseBearer(auth)
}

func parseBearer(auth string) (bearer string, ok bool) {
	const prefix = "Bearer "
	// 这里需要注意，Bearer 与 BasicAuth 不同，BasicAuth 中的 'Basic ' 不区分大小写
	// Bearer 是区分大小写的。这些协议里有说明。
	bearer, ok = strings.CutPrefix(auth, prefix)
	return bearer, ok
}

func ProcBearer(respWriter http.ResponseWriter, req *http.Request) {
	type BearerInfo struct {
		Authenticated bool   `json:"authenticated"`
		Token         string `json:"token"`
	}
	if bearer, ok := extractBearer(req); ok {
		bearerInfo := BearerInfo{
			Authenticated: true,
			Token:         bearer,
		}
		writeJSONResponse(bearerInfo, respWriter)
	} else {
		httputils.ErrorWithStatusCode(respWriter, http.StatusUnauthorized)
	}

}
