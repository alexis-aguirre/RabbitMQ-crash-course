package dto

import (
	"net/http"
)

type ImageReport struct {
	ID       string `json:"id"`
	Location string `json:"location"`
	Image    string `json:"image"`
}

func (ir ImageReport) Validate() bool {
	resp, err := http.Get(ir.Image)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode >= 200 && resp.StatusCode < 300
}
