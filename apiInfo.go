package main

import (
	"github.com/rickb777/date/period"
)

//Struct that contains info about the api
type apiInfo struct {
	Uptime  period.Period `json:"uptime"`
	Info    string        `json:"info"`
	Version string        `json:"version"`
}
