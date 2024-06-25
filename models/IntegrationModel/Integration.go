package integrationmodel

import "time"

type Integration struct {
	Id              int
	AppId           int
	ChannelId       int
	IntegrationData string
	Status          bool
	DeletedAt       time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	IsPrimary       bool
	BotId           int
	BusinessUnitId  int
	GroupId         int
	ForwardTo       int
	IsForward       bool
	ForwardUrl      string
	ListenMention   bool
	BotType         string
}
