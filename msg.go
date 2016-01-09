package hitbot

// imports

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

//JoinChannel joins channel specified in the parameter.
func (bot *Hitbot) JoinChannel(channel string) {
	message := map[string]interface{}{
		"name": "message",
		"args": []map[string]interface{}{
			{
				"method": "joinChannel",
				"params": map[string]interface{}{
					"channel": channel,
					"name":    bot.Name,
					"token":   bot.auth.Token,
					"isAdmin": false}}}}
	var js []byte
	js, _ = json.Marshal(message)
	msg := "5:::" + string(js)
	bot.conn.WriteMessage(websocket.TextMessage, []byte(msg))
	log.Print("Login sent!")
}

//MessageHandler processes messages recieved from chat server.
func (bot *Hitbot) MessageHandler() {
	for {
		_, p, err := bot.conn.ReadMessage()
		if err != nil {
			return
		}
		//log.Printf("Message: %v", string(p)) //debug info
		if string(p[:3]) == "2::" {
			bot.conn.WriteMessage(websocket.TextMessage, []byte("2::"))
			log.Print("Ping!")
			continue
		} else if string(p[:3]) == "1::" {
			log.Print("Server confirmed connection \\o/")
			bot.JoinChannel(bot.channels[0])
			continue
		}
	}
}

//Connect starts connection to active server, and stores its pointer in Hitbot struct.
func (bot *Hitbot) Connect(channels ...string) {
	dialer := websocket.Dialer{}
	c, _, err := dialer.Dial("ws://"+bot.servers[bot.activeServer].ServerIP+"/socket.io/1/websocket/"+bot.connID, nil)
	if err != nil {
		log.Fatal(err)
	}
	bot.conn = c
	bot.channels = channels
}
