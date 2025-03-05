package notify

import (
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"
	"github.com/kr/pretty"
)

type NoopEmail struct {
}

func (e *NoopEmail) Send(req *email.EmailReq) error {
	_, err := pretty.Printf("%v\n", req)
	if err != nil {
		return err
	}
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
