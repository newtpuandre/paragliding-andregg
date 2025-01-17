package main

import (
	"testing"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Used for tests
func TestDbInit(t *testing.T) {
	Credentials.TrackCollectionString = "test_tracks"
	Credentials.WebhookCollectionString = "test_webhooks"
	Credentials.DBString = "paragliding"
	Credentials.ConnectionString = "mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding"

}

//setupDB returns an DBInfo struct with filled variables
func setupDB(t *testing.T) *DBInfo {
	var testDB DBInfo
	testDB.TrackCollectionString = "test_tracks"
	testDB.WebhookCollectionString = "test_webhooks"
	testDB.DBString = "paragliding"
	testDB.ConnectionString = "mongodb://newtpu:database1@ds239903.mlab.com:39903/paragliding"

	return &testDB
}

//Clears track collection
func clearTrackCol(t *testing.T, db *DBInfo) {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		t.Error(err)
	}
	defer session.Close()

	err = session.DB(db.DBString).C(db.TrackCollectionString).DropCollection()
	if err != nil {
		t.Error(err)
	}

}

//Clears Webhook Collection
func clearHookCol(t *testing.T, db *DBInfo) {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		t.Error(err)
	}
	defer session.Close()
	err = session.DB(db.DBString).C(db.WebhookCollectionString).DropCollection()
	if err != nil {
		t.Error(err)
	}
}

//Tests the InsertTrack function
func TestInsertTrack(t *testing.T) {
	testDB := setupDB(t)
	defer clearTrackCol(t, testDB)

	var newTrack Track
	newTrack.H_date = "2016-02-19 00:00:00 +0000 UTC"
	newTrack.Pilot = "Miguel Angel Gordillo"
	newTrack.Glider = "RV8"
	newTrack.Glider_id = "EC-XLL"
	newTrack.Track_length = 443.2573603705269
	newTrack.Track_src_url = "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"

	insertTrack(&newTrack, testDB)

	session, err := mgo.Dial(testDB.ConnectionString)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	resultTrack := Track{}
	err = session.DB(testDB.DBString).C(testDB.TrackCollectionString).Find(newTrack).One(&resultTrack)

	if err != nil {
		t.Errorf("error in FindId(): %v", err.Error())
		return
	}

	if countTrack(testDB) != 1 {
		t.Error("adding new track failed.")
	}

}

//Tests the CountTrack function
func TestCountTrack(t *testing.T) {
	testDB := setupDB(t)
	defer clearTrackCol(t, testDB)

	var newTrack Track
	newTrack.H_date = "2016-02-19 00:00:00 +0000 UTC"
	newTrack.Pilot = "Miguel Angel Gordillo"
	newTrack.Glider = "RV8"
	newTrack.Glider_id = "EC-XLL"
	newTrack.Track_length = 443.2573603705269
	newTrack.Track_src_url = "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"

	insertTrack(&newTrack, testDB)

	if countTrack(testDB) != 1 {
		t.Error("Count is not equal to one")
	}

}

//Tests the GetAllTracks Function
func TestGetAllTracks(t *testing.T) {
	testDB := setupDB(t)
	defer clearTrackCol(t, testDB)

	var newTrack Track
	newTrack.H_date = "2016-02-19 00:00:00 +0000 UTC"
	newTrack.Pilot = "Miguel Angel Gordillo"
	newTrack.Glider = "RV8"
	newTrack.Glider_id = "EC-XLL"
	newTrack.Track_length = 443.2573603705269
	newTrack.Track_src_url = "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"

	//Insert two tracks
	insertTrack(&newTrack, testDB)
	insertTrack(&newTrack, testDB)

	tracks := getAllTracks(testDB)

	if len(tracks) != 2 {
		t.Fatal("Could not retrieve two tracks")
	}

}

//Tests the DeleteTrackCollection Function
func TestDeleteTrackCollection(t *testing.T) {
	testDB := setupDB(t)

	var newTrack Track
	newTrack.H_date = "2016-02-19 00:00:00 +0000 UTC"
	newTrack.Pilot = "Miguel Angel Gordillo"
	newTrack.Glider = "RV8"
	newTrack.Glider_id = "EC-XLL"
	newTrack.Track_length = 443.2573603705269
	newTrack.Track_src_url = "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"

	//Insert two tracks
	insertTrack(&newTrack, testDB)
	insertTrack(&newTrack, testDB)

	count := deleteTrackCollection(testDB)

	if count != 2 {
		t.Fatal("Could not delete two objects")
	}
}

//Tests the GetWebhooks Function
func TestGetWebhooks(t *testing.T) {
	testDB := setupDB(t)
	defer clearHookCol(t, testDB)

	var hook webhookStruct
	hook.WebhookURL = "test.test/test"
	hook.MinTriggerValue = 1

	//Insert two tracks
	insertWebhook(&hook, testDB)
	insertWebhook(&hook, testDB)

	hooks := getWebHooks(testDB)

	if len(hooks) != 2 {
		t.Fatal("Could not retrieve two Webhooks")
	}

}

//Tests the CountWebhook Function
func TestCountWebhook(t *testing.T) {
	testDB := setupDB(t)
	defer clearHookCol(t, testDB)

	var hook webhookStruct
	hook.WebhookURL = "test.test/test"
	hook.MinTriggerValue = 1

	insertWebhook(&hook, testDB)

	if countWebhook(testDB) != 1 {
		t.Error("Count is not equal to one")
	}

}

//Tests the InsertWebhook Function
func TestInsertWebhook(t *testing.T) {
	testDB := setupDB(t)
	defer clearHookCol(t, testDB)

	var hook webhookStruct
	hook.WebhookURL = "test.test/test"
	hook.MinTriggerValue = 1

	insertWebhook(&hook, testDB)

	session, err := mgo.Dial(testDB.ConnectionString)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	resultHook := webhookStruct{}
	err = session.DB(testDB.DBString).C(testDB.WebhookCollectionString).Find(hook).One(&resultHook)

	if err != nil {
		t.Errorf("error in FindId(): %v", err.Error())
		return
	}

	if countWebhook(testDB) != 1 {
		t.Error("adding new track failed.")
	}

}

//Tests the UpdateWebhook Function
func TestUpdateWebhook(t *testing.T) {
	testDB := setupDB(t)
	defer clearHookCol(t, testDB)

	var hook webhookStruct
	hook.ID = bson.NewObjectId()
	hook.WebhookURL = "test.test/test"
	hook.MinTriggerValue = 1

	insertWebhook(&hook, testDB)

	hook.WebhookURL = "nottest.test/nottest"

	updateWebhook(&hook, testDB)

	hooks := getWebHooks(testDB)

	if hooks[0].WebhookURL != "nottest.test/nottest" {
		t.Fatal("Object was not updated")
	}

}

func TestDeleteWebhook(t *testing.T) {
	testDB := setupDB(t)

	var hook webhookStruct
	hook.ID = bson.NewObjectId()
	hook.WebhookURL = "test.test/test"
	hook.MinTriggerValue = 1

	insertWebhook(&hook, testDB)

	deleteWebhook(&hook, testDB)

}
