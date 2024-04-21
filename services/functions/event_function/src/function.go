package function

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func init() {
	functions.CloudEvent("CloudEventFunction", cloudEventFunction)
}

func cloudEventFunction(_ context.Context, _ cloudevents.Event) error {
	fmt.Println("OK")
	return nil
}
