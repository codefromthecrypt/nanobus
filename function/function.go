package function

import (
	"context"
)

type Function struct {
	Namespace string `json:"namespace" msgpack:"namespace"`
	Operation string `json:"operation" msgpack:"operation"`
}

type functionKey struct{}

func FromContext(ctx context.Context) Function {
	v := ctx.Value(functionKey{})
	if v == nil {
		return Function{}
	}
	c, _ := v.(Function)

	return c
}

func ToContext(ctx context.Context, function Function) context.Context {
	return context.WithValue(ctx, functionKey{}, function)
}
