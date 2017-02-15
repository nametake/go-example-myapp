package myapp

type User struct {
	ID   int
	Name string
}

type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	CreateUser(u *User) error
}

type UserCache struct {
	cache   map[int]*User
	service UserService
}

func NewUserCache(service UserService) *UserCache {
	return &UserCache{
		cache:   make(map[int]*User),
		service: service,
	}
}

func (c *UserCache) User(id int) (*User, error) {
	if u := c.cache[id]; u != nil {
		return u, nil
	}

	u, err := c.service.User(id)
	if err != nil {
		return nil, err
	} else if u != nil {
		c.cache[id] = u
	}
	return u, err
}

type TransactionService interface {
}
