package user

import (
	"github.com/jinzhu/gorm"
	"github.com/sarulabs/di"
)

type userUsecase struct {
	ctn  di.Container
	repo Repository
}

//NewUserUsecase will create a user.UseCase interface representation
func NewUserUsecase(ctn di.Container) UseCase {
	return &userUsecase{
		ctn:  ctn,
		repo: ctn.Get("user-repository").(Repository),
	}
}

type mysqldbUserRepository struct {
	conn *gorm.DB
}

func NewMysqlDBUserRepository(conn *gorm.DB) Repository {
	return &mysqldbUserRepository{
		conn: conn,
	}
}
