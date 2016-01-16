package hitbot

import (
	"log"
	"strings"
)

//RegisterHandler registers specified handler for command of provided name. Will overwrite if command already exists.
func (bot *Hitbot) RegisterHandler(name string, handler func(map[string]interface{}) (string, string)) {
	bot.cmdHandlers[name] = handler
	log.Printf("Registered '%v' command", name)
}

type basicCmd struct {
	response string
}

func (cmd basicCmd) Handle(params map[string]interface{}) (string, string) {
	return params["channel"].(string), cmd.response
}

//BasicCmd creates and registers basic cmd handler.
func (bot *Hitbot) BasicCmd(name string, response string) {
	bot.RegisterHandler(name, basicCmd{response: response}.Handle)
}

func (bot *Hitbot) dispatchCommand(params map[string]interface{}) {
	cmd := strings.Split(params["text"].(string), " ")
	//log.Printf("%v invoked command '%v'", params["name"].(string), cmd[0][1:]) //debug stuff
	if handler, prs := bot.cmdHandlers[cmd[0][1:]]; prs {
		bot.sendMessage(handler(params))
	}
}
