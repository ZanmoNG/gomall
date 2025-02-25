package dal

import (
	"github.com/ZanmoNG/gomall/app/cart/biz/dal/mysql"
	"github.com/ZanmoNG/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
