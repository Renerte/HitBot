/*
Package hitbot allows for easy bot creation for Hitbox.tv streaming platform.

Setup:

To create bot instance use:
    bot := hitbox.NewBot("name")

Then you need to get server list:
    bot.GetServers()

After that you need to get connection id for one of the servers:
    bot.GetID()


Authentication:

Before you can connect, you have to login to Hitbox.tv for access token. Following command does it for you:
    bot.Auth("name","pass")
*/
package hitbot
