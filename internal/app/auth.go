package service

import (
	"net/http"
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

	const statusNotFoundBody = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>404 Not Found</title>
<h1>Not Found</h1>
<p>The requested URL was not found on the server.  If you entered the URL manually please check your spelling and try again.</p>`

	pathParams := parsePathParams(req.URL.Path, 1)

	// 从 URL 中解析出用于测试的 用户名、口令
	// 如果 长度不为 2，则不是 /basic-auth/username/password 模式
	// 返回 404
	if len(pathParams) != 2 {
		http.Error(respWriter, statusNotFoundBody, http.StatusNotFound)
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
		http.Error(respWriter, "", http.StatusUnauthorized)
		respWriter.WriteHeader(401)
	} else { // 如果是 隐式模式，则 返回 404
		http.Error(respWriter, "", http.StatusNotFound)
	}
}
