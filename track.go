package main

//Track stores information about one Track and stores it in DB
//Struct variables need to be with underscore inorder for the reflect magic to work..
type Track struct {
	ID            int     `json:"id"`
	Timestamp     int64   `json:"timestamp"`
	H_date        string  `json:"H_date"`
	Pilot         string  `json:"pilot"`
	Glider        string  `json:"glider"`
	Glider_id     string  `json:"glider_id"`
	Track_length  float64 `json:"track_length"`
	Track_src_url string  `json:"track_src_url"`
}

//TrackNoTimestamp is used to return when a request is given.
//It does not need some of the information that the Track one needs
type TrackNoTimestamp struct {
	H_date        string  `json:"H_date"`
	Pilot         string  `json:"pilot"`
	Glider        string  `json:"glider"`
	Glider_id     string  `json:"glider_id"`
	Track_length  float64 `json:"track_length"`
	Track_src_url string  `json:"track_src_url"`
}
