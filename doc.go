/*
Package hitbot allows for easy bot creation for Hitbox.tv streaming platform.


Setup:

To load bot from config file use:
    bot := hitbot.LoadBot("/path/to/file.json", true) //where true is verbose flag

After that register any custom handlers you want with:
    bot.RegisterHandler("name", HandlerInit)

When you are ready, load the commands:
    bot.LoadCommands()

Bot is now ready to start:
    bot.MessageHandler()
*/
package hitbot
