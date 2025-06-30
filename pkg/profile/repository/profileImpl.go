package repository

import (
	"github.com/yuri7772d/reditCloneApi/databases"
	"github.com/yuri7772d/reditCloneApi/entities"
	"github.com/yuri7772d/reditCloneApi/pkg/profile/model"
	"gorm.io/gorm"
)

type profileImpl struct {
	db databases.Database
}

func New(db databases.Database) Profile {
	return &profileImpl{db}
}

func (r *profileImpl) Begin() (*gorm.DB, error) {
	tx := r.db.Connet().Begin()
	return tx, tx.Error
}
func (r *profileImpl) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}
func (r *profileImpl) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *profileImpl) Create(profile *entities.Profile) error {
	db := r.db.Connet()
	if err := db.Create(profile); err != nil {
		return err.Error
	}
	return nil
}
func (r *profileImpl) GetBrief(ids ...int) ([]model.Brief, error) {
	db := r.db.Connet()

	var briefs []model.Brief
	if err := db.Model(&entities.Profile{}).Select("id", "name", "image").Where("id IN ? ", ids).Scan(&briefs).Error; err != nil {
		return nil, err
	}

	return briefs, nil
}
