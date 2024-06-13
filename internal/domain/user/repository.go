package user

type UserRepository interface {
	Add(user User) error
	FindByID(id string) (User, error)
}