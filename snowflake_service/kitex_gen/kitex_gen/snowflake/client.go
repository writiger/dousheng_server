// Code generated by Kitex v0.4.4. DO NOT EDIT.

package snowflake

import (
	"context"
	kitex_gen "dousheng_server/snowflake_service/kitex_gen/kitex_gen"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	NewID(ctx context.Context, Req *kitex_gen.NewIDRequest, callOptions ...callopt.Option) (r *kitex_gen.NewIDResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kSnowflakeClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kSnowflakeClient struct {
	*kClient
}

func (p *kSnowflakeClient) NewID(ctx context.Context, Req *kitex_gen.NewIDRequest, callOptions ...callopt.Option) (r *kitex_gen.NewIDResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.NewID(ctx, Req)
}