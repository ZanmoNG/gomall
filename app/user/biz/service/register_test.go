package service

import (
	"context"
	"github.com/ZanmoNG/gomall/app/user/biz/dal/mysql"
	user "github.com/ZanmoNG/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
