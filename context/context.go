// Lib for with some context helpers
package context

import (
	"context"
	"log"

	"github.com/aidarkhanov/nanoid/v2"
)

const TraceField = "trace_id"

// Create a context with TraceID
func CreateWithTrace() context.Context {
	ctx := context.Background()

	traceID, err := nanoid.New()
	if err != nil {
		log.Fatalln(err)
	}

	return context.WithValue(ctx, TraceField, traceID)
}
