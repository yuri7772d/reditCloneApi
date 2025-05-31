package controller

import _topicService "github.com/yuri7772d/reditCloneApi/pkg/topic/service"

type topicControllerImpl struct {
	topicService _topicService.TopicService
}

func newCommentController(topicService _topicService.TopicService) TopicController {
	return &topicControllerImpl{topicService}
}
