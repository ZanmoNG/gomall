package service

import (
	"context"
	frontendUtils "github.com/ZanmoNG/gomall/app/frontend/biz/utils"
	"github.com/ZanmoNG/gomall/app/frontend/infra/rpc"

	cart "github.com/ZanmoNG/gomall/app/frontend/hertz_gen/fronted/cart"
	common "github.com/ZanmoNG/gomall/app/frontend/hertz_gen/fronted/common"
	rpccart "github.com/ZanmoNG/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  uint32(req.ProductNum),
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
