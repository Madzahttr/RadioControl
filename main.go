package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

/*
currentSong: {
	name: "Faded",
	artist: "Alan Walker",
	album: "Faded",
	releaseYear: "2015",
	genre: "EDM",
	file: "https://dl.madzahttr.com/Faded.mp3"
},
*/

type Song struct {
	ID          string `json:"_id"`
	Rev         string `json:"_rev,omitempty"`
	name        string `json:"name"`
	artist      string `json:"artist"`
	album       string `json:"album"`
	releaseYear string `json:"releaseYear`
	genre       string `json:"genre"`
	file        string `json:"file"`
}

func main() {
	/*client, err := kivik.New(context.GODO(), "couch", "http://Admin:abc123@localhost:5984/")

	song := Song{
		ID:"thisisanid", 
		name: "Faded",
		artist: "Alan Walker",
		album: "Faded",
		releaseYear: "2015",
		genre: "EDM",
		file: "C:\Users\Ethan\Music\Radio\AlanWalker"
	}

	rev, err := db.Put(context.TODO(), "songs", song)
	if err != nil {
		panic(err)
	}*/

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Resizable: true,
		Title:     "RadioControl",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
	})
	app.Run()
}
