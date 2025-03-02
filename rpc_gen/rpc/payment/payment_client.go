package payment

import (
	"context"
	payment "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment/chargeservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() chargeservice.Client
	Service() string
	Charge(ctx context.Context, Req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := chargeservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient chargeservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() chargeservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Charge(ctx context.Context, Req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error) {
	return c.kitexClient.Charge(ctx, Req, callOptions...)
}
