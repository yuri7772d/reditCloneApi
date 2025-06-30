package repository

import (
	"github.com/yuri7772d/reditCloneApi/databases"
	"github.com/yuri7772d/reditCloneApi/entities"
	"gorm.io/gorm"
)

type topicImpl struct {
	db databases.Database
}

func New(db databases.Database) Topic {
	return &topicImpl{db}
}

func (r *topicImpl) Begin() (*gorm.DB, error) {
	tx := r.db.Connet().Begin()
	return tx, tx.Error
}
func (r *topicImpl) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}
func (r *topicImpl) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *topicImpl) Create(topic *entities.Topic) error {
	db := r.db.Connet()

	if err := db.Create(topic).Error; err != nil {
		return err
	}
	if err := db.Create(&entities.TopicLike{TopicID: topic.ID}).Error; err != nil {
		return err
	}
	return nil
}
func (r *topicImpl) Listing() ([]*entities.Topic, error) {
	db := r.db.Connet()
	topics := make([]*entities.Topic, 0)
	if err := db.Preload("Like").Preload("Images").Find(&topics).Error; err != nil {
		return nil, err
	}

	return topics, nil
}
