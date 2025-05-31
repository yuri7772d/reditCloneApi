package controller

import _userService "github.com/yuri7772d/reditCloneApi/pkg/user/service"

type userControllerImpl struct {
	usertService _userService.UserService
}

func newCommentController(usertService _userService.UserService) UserController {
	return &userControllerImpl{usertService}
}
