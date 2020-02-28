package sqlike

import "context"

type ExecutedMarkerValue int

const (
	executedMarker = ExecutedMarkerValue(1)
)

func MarkAsExecuted(ctx context.Context) context.Context {
	return context.WithValue(ctx, executedMarker, true)
}

func isExecuted(ctx context.Context) bool {
	executedV := ctx.Value(executedMarker)
	if executedV != nil {
		if executed, ok := executedV.(bool); ok && executed {
			return true
		}
	}
	return false
}
