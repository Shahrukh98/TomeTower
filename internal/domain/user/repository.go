package user

type UserRepository interface {
	Add(user User) error
	FindByEmail(email string) (User, error)
	FindByID(id string) (User, error)
	UpdateNick(id string, nick string) error
}
