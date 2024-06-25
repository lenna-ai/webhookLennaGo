package messagemodel

import "time"

type Message struct {
	Id               int
	Channel_data     string
	RoomId           int
	MessageableType  string
	MessageableId    int
	Content          string
	Data             string
	DeletedAt        time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Encrypt          bool
	IsDiscussion     bool
	TicketId         int
	UserId           int
	ChannelMessageId int
}
