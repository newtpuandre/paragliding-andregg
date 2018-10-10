package main

//URL is used when a user posts a JSON object with an URL
type URL struct {
	URL string `json:"Url"`
}

//URLID is used to return a JSON object with a new ID to the User
type URLID struct {
	ID int `json:"id"`
}
