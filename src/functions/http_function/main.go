package main

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("httpFunction", HttpFunction)
}

func HttpFunction(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "OK")
}
