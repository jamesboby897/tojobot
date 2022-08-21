package main

import (
        "tojobot/bot"
        "tojobot/config"
)

func main() {
        keys := config.GetKeys()
        bot.InitBot(keys.BotToken)
}