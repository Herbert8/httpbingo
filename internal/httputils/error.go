package httputils

import (
	"fmt"
	"net/http"
)

func Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, _ = fmt.Fprint(w, error)
}

func ErrorWithStatusCode(w http.ResponseWriter, code int) {
	Error(w, "", code)
}

func ErrorWithDefaultStatusText(w http.ResponseWriter, code int) {
	Error(w, http.StatusText(code), code)
}

func ErrorNotFound(w http.ResponseWriter) {
	const statusNotFoundBody = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>404 Not Found</title>
<h1>Not Found</h1>
<p>The requested URL was not found on the server.  If you entered the URL manually please check your spelling and try again.</p>`

	Error(w, statusNotFoundBody, http.StatusNotFound)
}
