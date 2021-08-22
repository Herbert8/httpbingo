package service

import (
	"net/http"
)


func ProcResponseHeader(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()
	values := r.Form

	for key, values := range values {
		for _, val := range values {
			//println(key, val)
			w.Header().Add(key, val)
		}
	}

	writeJSONResponse(values2Map(values), w)

}