package user

type UserRepository interface {
	Add(user User) error
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
	UpdateNick(id string, nick string) error
}
