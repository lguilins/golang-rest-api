package registry

import (
	"github.com/lguilins/oraculosweb-api/infra"
	"github.com/sarulabs/di"
)

func buildMysqlDBConn(ctn di.Container) (interface{}, error) {
	conn, err := infra.ConnectMysqlDB()
	return conn, err
}
