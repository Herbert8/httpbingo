package service

import (
	"fmt"
	"net/http"
	"strconv"
)

func ProcRedirectTo(w http.ResponseWriter, r *http.Request) {

	statusCode, err := strconv.Atoi(r.FormValue("status_code"))

	if err != nil {
		statusCode = 302
	}

	redirectTo(w, statusCode, r.FormValue("url"))
}

func redirectTo(responseWriter http.ResponseWriter, statusCode int, redirectToLocation string) {
	if statusCode < 300 || statusCode > 399 {
		statusCode = 302
	}
	responseWriter.Header().Set("Location", redirectToLocation)
	responseWriter.WriteHeader(statusCode)
}

func ProcWebRedirectTo(w http.ResponseWriter, r *http.Request) {
	const HTMLTemplate = `<html>
<meta http-equiv="refresh" content="%[1]d;url=%[2]s">
<body>
%s
</body>
</html>
`
	var delaySeconds int
	if delaySeconds, _ = strconv.Atoi(r.FormValue("delay")); delaySeconds > 10 {
		delaySeconds = 10
	}
	var sUrl string
	if sUrl = r.FormValue("url"); sUrl == "" {
		sUrl = "http://baidu.com"
	}
	var sBody string
	if delaySeconds != 0 {
		sBody = fmt.Sprintf("Wait for %d seconds and then redirect to %s.", delaySeconds, sUrl)
	}

	retHTML := fmt.Sprintf(HTMLTemplate, delaySeconds, sUrl, sBody)
	_, _ = w.Write([]byte(retHTML))
}
