package domain

type User struct {
	Id       int
	Email    string
	Password string
}

type UserAdaptorI interface {
	Save(u User) error
	Load() ([]User, error)
}

type UserRepository interface {
	Save(u User) error
	Load() ([]User, error)
}
