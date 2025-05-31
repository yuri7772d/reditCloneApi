package service

import _repository "github.com/yuri7772d/reditCloneApi/pkg/user/repository"

type userServiceImpl struct {
	userRepository _repository.UserRepository
}

func newCommentService(userRepository _repository.UserRepository) UserService {
	return &userServiceImpl{userRepository}
}
