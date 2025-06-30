package repository

import (
	"github.com/yuri7772d/reditCloneApi/entities"
	"gorm.io/gorm"
)

type Topic interface {
	Begin() (*gorm.DB, error)
	Commit(*gorm.DB) error
	Rollback(*gorm.DB) error
	Create(*entities.Topic) error
	Listing() ([]*entities.Topic, error)
}
