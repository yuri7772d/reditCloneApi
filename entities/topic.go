package entities

import (
	"time"
)

type (
	Topic struct {
		ID        int          `gorm:"primaryKey;autoIncrement"`
		Title     string       `gorm:"type:varchar(500);not null;"`
		ProfileID int          `gorm:"index;not null;"`
		Profile   Profile      `gorm:"foreignKey:ProfileID;references:ID"`
		CreatedAt time.Time    `gorm:"type:timestamptz;default:now()"`
		Images    []TopicImage `gorm:"foreignkey:TopicID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Like      TopicLike    `gorm:"foreignkey:TopicID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Replieds  []Replied    `gorm:"polymorphic:Parent"`
	}
	TopicImage struct {
		ID      int    `gorm:"primaryKey;autoIncrement"`
		TopicID int    `gorm:"index;not null;"`
		Image   string `gorm:"type:varchar(64);not null;"`
	}
	TopicLike struct {
		ID        int                 `gorm:"primaryKey;autoIncrement;"`
		TopicID   int                 `gorm:"uniqueIndex;not null;"`
		Count     int                 `gorm:"default:0"`
		Trackings []TopicLikeTracking `gorm:"foreignkey:TopicLikeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}
	TopicLikeTracking struct {
		ID          int       `gorm:"primaryKey;autoIncrement;"`
		ProfileID   int       `gorm:"index;not null;"`
		Profile     Profile   `gorm:"foreignKey:ProfileID;references:ID"`
		TopicLikeID int       `gorm:"index;not null;"`
		CreatedAt   time.Time `gorm:"type:timestamptz;default:now()"`
	}
)
