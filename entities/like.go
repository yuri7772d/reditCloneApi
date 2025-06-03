package entities

import "time"

type LikeTranking struct {
	Id          int
	LikeCountId uint64
	UserId      uint64
	CreateAt    time.Time
}
