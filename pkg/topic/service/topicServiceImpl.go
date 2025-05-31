package service

import _repository "github.com/yuri7772d/reditCloneApi/pkg/topic/repository"

type topicServiceImpl struct {
	topicRepository _repository.TopicRepository
}

func newCommentService(topicRepository _repository.TopicRepository) TopicService {
	return &topicServiceImpl{topicRepository}
}
