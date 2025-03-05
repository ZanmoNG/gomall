package email

import (
	"github.com/ZanmoNG/gomall/app/email/infra/mq"
	"github.com/ZanmoNG/gomall/app/email/infra/notify"
	"github.com/ZanmoNG/gomall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}

		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
	})

	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		err := sub.Unsubscribe()
		if err != nil {
			return
		}
		mq.Nc.Close()
	})
}
