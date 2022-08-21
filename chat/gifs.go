package chat

import (
	"io/ioutil"
	"net/http"
)

func GetGif(name string)string{
	resp, err := http.Get("http://api.nekos.fun:8080/api/"+name)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	link := string(body)
	return link[10:len(link)-2]
}