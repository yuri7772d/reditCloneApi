package controller

import _commentService "github.com/yuri7772d/reditCloneApi/pkg/comment/service"

type commentControllerImpl struct {
	commentService _commentService.CommentService
}

func newCommentController(commentService _commentService.CommentService) CommentController {
	return &commentControllerImpl{commentService}
}
