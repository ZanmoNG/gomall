.PHONY: gen-frontend
gen-frontend:
	cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/cloudwego/biz-demo/gomall/app/frontend -I ../../idl
