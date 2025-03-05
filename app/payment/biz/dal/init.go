package dal

import (
	"github.com/cloudwego/biz-demo/gomall/app/payment/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
