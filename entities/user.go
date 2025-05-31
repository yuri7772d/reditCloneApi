package entities

import "time"

type User struct {
	Id       int
	Name     string
	CreateAt time.Time
}
