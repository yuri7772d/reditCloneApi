package databases

import "gorm.io/gorm"

type Database interface {
	Connet() *gorm.DB
}
