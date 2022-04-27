// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package product

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Product interface {
		Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
		Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
		Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
		Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	client := NewProductClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultProduct) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	client := NewProductClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

func (m *defaultProduct) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	client := NewProductClient(m.cli.Conn())
	return client.Remove(ctx, in, opts...)
}

func (m *defaultProduct) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	client := NewProductClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}
