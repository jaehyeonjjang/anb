package chat

type Room struct {
	Id       int
	Name     string
	Max      int
	Secret   bool
	Password string
	Admin    string
	Users    map[int64]*User
}

func (c *Room) GetUsers() map[int64]*User {
	return c.Users
}

func (c *Room) GetUsersCount() int {
	return len(c.Users)
}

func (c *Room) Join(user *User) bool {
	c.Users[user.Id] = user

	return true
}

func (c *Room) Exit(user *User) bool {
	delete(c.Users, user.Id)

	return true
}
