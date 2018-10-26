package main

import "github.com/globalsign/mgo/bson"

type webhookStruct struct {
	ID              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	WebhookURL      string        `json:"webhookURL"`
	MinTriggerValue int           `json:"minTriggerValue"`
	NewTracks       int           `json:"newTracks"`
}

type webhookStructResponse struct {
	WebhookURL      string `json:"webhookURL"`
	MinTriggerValue int    `json:"minTriggerValue"`
}

type discordMessage struct {
	Content string `json:"content"`
}
