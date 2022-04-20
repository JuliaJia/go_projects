package impl

import (
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go_learn/restful-api-demo/conf"
)

var Service *impl = &impl{}

type impl struct {
	log logger.Logger
	db  *sql.DB
}

func (i *impl) Init() error {
	db, err := conf.C().Mysql.GetDB()
	if err != nil {
		return err
	}
	i.log = zap.L().Named("Host")
	i.db = db
	return nil
}
