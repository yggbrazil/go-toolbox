// Lib for with some context helpers
package context

import (
	"context"
	"log"

	"github.com/aidarkhanov/nanoid/v2"
)

const TraceField = "trace_id"

// Create a context with trace_id
func CreateWithTrace() context.Context {
	traceID, err := nanoid.New()
	if err != nil {
		log.Fatalln(err)
	}

	return context.WithValue(context.Background(), TraceField, traceID)
}

// AddTrace add trace_id in context
func AddTrace(c context.Context) context.Context {
	traceID, err := nanoid.New()
	if err != nil {
		log.Fatalln(err)
	}

	return context.WithValue(c, TraceField, traceID)
}
