package entities

import "time"

type Comment struct {
	Id       uint64
	Content  string
	Picture  string
	TopicId  uint64
	ParentId uint64
	UserId   uint64
	CreateAt time.Time
}
