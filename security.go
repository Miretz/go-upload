package main

import (
	"net/http"
	"strings"
	"encoding/base64"
)

func IsUnauthorized(w http.ResponseWriter, r *http.Request) bool {
	if !BasicAuth(w, r) {
		RequestAuth(w, r)
		return true
	}
	return false
}

func BasicAuth(w http.ResponseWriter, r *http.Request) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}
	return pair[0] == string(username) && pair[1] == string(pass)
}

func RequestAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Beware! Protected website!"`)
	w.WriteHeader(401)
	w.Write([]byte("401 IsUnauthorized\n"))
	return
}

