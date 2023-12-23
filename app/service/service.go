package service

import (
	"time"
	"wusthelper-todolist-service/app/conf"
	"wusthelper-todolist-service/app/dao"
)

type Service struct {
	config   *conf.Config
	dao      *dao.Dao
	timezone *time.Location
}

func New(c *conf.Config) (service *Service) {
	service = &Service{
		config: c,
		dao:    dao.New(c),
	}

	return service
}
