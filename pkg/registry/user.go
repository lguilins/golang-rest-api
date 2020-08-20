package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/lguilins/oraculosweb-api/pkg/user"
	"github.com/sarulabs/di"
)

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	service := user.NewUserUsecase(ctn)
	return service, nil
}

func buildUserMysqlRepository(ctn di.Container) (interface{}, error) {
	db := ctn.Get("mysqldb-conn").(*gorm.DB)
	repo := user.NewMysqlDBUserRepository(db)
	return repo, nil
}
