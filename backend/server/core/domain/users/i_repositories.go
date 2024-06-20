package users

type IUsersRepository interface {
	Register(user *UserModel) error
	Login(user *UserModel) (string, error)
}
