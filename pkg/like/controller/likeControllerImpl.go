package controller

import _likeService "github.com/yuri7772d/reditCloneApi/pkg/like/service"

type likeControllerImpl struct {
	commentService _likeService.LikeService
}

func newCommentController(commentService _likeService.LikeService) LikeController {
	return &likeControllerImpl{commentService}
}
