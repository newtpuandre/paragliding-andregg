//Every 10 min we check if new tracks have been added.
//we send a message to a discord channel if there are new tracks

//For the convenince of having ONE repository the clock trigger is located here.
//It is however a single executable located on openstack and have no affiliation with
//the api codebase

package clocktrigger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var hook = "https://discordapp.com/api/webhooks/506068687875342336/f1tbFV36A-c4JqRbSHH6Jdfbwce2EE_QHI2M0M1Z4m_nvKkRJlo7yWeJpb0bPJI7zj4S"

var latestTimeStamp = 0

func main() {
	for t := range time.NewTicker(10 * time.Minute).C {
		checkNewTracks(t)
	}
}

func checkNewTracks(t time.Time) {
	var timeStamp int

	resp, err := http.Get("https://igcviewer-andregg.herokuapp.com/paragliding/api/ticker/latest")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodyString := string(body)
	bodyString = strings.Trim(bodyString, "\n")
	timeStamp, err = strconv.Atoi(bodyString)

	if err != nil { //Either error in json format or no tracks in the API
		fmt.Println(err)
	}

	if timeStamp > latestTimeStamp {
		fmt.Println("New tracks!")
		tstruct := getIds(latestTimeStamp)
		sendMessage(timeStamp, tstruct)
	}

	latestTimeStamp = timeStamp
}

func getIds(timeStamp int) tickerStruct {
	newTime := strconv.Itoa(timeStamp)
	resp, err := http.Get("https://igcviewer-andregg.herokuapp.com/paragliding/api/ticker/" + newTime)

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	//Create a struct to hold info
	var tempStruct tickerStruct
	decoder := json.NewDecoder(resp.Body)
	decoderr := decoder.Decode(&tempStruct)
	if decoderr != nil {
		fmt.Println(err)
	}

	return tempStruct
}

func sendMessage(timeStamp int, t tickerStruct) {
	var message discordMessage
	message.Content = "Latest timestamp: "
	stringTimestamp := strconv.Itoa(timeStamp)
	message.Content += stringTimestamp
	message.Content += ". New track ids are: ["

	fmt.Println(len(t.Tracks))
	for i := range t.Tracks {
		trackID := strconv.Itoa(t.Tracks[i])

		if i != len(t.Tracks)-1 { //This is only for nice formatting
			message.Content = message.Content + trackID + ", "
		} else {
			message.Content = message.Content + "" + trackID
		}
	}

	message.Content += "]. Request took (" + t.Processing.String() + ")."

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(message)

	req, err := http.Post(hook, "application/json", b)
	if err != nil {
		fmt.Println(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusNoContent {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

}
