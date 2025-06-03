package entities

type LikeCount struct {
	Id           int
	TargetId     uint64
	TargetTypeId uint8
	Count        int
}
