package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/boyzhujian/golib/ciscospark"
)

type Config struct {
	Name        string
	Accesstoken string
}

func main() {

	var conf Config
	//for secruity ,I put every thing not public in secret folder which is ignored by .gitignore file,you must manually constuct toml file to run this example.
	if _, err := toml.DecodeFile("../secret/ciscosparksecret.toml", &conf); err != nil {
		// handle error
	}

	b := ciscospark.Bot{Name: conf.Name, Accesstoken: conf.Accesstoken}
	roomlist := b.ListRooms()
	fmt.Println(roomlist)
	for _, room := range roomlist.Rooms {
		fmt.Println(room.ID, room.Title)
	}
}
