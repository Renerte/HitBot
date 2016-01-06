/*
Package hitbot allows for easy bot creation for Hitbox.tv streaming platform.

To create bot instance use:
    bot := hitbox.NewBot("name")

Then you need to get server list:
    bot.GetServers()

After that you need to get connection id for one of the servers:
    bot.GetID()
*/
package hitbot
