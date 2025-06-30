package repository

import (
	"github.com/yuri7772d/reditCloneApi/entities"
	"github.com/yuri7772d/reditCloneApi/pkg/profile/model"
	"gorm.io/gorm"
)

type Profile interface {
	Begin() (*gorm.DB, error)
	Commit(*gorm.DB) error
	Rollback(*gorm.DB) error
	Create(*entities.Profile) error
	GetBrief(...int) ([]model.Brief, error)
}
