package dal

import (
	"github.com/ZanmoNG/gomall/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
