package main

import (
	"fmt"

	"github.com/yuri7772d/reditCloneApi/config"
	"github.com/yuri7772d/reditCloneApi/databases"
	_profileRepository "github.com/yuri7772d/reditCloneApi/pkg/profile/repository"
	_profileService "github.com/yuri7772d/reditCloneApi/pkg/profile/service"
	_topicRepository "github.com/yuri7772d/reditCloneApi/pkg/topic/repository"
	_topicService "github.com/yuri7772d/reditCloneApi/pkg/topic/service"
)

func main() {
	conf, _ := config.Get()
	db := databases.NewPosgresDatabase(conf.Database)

	profileRepository := _profileRepository.New(db)
	profileservice := _profileService.New(profileRepository)
	TopicRepository := _topicRepository.New(db)
	TopicService := _topicService.New(TopicRepository)
	topicList := TopicService.Listing(profileservice)
	if topicList == nil {
		fmt.Println("topicList is nil!!")
	}

	for _, topic := range topicList {
		fmt.Printf("[%v]Name : %v \n", topic.ProfileID, topic.ProfileName)

		fmt.Println("Profile Image : ", topic.ProfileImage)
		fmt.Println("Title : ", topic.Tiltle)
		fmt.Printf("Image : ")
		for _, img := range topic.Images {
			fmt.Printf("%v,", img)
		}
		fmt.Printf("\n")
		fmt.Println("Like count : ", topic.LikeCount)
		fmt.Println("Is like : ", topic.Islike)
	}

	//profileBriefs := profileservice.GetBrief(1, 2, 3, 4, 6)

	/*for _, profileBrief := range profileBriefs {
		fmt.Println(profileBrief.Name, " : ", profileBrief.Image)
	}*/
}
