package dao

import (
	"github.com/golang/glog"
	"github.com/oikomi/FishChatServer2/common/dao/xredis"
	"github.com/oikomi/FishChatServer2/server/register/conf"
)

type Dao struct {
	redis   *xredis.Pool
	Mysql   *Mysql
	MongoDB *MongoDB
}

func NewDao() (dao *Dao) {
	mysql := NewMysql()
	mongo, err := NewMongoDB()
	if err != nil {
		glog.Error(err)
		return
	}
	dao = &Dao{
		redis:   xredis.NewPool(conf.Conf.Redis.Redis),
		Mysql:   mysql,
		MongoDB: mongo,
	}
	return
}
