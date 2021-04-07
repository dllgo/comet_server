// Code generated by goctl. DO NOT EDIT!
// Source: comet_server.proto

//go:generate mockgen -destination ./cometserver_mock.go -package cometserverclient -source $GOFILE

package cometserverclient

import (
	"context"

	"github.com/dllgo/comet_server/comet_server"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Request     = comet_server.Request
	RequestAck  = comet_server.RequestAck
	Response    = comet_server.Response
	ResponseAck = comet_server.ResponseAck

	CometServer interface {
		Send2User(ctx context.Context, in *Request) (*Response, error)
		Send2Friend(ctx context.Context, in *Request) (*Response, error)
		Send2Group(ctx context.Context, in *Request) (*Response, error)
		Send2Room(ctx context.Context, in *Request) (*Response, error)
		Send2Push(ctx context.Context, in *Request) (*Response, error)
	}

	defaultCometServer struct {
		cli zrpc.Client
	}
)

func NewCometServer(cli zrpc.Client) CometServer {
	return &defaultCometServer{
		cli: cli,
	}
}

func (m *defaultCometServer) Send2User(ctx context.Context, in *Request) (*Response, error) {
	client := comet_server.NewCometServerClient(m.cli.Conn())
	return client.Send2User(ctx, in)
}

func (m *defaultCometServer) Send2Friend(ctx context.Context, in *Request) (*Response, error) {
	client := comet_server.NewCometServerClient(m.cli.Conn())
	return client.Send2Friend(ctx, in)
}

func (m *defaultCometServer) Send2Group(ctx context.Context, in *Request) (*Response, error) {
	client := comet_server.NewCometServerClient(m.cli.Conn())
	return client.Send2Group(ctx, in)
}

func (m *defaultCometServer) Send2Room(ctx context.Context, in *Request) (*Response, error) {
	client := comet_server.NewCometServerClient(m.cli.Conn())
	return client.Send2Room(ctx, in)
}

func (m *defaultCometServer) Send2Push(ctx context.Context, in *Request) (*Response, error) {
	client := comet_server.NewCometServerClient(m.cli.Conn())
	return client.Send2Push(ctx, in)
}
