package webchatservice

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"webhooklenna/helpers"
	formatter "webhooklenna/models/Formatter"

	"github.com/gofiber/fiber/v2"
)

func (wwsi *WebhookWebchatServiceImpl) WebhookWebchat(c *fiber.Ctx, senderId int, integrationId int, event string, text string) (any, error) {
	integration, err := wwsi.webhookWebchatRepository.GetIntegrationById(integrationId)
	if err != nil {
		log.Println("wwsi.webhookWebchatRepository.GetIntegrationById(integrationId)")
		log.Println(err.Error())
		return integration, nil
	}

	user, err := wwsi.webhookWebchatRepository.GetUser(senderId)
	if err != nil {
		log.Println("wwsi.webhookWebchatRepository.GetUser(senderId)")
		log.Println(err.Error())
		return user, nil
	}

	room, err := wwsi.webhookWebchatRepository.GetRoomByCreatedBy(user.Id)
	if err != nil {
		if err.Error() == "room not found" {
			wwsi.webhookWebchatRepository.CreateRoomByCreatedBy(user.AppId, 5, user.Id)
		} else {
			log.Println("wwsi.webhookWebchatRepository.GetRoomByCreatedBy(user.Id)")
			log.Println(err.Error())
			return room, nil
		}
	}

	var resultMessageTypeText []interface{}

	messageTypeText := formatter.MessageTypeText{
		Type:   "text",
		Text:   text,
		Speech: text,
		Id:     0,
	}

	resultMessageTypeText = append(resultMessageTypeText, messageTypeText)

	_, err = wwsi.webhookWebchatRepository.CreateMessage("", room.Id, resultMessageTypeText, user)
	if err != nil {
		log.Println("wwsi.webhookWebchatRepository.CreateMessage('', room.Id, resultMessageTypeText, user)")
		log.Println(err.Error())
		return user, nil
	}

	interaction, err := wwsi.webhookWebchatRepository.CheckInteraction(room.Id, user)
	if err != nil {
		if err.Error() == "interaction not found" {
			return interaction, nil
		}
		return interaction, nil
	}
	password, err := wwsi.EncryptText("k384y0r4nTRIANGLE", "544:1719197008104@gmail.com")
	if err != nil {
		return password, err
	}
	backendDataApi := fiber.Map{
		"email":          user.Email,
		"password":       password,
		"fcm_token":      user.FcmToken,
		"client":         user.Client,
		"integration_id": user.IntegrationId,
		"success":        true,
	}

	backendServiceHitLoginChannel := make(chan map[string]any)
	go helpers.HitApiJsonWithoutAuth(c, "https://staging.lenna.ai/backend/api/35d7xy/auth/login", backendDataApi, backendServiceHitLoginChannel)
	backendServiceHitLoginResponse := <-backendServiceHitLoginChannel

	backendServiceHitLoginResponseByte, err := json.Marshal(backendServiceHitLoginResponse["body"])
	var dataResult map[string]any
	err = json.Unmarshal(backendServiceHitLoginResponseByte, &dataResult)
	if err != nil {
		panic(err.Error())
	}

	integrationAppId, err := wwsi.webhookWebchatRepository.GetIntegrationJoinByAppId(integration.AppId)
	if err != nil {
		fmt.Println("aa")
		panic(err.Error())
	}

	var Integration_IntegrationData map[string]interface{}
	err = json.Unmarshal([]byte((*integrationAppId)["integration_data"].(string)), &Integration_IntegrationData)
	if err != nil {
		panic(err.Error())
	}
	var botId int = int(Integration_IntegrationData["botId"].(float64))

	dataNlpResponse := map[string]string{
		"bot_type": "full_assistant",
		"bot_id":   strconv.Itoa(botId),
		"user_id":  strconv.Itoa(user.Id),
		"query":    text,
		"location": "0,0",
		"language": "id",
	}
	nlpResponseAsyncChannel := make(chan map[string]any)
	go helpers.HitApiBasicAuthJson(c, "http://34.126.152.255/nlp/chat", dataNlpResponse, nlpResponseAsyncChannel)

	nlpResponse := <-nlpResponseAsyncChannel

	nlpResponseAsyncChannelResponseByte, err := json.Marshal(nlpResponse["body"])
	var dataResultNlpResponse map[string]any
	err = json.Unmarshal(nlpResponseAsyncChannelResponseByte, &dataResultNlpResponse)
	if err != nil {
		panic(err.Error())
	}

	return dataResultNlpResponse, nil

}

func (wwsi *WebhookWebchatServiceImpl) EncryptText(key string, plaintext string) (string, error) {
	// uQnlZjZPyjlpe/JpHBoJ2NYOA+NplLm8vZbO65HBjPM=

	// var plaintext = "544:1719197008104@gmail.com"
	// var key = "k384y0r4nTRIANGLE"
	// var cipherString = "AES-256-CBC"
	iv := "0000000000000000"

	for i := len(key); i < 32; i++ {
		key = key + "0"
	}

	var plainTextBlock []byte
	length := len(plaintext)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	str := base64.StdEncoding.EncodeToString(ciphertext)

	return str, nil
}
