package main

import (
	"fmt"

	"github.com/yuri7772d/reditCloneApi/entities"
)

func viewTrack(t *entities.TopicLike) {
	likeTrack := t.Tracked()
	for _, r := range likeTrack {
		fmt.Printf("%v ", r)
	}
	fmt.Println("\nTotal : ", t.Total())

}

func main() {

	topicLike := entities.NewTopicLike("asdkjhadj")
	topicLike.Update([]string{"adasdadssad", "asdad", "adasdadfiudf9we8"})
	viewTrack(&topicLike)
	topicLike.Remove("adasdadfiudf9we8")
	viewTrack(&topicLike)

	topicLike.Add("whrewurowejflsdkfjiefuwefko")
	topicLike.Add("qweqeowejflsdkfjiarewewe")
	viewTrack(&topicLike) //*/

}
