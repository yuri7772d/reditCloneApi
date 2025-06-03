package entities

import "time"

type Topic struct {
	Id       uint64
	Title    string
	Picture  string
	UserId   uint64
	CreateAt time.Time
}
