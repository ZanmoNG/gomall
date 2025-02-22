package dal

import (
	"awesomeProject/app/user/biz/dal/mysql"
	"awesomeProject/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
