package dal

import (
	"github.com/ZanmoNG/gomall/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
