package main

import "github.com/globalsign/mgo/bson"

//webhookStruct is used to insert and retrive data from DB
type webhookStruct struct {
	ID              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	WebhookID       int           `json:"webhookid"`
	WebhookURL      string        `json:"webhookURL"`
	MinTriggerValue int           `json:"minTriggerValue"`
	NewTracks       int           `json:"newTracks"`
	LastTrackID     int           `json:"lastTrackID"`
}

//webhookStructResponse is used when a user is posting or getting a response
type webhookStructResponse struct {
	WebhookURL      string `json:"webhookURL"`
	MinTriggerValue int    `json:"minTriggerValue"`
}

//discordMessage struct is how discord want formatting to be
type discordMessage struct {
	Content string `json:"content"`
}
