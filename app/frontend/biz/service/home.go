package service

import (
	"context"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {

	// resp := []map[string]any{
	// 	{"Name": "创造性AI", "Price": "119", "Picture": "/static/image/bot_AI.png"},
	// 	{"Name": "冷静头脑", "Price": "56", "Picture": "/static/image/bot_calm.png"},
	// 	{"Name": "偏差认知", "Price": "116", "Picture": "/static/image/bot_cognition.png"},
	// 	{"Name": "回想形态", "Price": "122", "Picture": "/static/image/bot_echo.png"},
	// 	{"Name": "冰寒", "Price": "76", "Picture": "/static/image/bot_freezing.png"},
	// }
	// return resp, nil
	//
	// args need changed to []map[string]any too
	//
	// change resp to a complicate map, whose Items is a map list
	// with each item as list member

	// var resp = make(map[string]any)
	// items := []map[string]any{
	// 	{"Name": "creative AI", "Price": "119", "Picture": "/static/image/bot_AI.png"},
	// 	{"Name": "calm down bot", "Price": "56", "Picture": "/static/image/bot_calm.png"},
	// 	{"Name": "cognition?", "Price": "116", "Picture": "/static/image/bot_cognition.png"},
	// 	{"Name": "echo mode", "Price": "122", "Picture": "/static/image/bot_echo.png"},
	// 	{"Name": "freezing", "Price": "76", "Picture": "/static/image/bot_freezing.png"},
	// }
	// resp["Title"] = "Hot Sales"
	// resp["Items"] = items

	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Hot sale",
		"items": products.Products,
	}, nil
}
