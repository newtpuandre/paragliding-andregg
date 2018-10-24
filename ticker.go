package main

type tickerStruct struct {
	tLatest    string  `json:"t_latest"`
	tStart     string  `json:"t_start"`
	tStop      string  `json:"t_stop"`
	tracks     []Track `json:"tracks"`
	processing string  `json:"processing"`
}
