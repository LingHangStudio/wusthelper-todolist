package dao

import (
	"wusthelper-todolist-service/app/conf"
	"wusthelper-todolist-service/library/database"
	"wusthelper-todolist-service/library/log"
	"xorm.io/xorm"
)

type Dao struct {
	db *xorm.Engine
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		db: database.NewMysql(&c.Database),
	}

	return
}

func (d *Dao) Close() {
	dbErr := d.db.Close()
	if dbErr != nil {
		log.Warn("[dao]关闭数据库连接出错")
	}
}
