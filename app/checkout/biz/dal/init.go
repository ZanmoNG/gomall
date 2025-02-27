package dal

import (
	"github.com/ZanmoNG/gomall/app/checkout/biz/dal/mysql"
	"github.com/ZanmoNG/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
