package service

import (
	profileService "github.com/yuri7772d/reditCloneApi/pkg/profile/service"
	"github.com/yuri7772d/reditCloneApi/pkg/topic/model"
)

type Topic interface {
	Post(int, string, ...string)
	Listing(profileService.Profile) []*model.Topic
}
