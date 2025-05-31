package service

import _repository "github.com/yuri7772d/reditCloneApi/pkg/like/repository"

type likeServiceImpl struct {
	likeRepository _repository.LikeRepository
}

func newLikeService(likeRepository _repository.LikeRepository) LikeService {
	return &likeServiceImpl{likeRepository}
}
