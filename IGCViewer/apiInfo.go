package IGCViewer

import (
	"github.com/rickb777/date/period"
)

type apiInfo struct {
	Uptime  period.Period `json:"uptime"`
	Info    string        `json:"info"`
	Version string        `json:"version"`
}
