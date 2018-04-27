// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/bottos-project/core/api/core-api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	github.com/bottos-project/core/api/core-api.proto

It has these top-level messages:
	PushTrxResponse
	QueryTrxRequest
	QueryTrxResponse
	QueryBlockRequest
	QueryBlockResponse
	QueryChainInfoRequest
	QueryChainInfoResponse
	QueryAccountRequest
	QueryAccountResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/bottos-project/core/common/types"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CoreApi service

type CoreApiClient interface {
	PushTrx(ctx context.Context, in *types.Transaction, opts ...client.CallOption) (*PushTrxResponse, error)
	QueryTrx(ctx context.Context, in *QueryTrxRequest, opts ...client.CallOption) (*QueryTrxResponse, error)
	QueryBlock(ctx context.Context, in *QueryBlockRequest, opts ...client.CallOption) (*QueryBlockResponse, error)
	QueryChainInfo(ctx context.Context, in *QueryChainInfoRequest, opts ...client.CallOption) (*QueryChainInfoResponse, error)
	QueryAccount(ctx context.Context, in *QueryAccountRequest, opts ...client.CallOption) (*QueryAccountResponse, error)
}

type coreApiClient struct {
	c           client.Client
	serviceName string
}

func NewCoreApiClient(serviceName string, c client.Client) CoreApiClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "api"
	}
	return &coreApiClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *coreApiClient) PushTrx(ctx context.Context, in *types.Transaction, opts ...client.CallOption) (*PushTrxResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CoreApi.PushTrx", in)
	out := new(PushTrxResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreApiClient) QueryTrx(ctx context.Context, in *QueryTrxRequest, opts ...client.CallOption) (*QueryTrxResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CoreApi.QueryTrx", in)
	out := new(QueryTrxResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreApiClient) QueryBlock(ctx context.Context, in *QueryBlockRequest, opts ...client.CallOption) (*QueryBlockResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CoreApi.QueryBlock", in)
	out := new(QueryBlockResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreApiClient) QueryChainInfo(ctx context.Context, in *QueryChainInfoRequest, opts ...client.CallOption) (*QueryChainInfoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CoreApi.QueryChainInfo", in)
	out := new(QueryChainInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreApiClient) QueryAccount(ctx context.Context, in *QueryAccountRequest, opts ...client.CallOption) (*QueryAccountResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CoreApi.QueryAccount", in)
	out := new(QueryAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CoreApi service

type CoreApiHandler interface {
	PushTrx(context.Context, *types.Transaction, *PushTrxResponse) error
	QueryTrx(context.Context, *QueryTrxRequest, *QueryTrxResponse) error
	QueryBlock(context.Context, *QueryBlockRequest, *QueryBlockResponse) error
	QueryChainInfo(context.Context, *QueryChainInfoRequest, *QueryChainInfoResponse) error
	QueryAccount(context.Context, *QueryAccountRequest, *QueryAccountResponse) error
}

func RegisterCoreApiHandler(s server.Server, hdlr CoreApiHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&CoreApi{hdlr}, opts...))
}

type CoreApi struct {
	CoreApiHandler
}

func (h *CoreApi) PushTrx(ctx context.Context, in *types.Transaction, out *PushTrxResponse) error {
	return h.CoreApiHandler.PushTrx(ctx, in, out)
}

func (h *CoreApi) QueryTrx(ctx context.Context, in *QueryTrxRequest, out *QueryTrxResponse) error {
	return h.CoreApiHandler.QueryTrx(ctx, in, out)
}

func (h *CoreApi) QueryBlock(ctx context.Context, in *QueryBlockRequest, out *QueryBlockResponse) error {
	return h.CoreApiHandler.QueryBlock(ctx, in, out)
}

func (h *CoreApi) QueryChainInfo(ctx context.Context, in *QueryChainInfoRequest, out *QueryChainInfoResponse) error {
	return h.CoreApiHandler.QueryChainInfo(ctx, in, out)
}

func (h *CoreApi) QueryAccount(ctx context.Context, in *QueryAccountRequest, out *QueryAccountResponse) error {
	return h.CoreApiHandler.QueryAccount(ctx, in, out)
}