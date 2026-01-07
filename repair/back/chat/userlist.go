package chat

import socketio "github.com/googollee/go-socket.io"

type Userlist struct {
	Items map[int64]*User
}

var _users Userlist

func NewUserlist() *Userlist {
	var userlist Userlist
	userlist.Items = make(map[int64]*User)

	return &userlist
}

func (c *Userlist) Add(id int64, loginid string, nickname string, so socketio.Conn) *User {
	if item, ok := c.Items[id]; ok {
		return item
	}

	var user User
	user.Id = id
	user.Loginid = loginid
	user.Nickname = nickname
	user.Status = WAIT
	user.Room = 0
	user.Socket = so
	c.Items[id] = &user

	return &user
}

func (c *Userlist) Remove(id int64) {
	if _, ok := c.Items[id]; ok {
		delete(c.Items, id)
	}
}

func (c *Userlist) Find(id int64) *User {
	if user, ok := c.Items[id]; ok {
		return user
	}

	return nil
}

func (c *Userlist) FindBySocketio(id string) *User {
	for _, user := range c.Items {
		if user.Socket.ID() == id {
			return user
		}
	}

	return nil
}

func (c *Userlist) FindByLoginid(loginid string) *User {
	for _, user := range c.Items {
		if user.Loginid == loginid {
			return user
		}
	}

	return nil
}
