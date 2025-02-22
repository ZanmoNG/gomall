package dal

import (
	"github.com/ZanmoNG/gomall/app/user/biz/dal/mysql"
	"github.com/ZanmoNG/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
