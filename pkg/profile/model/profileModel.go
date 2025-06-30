package model

import (
	"time"
)

type Profile struct {
	ID        int
	Name      string
	CreatedAt time.Time
}
type Brief struct {
	ID    int
	Name  string
	Image string
}
