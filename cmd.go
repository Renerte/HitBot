package hitbot

import (
	"log"
	"strings"
)

//HandlerFunc is a type of function capable of handling commands.
type HandlerFunc func(ChatParams) (string, string)

//HandlerData holds data provided for handler to use.
type HandlerData interface{}

//HandlerInit prepares handler and applies HandlerData provided to it.
type HandlerInit func(HandlerData) HandlerFunc

//RegisterHandler registers specified handler for command of provided name. Will overwrite if command already exists.
func (bot *Hitbot) RegisterHandler(name string, handler HandlerInit) {
	bot.registerHandler(name, handler)
	if bot.verbose {
		log.Printf("Registered %v handler", name)
	}
}

func (bot *Hitbot) registerHandler(name string, handler HandlerInit) {
	bot.handlers[name] = handler
}

//RegisterCommand registers command with specified name, handler, and data.
func (bot *Hitbot) RegisterCommand(name string, handler string, data HandlerData) {
	bot.registerCommand(name, handler, data)
	if bot.verbose {
		log.Printf("Registered '%v' command with handler '%v'", name, handler)
	}
}

func (bot *Hitbot) registerCommand(name string, handler string, data HandlerData) {
	bot.cmdHandlers[name] = cmd{Handler: handler, Data: data}
	bot.cmds[name] = bot.handlers[handler](data)
}

type basicCmd struct {
	response string
}

func (cmd basicCmd) Handle(params ChatParams) (string, string) {
	return params.Channel, cmd.response
}

func basicInit(data HandlerData) HandlerFunc {
	return basicCmd{response: data.(string)}.Handle
}

//BasicCmd creates and registers basic cmd handler.
func (bot *Hitbot) BasicCmd(name string, response string) {
	bot.RegisterCommand(name, "basic", response)
}

func (bot *Hitbot) dispatchCommand(params ChatParams) {
	cmd := strings.Split(params.Text, " ")
	if handler, prs := bot.cmds[cmd[0][1:]]; prs {
		bot.sendMessage(handler(params))
	}
}
