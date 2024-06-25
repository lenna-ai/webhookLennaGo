package controllers

import webchatcontroller "webhooklenna/controllers/webhookController/webchatController"

type AllController struct {
	WebhookWebchatController webchatcontroller.WebhookWebchatController
}

func NewAllController(webhookWebchatController webchatcontroller.WebhookWebchatController) *AllController {
	return &AllController{
		WebhookWebchatController: webhookWebchatController,
	}
}
