.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/cloudwego/biz-demo/gomall/app/frontend -I ../../idl

.PHONY: gen-product
gen-frontend:
	@cd rpc_gen && cwgo client --type RPC --service product --module github.com/cloudwego/biz-demo/gomall/rpc_gen -I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module github.com/cloudwego/biz-demo/gomall/app/product --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto