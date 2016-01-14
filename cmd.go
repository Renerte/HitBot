package hitbot

import (
	"log"
	"strings"
)

//CmdHandler is interface for command handler creation.
type CmdHandler interface {
	handle(body []string)
}

//BasicCmd is command handler for simple text responses.
type BasicCmd struct {
	response string
}

func (cmd BasicCmd) handle(body []string) {
	log.Print(cmd.response)
}

//RegisterHandler registers specified handler for command of provided name. Will overwrite is command already exists.
func (bot *Hitbot) RegisterHandler(name string, handler CmdHandler) {
	bot.cmdHandlers[name] = handler
	log.Printf("Registered '%v' command", name)
}

//GetBasicCmd creates instance of basic cmd handler.
func GetBasicCmd(response string) BasicCmd {
	return BasicCmd{response: response}
}

func (bot *Hitbot) dispatchCommand(params map[string]interface{}) {
	cmd := strings.Split(params["text"].(string), " ")
	log.Printf("%v invoked command '%v'", params["name"].(string), cmd[0][1:])
	if handler, prs := bot.cmdHandlers[cmd[0][1:]]; prs {
		handler.handle(cmd[1:])
	}
}
