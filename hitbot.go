package hitbot

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/gorilla/http"
)

//Hitbot struct contains all required fields for a bot.
type Hitbot struct {
	Name    string
	servers []server
	connID  string
	token   string
}

type server struct {
	ServerIP string `json:"server_ip"`
}

//NewBot creates bot with specified name.
func NewBot(name string) Hitbot {
	log.Printf("%v - based on hitbot from Renerte", name)
	return Hitbot{Name: name}
}

//GetServers retrieves list of available servers.
func (bot *Hitbot) GetServers() {
	bot.servers = make([]server, 0, 5)
	buf := new(bytes.Buffer)
	if _, err := http.Get(buf, "http://api.hitbox.tv/chat/servers.json?redis=true"); err != nil {
		log.Fatalf("Could not get server list: %v", err)
	}
	if err := json.Unmarshal(buf.Bytes(), &bot.servers); err != nil {
		log.Fatalf("Could not parse JSON: %v", err)
	}
	log.Printf("Found %v servers", len(bot.servers))
}

//GetID tries to get connection id for the first server available.
func (bot *Hitbot) GetID() {
	buf := new(bytes.Buffer)
	for i := 0; i < len(bot.servers); i++ {
		if _, err := http.Get(buf, "http://"+bot.servers[i].ServerIP+"/socket.io/1"); err == nil {
			temp := strings.Split(buf.String(), ":")
			bot.connID = temp[0]
			log.Printf("Connection ID was found properly (%v)", temp[0])
			return
		}
	}
	log.Fatal("Could not get connection IDs!!!")
}
