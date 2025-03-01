package dal

import (
	"github.com/ZanmoNG/gomall/app/order/biz/dal/mysql"
	"github.com/ZanmoNG/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
