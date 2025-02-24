package service

import (
	"context"
	"github.com/ZanmoNG/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/ZanmoNG/gomall/app/frontend/hertz_gen/fronted/product"
	rpcproduct "github.com/ZanmoNG/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	products, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"items": products.Results,
		"q":     req.Q,
	}, nil
}
