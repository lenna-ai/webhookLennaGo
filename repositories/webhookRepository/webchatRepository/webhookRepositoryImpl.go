package webchatrepository

import (
	"errors"
	"log"
	"time"
	integrationmodel "webhooklenna/models/IntegrationModel"
	interactionmodel "webhooklenna/models/InteractionModel"
	messagemodel "webhooklenna/models/MessageModel"
	roommodel "webhooklenna/models/RoomModel"
	usermodel "webhooklenna/models/UserModel"
)

func (wwri *WebhookWebchatRepositoryImpl) GetIntegrationById(id int) (*integrationmodel.Integration, error) {
	integration := integrationmodel.Integration{}
	if err := wwri.DB.Take(&integration, "id = ?", id).Error; err != nil {
		log.Println(err.Error())
		return &integration, err
	}

	return &integration, nil
}

func (wwri *WebhookWebchatRepositoryImpl) GetIntegrationJoinByAppId(appId int) (*map[string]any, error) {
	resultDataIntegrationChannel := new(map[string]any)
	// 1098
	if err := wwri.DB.Model(&integrationmodel.Integration{}).Select("*").Joins("inner join omnichannel.channels on omnichannel.integrations.channel_id = omnichannel.channels.id").Where("omnichannel.integrations.app_id = ? and omnichannel.channels.name = 'bot' and omnichannel.integrations.status = true", appId).Scan(resultDataIntegrationChannel).Error; err != nil {
		log.Println(err.Error())
		return resultDataIntegrationChannel, err
	}

	return resultDataIntegrationChannel, nil
}

func (wwri *WebhookWebchatRepositoryImpl) GetUser(id int) (*usermodel.User, error) {
	user := new(usermodel.User)
	if err := wwri.DB.Take(&user, "id = ?", id).Error; err != nil {
		log.Println(err.Error())
		return user, err
	}

	return user, nil
}

func (wwri *WebhookWebchatRepositoryImpl) GetRoomByCreatedBy(user_id int) (*roommodel.Room, error) {
	room := new(roommodel.Room)
	roomSql := wwri.DB.Take(&room, "created_by = ?", user_id)
	if roomSql.Error != nil {
		log.Println(roomSql.Error.Error())
		return room, roomSql.Error
	}

	if roomSql.RowsAffected < 1 {
		return room, errors.New("room not found")
	}

	return room, nil
}

func (wwri *WebhookWebchatRepositoryImpl) CreateRoomByCreatedBy(app_id int, channel_id int, user_id int) (*roommodel.Room, error) {
	var room = &roommodel.Room{
		AppId:     app_id,
		ChannelId: channel_id,
		CreatedBy: user_id,
	}
	if err := wwri.DB.Create(room).Error; err != nil {
		return room, err

	}
	return room, nil
}

func (wwri *WebhookWebchatRepositoryImpl) CreateMessage(channel_data string, room_id int, content any, user *usermodel.User) (*messagemodel.Message, error) {
	var message = messagemodel.Message{
		RoomId:          room_id,
		MessageableType: "user",
		MessageableId:   user.Id,
	}
	if err := wwri.DB.Create(&message).Error; err != nil {
		return &message, nil
	}
	return &message, nil
}

func (wwri *WebhookWebchatRepositoryImpl) CheckInteraction(room_id int, user *usermodel.User) (*[]interactionmodel.Interaction, error) {
	var interaction = new([]interactionmodel.Interaction)
	var currentTime = time.Now()
	dateTime := currentTime.Format("2006-01-02")
	err := wwri.DB.Where("userable_id = ? and room_id = ? and date(created_at) = ?", user.Id, room_id, dateTime).Find(&interaction)
	if err.Error != nil {
		return interaction, err.Error
	} else if err.RowsAffected < 1 {
		return interaction, errors.New("interaction not found")
	}

	return interaction, err.Error
}
