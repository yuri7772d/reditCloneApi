package model

type Topic struct {
	ID           int
	Tiltle       string
	ProfileID    int
	ProfileName  string
	ProfileImage string
	Images       []string
	LikeCount    int
	Islike       bool
}
