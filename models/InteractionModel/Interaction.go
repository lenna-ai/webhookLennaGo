package interactionmodel

import "time"

type Interaction struct {
	Id           int
	RoomId       int
	Message      string
	UserableType string
	UserableId   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsNewUser    bool
	LastUserChat time.Time
	ChannelId    int
}
