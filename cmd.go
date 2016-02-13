package hitbot

import (
	"log"
	"strings"
)

//CmdHandler is a type of function capable of handling commands.
type CmdHandler func(ChatParams) (string, string)

//RegisterHandler registers specified handler for command of provided name. Will overwrite if command already exists.
func (bot *Hitbot) RegisterHandler(name string, handler CmdHandler) {
	bot.registerHandler(name, handler)
	if bot.verbose {
		log.Printf("Registered %v handler", name)
	}
}

func (bot *Hitbot) registerHandler(name string, handler CmdHandler) {
	bot.cmdHandlers[name] = handler
}

type basicCmd struct {
	response string
}

func (cmd basicCmd) Handle(params ChatParams) (string, string) {
	return params.Channel, cmd.response
}

//BasicCmd creates and registers basic cmd handler.
func (bot *Hitbot) BasicCmd(name string, response string) {
	bot.RegisterHandler(name, basicCmd{response: response}.Handle)
}

func (bot *Hitbot) dispatchCommand(params ChatParams) {
	cmd := strings.Split(params.Text, " ")
	if handler, prs := bot.cmdHandlers[cmd[0][1:]]; prs {
		bot.sendMessage(handler(params))
	}
}
