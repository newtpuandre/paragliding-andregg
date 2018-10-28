package main

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding

//DBInfo contains information about database connection and collections
type DBInfo struct {
	ConnectionString        string
	DBString                string
	TrackCollectionString   string
	WebhookCollectionString string
}

//Credentials is a global variable with currently set DBInfo
var Credentials DBInfo

//Initialize database credentials
func dbInit() {
	Credentials.TrackCollectionString = "tracks"
	Credentials.WebhookCollectionString = "webhooks"
	Credentials.DBString = "paragliding"
	Credentials.ConnectionString = "mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding"
}

//Inserts a track into the track collection
func insertTrack(t *Track, db *DBInfo) {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB(db.DBString).C(db.TrackCollectionString).Insert(t)
	if err != nil {
		fmt.Println(err)
	}

}

//Count all tracks in the track collection
func countTrack(db *DBInfo) int {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	count, err := session.DB(db.DBString).C(db.TrackCollectionString).Count()
	if err != nil {
		fmt.Println(err)
		return -1
	}

	return count
}

//Gets all tracks from collection and return an array
func getAllTracks(db *DBInfo) []Track {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	var all []Track

	err = session.DB(db.DBString).C(db.TrackCollectionString).Find(bson.M{}).All(&all)
	if err != nil {
		panic(err)
	}
	return all
}

//Get ids from database and update global variables
func updateIDFromDB(db *DBInfo) {
	count := countTrack(db)
	for i := 0; i < count; i++ {
		trackID = append(trackID, lastID)
		lastID++
	}

	hooks := getWebHooks(db)
	for i := 0; i < len(hooks); i++ {
		webhookID = append(webhookID, hooks[i].WebhookID)
	}

	if len(hooks) > 0 {
		lastWebhookID = webhookID[len(webhookID)-1] + 1
	} else {
		lastWebhookID = 0
	}

}

//Deletes everything in the database
func deleteTrackCollection(db *DBInfo) int {
	count := countTrack(db) //Get database count

	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB(db.DBString).C(db.TrackCollectionString).DropCollection()
	if err != nil {
		panic(err)
	}

	return count
}

//Gets all webhooks from the collection and returns an array
func getWebHooks(db *DBInfo) []webhookStruct {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	var all []webhookStruct

	err = session.DB(db.DBString).C(db.WebhookCollectionString).Find(bson.M{}).All(&all)
	if err != nil {
		panic(err)
	}
	return all
}

//Counts webhooks from the collection and returns an int
//Currently unused. Future proofing.
func countWebhook(db *DBInfo) int {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	count, err := session.DB(db.DBString).C(db.WebhookCollectionString).Count()
	if err != nil {
		fmt.Println(err)
		return -1
	}

	return count
}

//Insert a webhook into the database
func insertWebhook(w *webhookStruct, db *DBInfo) {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DBString).C(db.WebhookCollectionString).Insert(w)
	if err != nil {
		panic(err)
	}

}

//Updates a webhook based on MongoDB id
func updateWebhook(w *webhookStruct, db *DBInfo) {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB(db.DBString).C(db.WebhookCollectionString).Update(bson.M{"_id": w.ID}, w)
	if err != nil {
		panic(err)
	}
}

//Deletes a webhook based on mongoDB Id
func deleteWebhook(w *webhookStruct, db *DBInfo) {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	err = session.DB(db.DBString).C(db.WebhookCollectionString).Remove(bson.M{"_id": w.ID})
	if err != nil {
		panic(err)
	}
}

//clearWebhookID clears all last used ids in the webhook DB
//inorder for the webhook post to work correctly
func clearWebhookID(db *DBInfo) {
	hooks := getWebHooks(db)

	for i := range hooks {
		hooks[i].LastTrackID = 0
		hooks[i].NewTracks = 0
		updateWebhook(&hooks[i], db)
	}

}
