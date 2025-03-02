package main

import (
	"context"
	payment "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/biz-demo/gomall/app/payment/biz/service"
)

// ChargeServiceImpl implements the last service interface defined in the IDL.
type ChargeServiceImpl struct{}

// Charge implements the ChargeServiceImpl interface.
func (s *ChargeServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}
