package checkout

import (
	"context"

	"github.com/ZanmoNG/gomall/app/frontend/biz/service"
	"github.com/ZanmoNG/gomall/app/frontend/biz/utils"
	checkout "github.com/ZanmoNG/gomall/app/frontend/hertz_gen/fronted/checkout"
	common "github.com/ZanmoNG/gomall/app/frontend/hertz_gen/fronted/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Checkout .
// @router /checkout [GET]
func Checkout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCheckoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "checkout", utils.WarpResponse(ctx, c, resp))
}

func CheckoutWaiting(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout.CheckoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCheckoutWaitingService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, resp))
}

func CheckoutResult(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCheckoutResultService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "result", utils.WarpResponse(ctx, c, resp))
}
