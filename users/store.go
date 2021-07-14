package users

type UsersStore interface {
	Create(user *User) (*User, error)
	Get(id string) (*User, error)
	GetByUsernameAndPassword(username, password string) (*User, error)
	List() ([]User, error)
}
