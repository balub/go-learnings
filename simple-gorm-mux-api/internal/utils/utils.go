package utils

import (
	"io"
	"net/http"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"res": pong}`)
}
