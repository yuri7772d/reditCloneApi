package main

import (
	"fmt"

	"github.com/yuri7772d/reditCloneApi/config"
	"github.com/yuri7772d/reditCloneApi/databases"
	"github.com/yuri7772d/reditCloneApi/entities"
	"github.com/yuri7772d/reditCloneApi/server"
)

func viewTrack(t *entities.TopicLike) {
	likeTrack := t.Tracked()
	for _, r := range likeTrack {
		fmt.Printf("%v ", r)
	}
	fmt.Println("\nTotal : ", t.Total())

}

type ()

func main() {
	conf, _ := config.GetingConfig()
	db := databases.NewPosgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnetionGeting())
	server.Start()
}
