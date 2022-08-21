package YT

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"tojobot/config"
)

func SearchVideo(title string,searchResults int64)(id []string,titleReceived []string){
	keys := config.GetKeys()
	service, err := youtube.NewService(context.TODO(),option.WithAPIKey(keys.DeveloperKey))
	if err != nil {
		panic(err)
	}
	parameters := []string{
		"id",
		"snippet",
	}
	call := service.Search.List(parameters).
		Q(title).
		MaxResults(searchResults)

	response, err := call.Do()
	if err != nil {
		panic(err)
	}
	for i :=0; int64(i)<searchResults; i++ {
		id = append(id, response.Items[i].Id.VideoId)
		titleReceived = append(titleReceived,response.Items[i].Snippet.Title)
	}
	return
}