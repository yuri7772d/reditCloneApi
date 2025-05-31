package entities

import "time"

type Topic struct {
	Id       int
	Title    string
	Picture  string
	UserId   int
	CreateAt time.Time
}
