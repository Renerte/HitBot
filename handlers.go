package hitbot

type basicCmd struct {
	response string
}

func (cmd basicCmd) Handle(params ChatParams) (string, string) {
	return params.Channel, cmd.response
}

func basicInit(data HandlerData) HandlerFunc {
	return basicCmd{response: data.(string)}.Handle
}
