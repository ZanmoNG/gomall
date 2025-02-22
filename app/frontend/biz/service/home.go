package service

import (
	"context"

	common "github.com/ZanmoNG/gomall/app/frontend/hertz_gen/fronted/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "Shimoe Koharu", "Price": "Free", "Picture": "/static/image/delicious_1.jpg"},
		{"Name": "Yurisaki Mika", "Price": "Free", "Picture": "/static/image/delicious_2.jpg"},
		{"Name": "MNF L'Audacieux", "Price": "Free", "Picture": "/static/image/delicious_3.png"},
		{"Name": "鳳凰モヤ", "Price": "Free", "Picture": "/static/image/delicious_4.jpg"},
		{"Name": "Yurizono Seia", "Price": "Free", "Picture": "/static/image/delicious_5.jpg"},
		{"Name": "Kasumisawa Miyu", "Price": "Free", "Picture": "/static/image/delicious_6.png"},
	}
	resp["Title"] = "Hot Pantyhose"
	resp["Items"] = items
	return resp, nil
}
