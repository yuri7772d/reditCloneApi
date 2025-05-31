package service

import _repository "github.com/yuri7772d/reditCloneApi/pkg/comment/repository"

type commentServiceImpl struct {
	commentRepository _repository.CommentRepository
}

func newCommentService(commentRepository _repository.CommentRepository) CommentService {
	return &commentServiceImpl{commentRepository}
}
