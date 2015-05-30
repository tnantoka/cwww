package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const settingsJSON = "settings.json"

type Settings struct {
	RoomID   string
	APIToken string
}

func (settings Settings) Save() {
	data, err := json.Marshal(settings)
	if err != nil {
		log.Fatal(err)
	}
	path := storePath(settingsJSON)
	ioutil.WriteFile(path, data, 0644)
}

func (settings *Settings) Load() {
	path := storePath(settingsJSON)
	if fileExists(path) {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(data, &settings)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func NewSettings() Settings {
	settings := Settings{}
	settings.Load()
	return settings
}
