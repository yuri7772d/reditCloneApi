package entities

import (
	"time"
)

type Replied struct {
	ID         int            `gorm:"primaryKey;autoIncrement"`
	Content    string         `gorm:"type:varchar(500);not null;"`
	ProfileID  int            `gorm:"index;not null;"`
	Profile    Profile        `gorm:"foreignKey:ProfileID;references:ID"`
	CreatedAt  time.Time      `gorm:"type:timestamptz;default:now()"`
	ParentID   *int           `gorm:"index:idx_replied"`
	ParentType string         `gorm:"index:idx_replied"`
	Replieds   []Replied      `gorm:"polymorphic:Parent"`
	Images     []RepliedImage `gorm:"foreignkey:RepliedID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Like       RepliedLike    `gorm:"foreignkey:RepliedID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type RepliedImage struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	RepliedID int    `gorm:"index;not null;"`
	Image     string `gorm:"type:varchar(64);not null;"`
}

type RepliedLike struct {
	ID        int                   `gorm:"primaryKey;autoIncrement;"`
	RepliedID int                   `gorm:"index;not null;"`
	Count     int                   `gorm:"default:0"`
	Trackings []RepliedLikeTracking `gorm:"foreignkey:RepliedLikeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type RepliedLikeTracking struct {
	ID            int       `gorm:"primaryKey;autoIncrement;"`
	ProfileID     int       `gorm:"index;not null;"`
	Profile       Profile   `gorm:"foreignKey:ProfileID;references:ID"`
	RepliedLikeID int       `gorm:"index;not null;"`
	CreatedAt     time.Time `gorm:"type:timestamptz;default:now()"`
}
