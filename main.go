package main

import "fmt"

type (
	item interface {
		use() error
	}
	sword struct {
		Damage int
	}
)

func newItem(s int) item {
	return &sword{Damage: s}
}

func (s *sword) use() error {
	fmt.Printf("Use Sword!! tack Damage %v", s.Damage)
	return nil
}

func useItem(i item) error {
	return i.use()

}

func main() {
	_sword := &sword{Damage: 1}
	useItem(_sword)
	//fmt.Printf("hello world!")
}
