package databases

import "gorm.io/gorm"

type Database interface {
	ConnetionGeting() *gorm.DB
}
