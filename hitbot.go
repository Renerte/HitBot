package hitbot

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/gorilla/http"
	"github.com/gorilla/websocket"
)

//Hitbot struct contains all required fields for a bot.
type Hitbot struct {
	Name         string
	servers      []server
	activeServer int
	connID       string
	conn         *websocket.Conn
	auth         auth
	channels     []string
	cmdHandlers  map[string]CmdHandler
	color        string
}

type server struct {
	ServerIP string `json:"server_ip"`
}

type auth struct {
	Token string `json:"authToken"`
}

//NewBot creates bot with specified name.
func NewBot(name string) Hitbot {
	log.Printf("%v - based on hitbot made by Renerte (github.com/Renerte)", name)
	return Hitbot{Name: name, activeServer: -1, cmdHandlers: make(map[string]CmdHandler), color: "ffffff"}
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
	for i := bot.activeServer + 1; i < len(bot.servers); i++ {
		if _, err := http.Get(buf, "http://"+bot.servers[i].ServerIP+"/socket.io/1"); err == nil {
			temp := strings.Split(buf.String(), ":")
			bot.connID = temp[0]
			bot.activeServer = i
			log.Print("Connection ID was found properly")
			return
		}
	}
	log.Fatal("Could not get connection IDs!!!")
}

//Auth attempts to authenticate with Hitbox.tv to get access token, which is needed for chat connection.
func (bot *Hitbot) Auth(pass string) {
	temp := "login=" + bot.Name + "&pass=" + pass
	body := strings.NewReader(temp)
	headers := map[string][]string{"Content-Type": []string{"application/x-www-form-urlencoded"}}
	st, _, r, err := http.DefaultClient.Post("http://api.hitbox.tv/auth/token", headers, body)
	if err != nil {
		log.Fatal(err)
	}
	if r != nil {
		defer r.Close()
	}
	res := make([]byte, 56)
	r.Read(res)
	if err := json.Unmarshal(res, &bot.auth); err != nil {
		log.Fatalf("Could not parse JSON: %v", err)
	}
	if st.Code != 200 {
		log.Fatalf("Authentication failed! (status %v)", st.Code)
	}
	log.Print("Successfully authenticated with Hitbox.tv")
}

//NameColor sets color for bot name. (hex value without `#`)
func (bot *Hitbot) NameColor(color string) {
	bot.color = color
}
