package formattermessagemodel

// senderId = 1524940
// integrationId = 1098
// events = "message"
// app_id = 544
type CreateMessageModelFormatterRequest struct {
	SenderId      int    `validate:"required,numeric,gt=0"`
	IntegrationId int    `validate:"required,numeric,gt=0"`
	Events        string `validate:"required,gt=0"`
	AppId         int    `json:"app_id" validate:"required,numeric,gt=0"`
}
