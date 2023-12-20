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
