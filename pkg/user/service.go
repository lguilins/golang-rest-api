package user

func (u *userUsecase) GetAll() string {
	user := u.repo.FindAll()

	return user
}
