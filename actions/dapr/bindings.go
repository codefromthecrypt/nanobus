package dapr

import (
	"context"

	proto "github.com/dapr/dapr/pkg/proto/components/v1"

	"github.com/nanobus/nanobus/config"
	"github.com/nanobus/nanobus/resolve"
	"github.com/nanobus/nanobus/resource"
)

// Connection is the NamedLoader for a postgres connection.
func OutputBinding() (string, resource.Loader) {
	return "dapr/outputbinding.pluggable.v1", OutputBindingLoader
}

func OutputBindingLoader(ctx context.Context, with interface{}, resolver resolve.ResolveAs) (interface{}, error) {
	var c ComponentConfig
	if err := config.Decode(with, &c); err != nil {
		return nil, err
	}

	conn, err := DialConfig(ctx, &c)
	if err != nil {
		return nil, err
	}

	client := proto.NewOutputBindingClient(conn)
	_, err = client.Init(ctx, &proto.OutputBindingInitRequest{
		Metadata: &proto.MetadataRequest{
			Properties: c.Properties,
		},
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
