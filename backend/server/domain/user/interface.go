package user

// IUser is the interface for user
type IUser interface {
	Find(params map[string]interface{}) ([]UserModel, error)
	Create(req *UserModel) error
	Update(req *UserModel) error
	Delete(id *string) error
}
