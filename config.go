package hitbot

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Name      string    `json:"name"`
	Pass      string    `json:"pass"`
	NameColor string    `json:"nameColor"`
	Channels  []channel `json:"channels"`
}

type channel struct {
	Name     string    `json:"name"`
	Commands []command `json:"commands"`
}

type command struct {
	Name    string      `json:"name"`
	Handler string      `json:"handler"`
	Role    string      `json:"role"`
	Data    interface{} `json:"data"`
}

//LoadBot creates bot instance from config file.
func LoadBot(path string, verbose bool) Hitbot {
	raw := new(bytes.Buffer)
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Fatal error while opening config: %v", err)
	}
	_, err = raw.ReadFrom(f)
	if err != nil {
		log.Fatalf("Fatal error while reading config: %v", err)
	}
	var c config
	if err := json.Unmarshal(raw.Bytes(), &c); err != nil {
		log.Fatalf("Fatal error while processing config: %v", err)
	}
	bot := Hitbot{name: c.Name, verbose: verbose}
	bot.GetServers()
	bot.GetID()
	bot.Auth(c.Pass)
	channels := make([]string, 64)
	for _, channel := range c.Channels {
		channels = append(channels, channel.Name)
	}
	bot.Connect(channels...)
	bot.NameColor(c.NameColor)
	return bot
}
