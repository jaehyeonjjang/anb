package chat

import (
	"encoding/json"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

type Option struct {
	Size  int    `json:"size"`
	Color string `json:"color"`
}

type Request struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Room     int    `json:"room"`
}

type LoginRequest struct {
	Id       int64  `json:"id"`
	Loginid  string `json:"loginid"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type MessageRequest struct {
	Nickname string `json:"nickname"`
	Room     int    `json:"room"`
	Text     string `json:"text"`
	Option   Option `json:"option"`
}

type MessageToRequest struct {
	Nickname string `json:"nickname"`
	Room     int    `json:"room"`
	Text     string `json:"text"`
	To       string `json:"to"`
	Option   Option `json:"option"`
}

type MakeRoomRequest struct {
	Name     string `json:"name"`
	Max      int    `json:"max"`
	Secret   bool   `json:"secret"`
	Password string `json:"password"`
}

type Response struct {
	Id      string `json:"id"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Room    int    `json:"room"`
}

type MessageResponse struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Room     int    `json:"room"`
	Text     string `json:"text"`
	Option   Option `json:"option"`
}

type Chat struct {
	Userlist *Userlist
	Roomlist *Roomlist
	Server   *socketio.Server
}

func (c *Chat) Handler() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("connected:", s.ID())
		s.Join("chat")

		return nil
	})

	server.OnEvent("/", "login", func(s socketio.Conn, msg string) string {
		log.Println("login", msg)
		return c.Login(s, msg)
	})

	server.OnEvent("/", "chat", func(s socketio.Conn, msg string) {
		log.Println("chat", msg)
		c.Chat(s, msg)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("disconnection")
		c.Disconnect(s)
	})

	go server.Serve()

	c.Server = server
}

func NewChat() *Chat {
	var item Chat
	item.Userlist = NewUserlist()

	item.Roomlist = NewRoomlist()
	item.Roomlist.Add("Wait Room", 100000, false, "", "")

	item.Handler()

	return &item
}

func (c *Chat) MakeMessage(id string, nickname string, text string, option *Option) string {
	var message MessageResponse
	message.Id = id
	message.Nickname = nickname
	message.Text = text
	if option != nil {
		message.Option = *option
	} else {
		message.Option.Size = 0
		message.Option.Color = ""
	}

	ret, _ := json.Marshal(message)
	return string(ret)
}

func (c *Chat) Broadcating(roomid int, code string, message string) {
	room := c.Roomlist.Find(roomid)
	if room == nil {
		return
	}

	users := room.GetUsers()

	for _, user := range users {
		user.Socket.Emit(code, message)
	}
}

func (c *Chat) BroadcatingWithout(id int64, roomid int, code string, message string) {
	room := c.Roomlist.Find(roomid)
	if room == nil {
		return
	}

	users := room.GetUsers()

	for _, user := range users {
		if user.Id == id {
			continue
		}

		user.Socket.Emit(code, message)
	}
}

func (c *Chat) Login(so socketio.Conn, msg string) string {
	id := so.ID()
	var response Response
	response.Id = id
	response.Room = 0

	var item LoginRequest
	err := json.Unmarshal([]byte(msg), &item)
	if err != nil {
		response.Status = 0
		response.Message = "아이디를 찾을수가 없습니다"
	} else {
		response.Status = 1

		find := c.Userlist.Find(item.Id)

		if find == nil {
			c.Userlist.Add(item.Id, item.Loginid, item.Nickname, so)
		}
	}

	ret, _ := json.Marshal(response)
	return string(ret)
}

func (c *Chat) Disconnect(so socketio.Conn) {
	id := so.ID()
	item := c.Userlist.FindBySocketio(id)

	if item != nil {
		log.Printf("room = %v\n", item.Room)
		room := c.Roomlist.Find(item.Room)
		if room != nil {
			room.Exit(item)
			c.Broadcating(item.Room, "exitroom", c.MakeMessage(id, item.Nickname, "", nil))

			if room.Id != 0 && room.GetUsersCount() == 0 {
				c.Roomlist.Remove(room.Id)
			}
		}

		c.Userlist.Remove(item.Id)
	}
}

func Error(so socketio.Conn, err error) {
	log.Println("error:", err)
}

func (c *Chat) JoinRoom(so socketio.Conn, msg string) string {
	id := so.ID()
	item := c.Userlist.FindBySocketio(id)

	if item == nil {
		return ""
	}

	var req Request
	err := json.Unmarshal([]byte(msg), &req)

	if err != nil {
		return ""
	}

	if item.Room == req.Room {
		return ""
	}

	room := c.Roomlist.Find(item.Room)
	if room != nil {
		room.Exit(item)
		c.Broadcating(item.Room, "exitroom", c.MakeMessage(id, item.Nickname, "", nil))
	}

	room = c.Roomlist.Find(req.Room)
	if room != nil && room.Join(item) == true {
		c.Broadcating(req.Room, "joinroom", c.MakeMessage(id, item.Nickname, "", nil))
	}

	return ""
}

func (c *Chat) ExitRoom(so socketio.Conn, msg string) string {
	id := so.ID()
	item := c.Userlist.FindBySocketio(id)

	if item == nil {
		return ""
	}

	var req Request
	err := json.Unmarshal([]byte(msg), &req)

	if err != nil {
		return ""
	}

	if item.Room == req.Room {
		return ""
	}

	room := c.Roomlist.Find(req.Room)
	if room != nil {
		room.Exit(item)
		c.Broadcating(req.Room, "exitroom", c.MakeMessage(id, item.Nickname, "", nil))
	}

	room = c.Roomlist.Find(0)
	if room != nil && room.Join(item) == true {
		c.Broadcating(0, "joinroom", c.MakeMessage(id, item.Nickname, "", nil))
	}

	return ""
}

func (c *Chat) ListRoom(so socketio.Conn, msg string) string {
	rooms := c.Roomlist.GetList()
	e, err := json.Marshal(rooms)
	if err != nil {
		return ""
	}

	return string(e)
}

func (c *Chat) MakeRoom(so socketio.Conn, msg string) string {
	id := so.ID()
	item := c.Userlist.FindBySocketio(id)

	if item == nil {
		return ""
	}

	var req MakeRoomRequest
	err := json.Unmarshal([]byte(msg), &req)

	if err != nil {
		return ""
	}

	roomId := c.Roomlist.Add(req.Name, req.Max, req.Secret, req.Password, id)

	room := c.Roomlist.Find(item.Room)
	if room != nil {
		room.Exit(item)
		c.Broadcating(item.Room, "exitroom", c.MakeMessage(id, item.Nickname, "", nil))
	}

	room = c.Roomlist.Find(roomId)
	if room != nil && room.Join(item) == true {
	}

	return ""
}

func (c *Chat) EditRoom(so socketio.Conn, msg string) string {
	return ""
}

func (c *Chat) Ban(so socketio.Conn, msg string) string {
	return ""
}

func (c *Chat) Kick(so socketio.Conn, msg string) string {
	return ""
}

func (c *Chat) Chat(so socketio.Conn, msg string) string {
	id := so.ID()
	item := c.Userlist.FindBySocketio(id)

	if item == nil {
		return ""
	}

	var req MessageRequest
	err := json.Unmarshal([]byte(msg), &req)

	if err != nil {
		return ""
	}

	c.Broadcating(item.Room, "chat", c.MakeMessage(id, item.Nickname, req.Text, &req.Option))

	return ""
}
