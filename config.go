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
	Commands  []command `json:"commands"`
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
	bot := NewBot(c.Name)
	bot.Verbose(verbose)
	bot.GetServers()
	bot.GetID()
	bot.Auth(c.Pass)
	bot.RegisterBuiltinHandlers()
	channels := make([]string, 64)
	for _, channel := range c.Channels {
		channels = append(channels, channel.Name)
		bot.chCmds[channel.Name] = chCmd{cmds: make(map[string]HandlerFunc), cmdHandlers: make(map[string]cmd)}
		for _, comm := range channel.Commands {
			bot.chCmds[channel.Name].cmdHandlers[comm.Name] = cmd{Handler: comm.Handler, Role: comm.Role, Data: comm.Data}
		}
	}
	for _, comm := range c.Commands {
		bot.cmdHandlers[comm.Name] = cmd{Handler: comm.Handler, Role: comm.Role, Data: comm.Data}
	}
	bot.Connect(channels...)
	bot.NameColor(c.NameColor)
	return bot
}

//RegisterBuiltinHandlers registers all builtin handlers in bot instance for commands to use.
func (bot *Hitbot) RegisterBuiltinHandlers() {
	bot.registerHandler("basic", basicInit)
	if bot.verbose {
		log.Print("Registered builtin handlers")
	}
}

//LoadCommands loads commands from map created by either LoadBot, or RegisterCommand functions.
func (bot *Hitbot) LoadCommands() {
	for name, cmd := range bot.cmdHandlers {
		bot.cmds[name] = bot.handlers[cmd.Handler](cmd.Data)
	}
	for _, ch := range bot.chCmds {
		for name, cmd := range ch.cmdHandlers {
			ch.cmds[name] = bot.handlers[cmd.Handler](cmd.Data)
		}
	}
}
