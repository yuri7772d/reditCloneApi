package service

import (
	"github.com/yuri7772d/reditCloneApi/entities"
	"github.com/yuri7772d/reditCloneApi/pkg/profile/model"
	"github.com/yuri7772d/reditCloneApi/pkg/profile/repository"
)

type profileImpl struct {
	repository repository.Profile
}

func New(repository repository.Profile) Profile {
	return &profileImpl{repository}
}

func (s *profileImpl) Create(name string, about string, img string) {
	profile := &entities.Profile{Name: name, About: about}
	if img != "" {
		profile.Image = img
	}

	tx, _ := s.repository.Begin()
	if err := s.repository.Create(profile); err != nil {
		s.repository.Rollback(tx)
	}
	s.repository.Commit(tx)
}
func (s *profileImpl) GetBrief(ids ...int) []model.Brief {

	brief, _ := s.repository.GetBrief(ids...)

	return brief
}
