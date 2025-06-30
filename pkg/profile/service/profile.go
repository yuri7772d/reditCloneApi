package service

import "github.com/yuri7772d/reditCloneApi/pkg/profile/model"

type Profile interface {
	Create(string, string, string) // Create(name, about, img) creates a profile with basic info.
	GetBrief(...int) []model.Brief
}
