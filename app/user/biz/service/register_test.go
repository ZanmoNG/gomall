package service

import (
	"context"
	"testing"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "demo@admin.com",
		Password:        "12312",
		ConfirmPassword: "12312",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
