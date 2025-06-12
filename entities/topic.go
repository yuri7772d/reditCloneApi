package entities

import (
	"time"
)

type (
	Topic struct {
		ID       int
		Title    string
		Images   []string
		createAt time.Time
	}
	TopicRecord struct {
		ID       int       `gorm:"primaryKey;autoIncrement;"`
		Title    string    `gorm:"type:varchar(64);not null;"`
		CreateAt time.Time `gorm:"type:timestamptz"`
	}
	TopicImageRecord struct {
		TopicID  int         `gorm:"type:int(64);not null;"`
		Image    string      `gorm:"type:varchar(64);not null;"`
		TopicRef TopicRecord `gorm:"foreignKey:TopicID;references:ID;constraint:OnDelete:CASCADE"`
	}
)

func (TopicImageRecord) TableName() string {
	return "topic_image"
}

func (TopicRecord) TableName() string {
	return "topic"
}

func NewTopic(id int, title string, images []string, time time.Time) *Topic {
	return &Topic{ID: id, Title: title, Images: images, createAt: time}
}

func (t *Topic) GetCreateAt() time.Time {
	return t.createAt
}
func (t *Topic) ToRecord() (*TopicRecord, []*TopicImageRecord) {
	imageRecInstance := make([]*TopicImageRecord, len(t.Images))
	for i, r := range t.Images {
		imageRecInstance[i] = &TopicImageRecord{TopicID: t.ID, Image: r}
	}

	return &TopicRecord{ID: t.ID, Title: t.Title, CreateAt: t.createAt}, imageRecInstance
}
