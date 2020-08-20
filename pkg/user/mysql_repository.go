package user

//Repository is the user database contract
type Repository interface {
	FindAll() string
}

func (db *mysqldbUserRepository) FindAll() string {
	User := "LUIZZZZ"

	return User
}
