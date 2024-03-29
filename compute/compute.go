package compute

import (
	"context"
	"io"

	"github.com/nanobus/iota/go/wasmrs/invoke"
	"github.com/nanobus/iota/go/wasmrs/operations"
	"github.com/nanobus/iota/go/wasmrs/payload"
	"github.com/nanobus/iota/go/wasmrs/rx/flux"
	"github.com/nanobus/iota/go/wasmrs/rx/mono"

	"github.com/nanobus/nanobus/registry"
)

type (
	NamedLoader = registry.NamedLoader[Invoker]
	Loader      = registry.Loader[Invoker]
	Registry    = registry.Registry[Invoker]

	BusInvoker   func(ctx context.Context, namespace, service, function string, input interface{}) (interface{}, error)
	StateInvoker func(ctx context.Context, namespace, id, key string) ([]byte, error)

	Invoker interface {
		io.Closer
		Operations() operations.Table

		FireAndForget(context.Context, payload.Payload)
		RequestResponse(context.Context, payload.Payload) mono.Mono[payload.Payload]
		RequestStream(context.Context, payload.Payload) flux.Flux[payload.Payload]
		RequestChannel(context.Context, payload.Payload, flux.Flux[payload.Payload]) flux.Flux[payload.Payload]

		SetRequestResponseHandler(index uint32, handler invoke.RequestResponseHandler)
		SetFireAndForgetHandler(index uint32, handler invoke.FireAndForgetHandler)
		SetRequestStreamHandler(index uint32, handler invoke.RequestStreamHandler)
		SetRequestChannelHandler(index uint32, handler invoke.RequestChannelHandler)
	}
)
