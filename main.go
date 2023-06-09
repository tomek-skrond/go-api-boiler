package main

///////////// ENTRYPOINT /////////////

import (
	"net/http"
)

func main() {
	http.HandleFunc("/user", makeHTTPHandler(handleGetUserByID))
	http.ListenAndServe(":3000", nil)
}
