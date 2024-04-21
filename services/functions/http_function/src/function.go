package function

import (
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"net/http"
)

func init() {
	functions.HTTP("CloudHttpFunction", cloudHttpFunction)
}

func cloudHttpFunction(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello: cloudHttpFunction")
}
