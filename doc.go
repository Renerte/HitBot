/*
Package hitbot allows for easy bot creation for Hitbox.tv streaming platform.


Setup:


To create bot instance use:
    bot := hitbox.NewBot("name")

Then you need to get server list:
    bot.GetServers()

After that you need to get connection id for one of the servers:
    bot.GetID()

Before you can connect, you have to login to Hitbox.tv for access token. Following command does it for you:
    bot.Auth("pass")

At any point you can register commands by using provided BasicCmd handler factory, or create your own. To create basic cmd response:
    bot.BasicCmd("cmdname","response")

For custom handlers use:
    bot.RegisterHandler("cmdname", handler)
Where `handler` is:
    func handler(params map[string]interface{}) (string, string){
        //returns channel and response text
    }

If you don't want your bot's name to be white, you can set color with:
    bot.NameColor("hex value without `#`")

Then you can finally connect, and start MessageHandler:
    bot.Connect("channel")
    bot.MessageHandler()

Channels specified in Connect method will be joined as soon as MessageHandler recieves confirmation for connection, you can still join channels manually, just make sure it happens after confirmation.
Keep in mind, you can run MessageHandler as goroutine, so you can perform actions within your program.
*/
package hitbot
