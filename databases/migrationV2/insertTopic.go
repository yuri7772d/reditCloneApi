package main

import (
	"time"

	"github.com/yuri7772d/reditCloneApi/config"
	"github.com/yuri7772d/reditCloneApi/databases"
	"github.com/yuri7772d/reditCloneApi/entities"
)

var (
	conf *config.Config
	db   databases.Database
)

func postTopic(profileID int, title string, images []string) {
	topic := entities.Topic{Title: title, ProfileID: profileID}

	for _, img := range images {
		topic.Images = append(topic.Images, entities.TopicImage{Image: img})
	}
	db.Connet().Create(&topic)
	like := entities.TopicLike{Count: 0, TopicID: topic.ID}
	db.Connet().Create(&like)
}
func CreateProfile(name string) {
	profile := entities.Profile{Name: name, CreatedAt: time.Now()}
	db.Connet().Create(&profile)
}

func main() {
	conf, _ = config.Get()
	db = databases.NewPosgresDatabase(conf.Database)
	CreateProfile("pungza")
	postTopic(1, "top 10 anime in your hard..", []string{".png", "b.png"})
}
