package main

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HttpFunction", httpFunction)
}

func httpFunction(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello World: Http function")
}
