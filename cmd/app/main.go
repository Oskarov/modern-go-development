package main

import (
	"net/http"
	"webServer/pkg/stringutils"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		s := request.URL.Query()["s"]
		rev := stringutils.Rev(s[0])
		writer.Write([]byte(rev))
	}))
}
