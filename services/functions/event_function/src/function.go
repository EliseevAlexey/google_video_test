package function

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func init() {
	functions.CloudEvent("CloudEventFunction", cloudEventFunction)
}

func cloudEventFunction(_ context.Context, _ event.Event) error {
	fmt.Println("OK")
	return nil
}
