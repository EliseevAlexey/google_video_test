package function

import (
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"net/http"
)

func init() {
	functions.HTTP("HttpFunction", httpFunction)
}

func httpFunction(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello: httpFunction")
}
