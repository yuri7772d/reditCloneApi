package entities

import (
	"time"

	"github.com/yuri7772d/reditCloneApi/pkg/profile/model"
)

type Profile struct {
	ID                   int                   `gorm:"primaryKey;autoIncrement;"`
	Name                 string                `gorm:"type:varchar(64);not null;"`
	About                string                `gorm:"type:varchar(500);"`
	Image                string                `gorm:"type:varchar(64);default:'a.png';"`
	CreatedAt            time.Time             `gorm:"type:timestamptz;default:now()"`
	Topics               []Topic               `gorm:"foreignkey:ProfileID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TopicLikeTrackings   []TopicLikeTracking   `gorm:"foreignkey:ProfileID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Replieds             []Replied             `gorm:"foreignkey:ProfileID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RepliedLikeTrackings []RepliedLikeTracking `gorm:"foreignkey:ProfileID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *Profile) ToModel() model.Profile {
	return model.Profile{ID: p.ID, Name: p.Name, CreatedAt: p.CreatedAt}
}
