package webchatrepository

import (
	integrationmodel "webhooklenna/models/IntegrationModel"
	interactionmodel "webhooklenna/models/InteractionModel"
	messagemodel "webhooklenna/models/MessageModel"
	roommodel "webhooklenna/models/RoomModel"
	usermodel "webhooklenna/models/UserModel"

	"gorm.io/gorm"
)

type WebhookWebchatRepository interface {
	GetIntegrationById(id int) (*integrationmodel.Integration, error)
	GetIntegrationJoinByAppId(appId int) (*map[string]any, error)
	GetUser(id int) (*usermodel.User, error)
	GetRoomByCreatedBy(user_id int) (*roommodel.Room, error)
	CreateRoomByCreatedBy(app_id int, channel_id int, user_id int) (*roommodel.Room, error)
	CreateMessage(channel_data string, room_id int, content any, user *usermodel.User) (*messagemodel.Message, error)
	CheckInteraction(room_id int, user *usermodel.User) (*[]interactionmodel.Interaction, error)
}

type WebhookWebchatRepositoryImpl struct {
	DB *gorm.DB
}

func NewWebhookWebchatRepository(db *gorm.DB) *WebhookWebchatRepositoryImpl {
	return &WebhookWebchatRepositoryImpl{
		DB: db,
	}
}
