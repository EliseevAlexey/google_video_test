package function

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("CloudHttpFunction", cloudHttpFunction)
}

func cloudHttpFunction(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello: cloudHttpFunction")
}
