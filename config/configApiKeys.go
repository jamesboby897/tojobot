package config

import (
	"encoding/json"
	"os"
)
type Keys struct{
	BotToken string
	DeveloperKey string
}
func GetKeys()Keys{
	keys := Keys{}
	file,err := os.Open("keys.json")
	if err != nil{
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	fi,err:=file.Stat()
	if err != nil{
		panic(err)
	}
	data := make([]byte,fi.Size())
	_, err = file.Read(data)
	if err != nil{
		panic(err)
	}
	err = json.Unmarshal(data,&keys)
	if err != nil{
		panic(err)
	}
	return keys
}
