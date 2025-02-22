package dal

import (
	"github.com/ZanmoNG/gomall/app/frontend/biz/dal/mysql"
	"github.com/ZanmoNG/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
