package rpc

import (
	"github.com/ZanmoNG/gomall/app/frontend/conf"
	frontendUtils "github.com/ZanmoNG/gomall/app/frontend/utils"
	"github.com/ZanmoNG/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/ZanmoNG/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClient()
	})
}

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}
