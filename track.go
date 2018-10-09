package main

type Track struct {
	H_date       string  `json:"H_date"`
	Pilot        string  `json:"pilot"`
	Glider       string  `json:"glider"`
	Glider_id    string  `json:"glider_id"`
	Track_length float64 `json:"track_length"`
}
