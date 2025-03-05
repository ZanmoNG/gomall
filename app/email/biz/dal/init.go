package dal

import (
	"github.com/ZanmoNG/gomall/app/email/biz/dal/mysql"
	"github.com/ZanmoNG/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
