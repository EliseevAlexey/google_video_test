package main

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func main() {}

func init() {
	functions.HTTP("HttpFunction", HttpFunction)
}

func HttpFunction(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello World: Http function")
}
