package main

import (
	"fmt"
	"tojobot/bot"
	"tojobot/config"
)

func main(){
	err := config.ReadConfig()

	if err != nil{
		fmt.Println(err.Error())
		return
	}
	bot.Start()
	<-make(chan struct{})
	return
}
