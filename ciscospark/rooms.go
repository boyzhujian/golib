package ciscospark

import (
	"encoding/json"
	"time"

	"github.com/parnurzeal/gorequest"
)

//Roomlist is api return json
type Roomlist struct {
	Rooms []struct {
		ID           string    `json:"id"`
		Title        string    `json:"title"`
		Type         string    `json:"type"`
		IsLocked     bool      `json:"isLocked"`
		TeamID       string    `json:"teamId"`
		LastActivity time.Time `json:"lastActivity"`
		Created      time.Time `json:"created"`
	} `json:"items"`
}

//ListRooms  reference   https://developer.ciscospark.com/endpoint-rooms-get.html
//https://api.ciscospark.com/v1/rooms
func (u Bot) ListRooms() (rooms Roomlist) {
	var roomlist Roomlist
	r := gorequest.New()
	resp, body, _ := r.Get("https://api.ciscospark.com/v1/rooms").Set("Authorization", "Bearer "+u.Accesstoken).End()
	if resp.StatusCode == 200 {
		json.Unmarshal([]byte(body), &roomlist)
	}
	return roomlist
}
