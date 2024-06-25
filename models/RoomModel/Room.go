package roommodel

import "time"

type Room struct {
	Id        int
	AppId     int
	ChannelId int
	CreatedBy int
	CreatedAt time.Time
	UpdatedAt time.Time
}
