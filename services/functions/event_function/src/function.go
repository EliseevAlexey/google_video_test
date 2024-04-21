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

func cloudEventFunction(_ context.Context, e cloudevents.Event) error {
	fmt.Printf("OK: %s\n", e.Data())
	return nil
}

type TestEvent struct {
	name string
}
