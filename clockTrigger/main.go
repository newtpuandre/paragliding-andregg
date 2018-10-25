//Every 10 min we check if new tracks have been added.
//we send a message to a discord channel if there are new tracks

package clocktrigger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	for t := range time.NewTicker(30 * time.Second).C { //SET TO 10 MIN
		checkNewTracks(t)
	}
}

func checkNewTracks(t time.Time) {
	resp, err := http.Get("localhost:8080/paragliding/api/ticker/latest")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	var timeStamp int64
	decoder := json.NewDecoder(resp.Body)
	decoderr := decoder.Decode(&timeStamp)
	if decoderr != nil {
		fmt.Println(err)
	}

	fmt.Println(err)
}
