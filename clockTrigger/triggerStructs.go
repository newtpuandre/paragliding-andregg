package clocktrigger

import "time"

type tickerStruct struct {
	TLatest    int64         `json:"t_latest"`
	TStart     int64         `json:"t_start"`
	TStop      int64         `json:"t_stop"`
	Tracks     []int         `json:"tracks"`
	Processing time.Duration `json:"processing"`
}

type discordMessage struct {
	Content string `json:"content"`
}
