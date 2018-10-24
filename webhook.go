package main

type webhookStruct struct {
	WebhookURL      string `json:"webhookURL"`
	MinTriggerValue int    `json:"minTriggerValue"`
}

type discordMessage struct {
	Content string `json:"content"`
}
