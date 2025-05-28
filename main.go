package main

import "fmt"

type (
	item interface {
		use()
	}
	sword struct{}
)

func (s *sword) use() {
	fmt.Printf("Use Sword!!")

}

func useItem(i item) {
	i.use()

}

func main() {
	Sword := sword{}
	useItem(Sword)
	fmt.Printf("hello world!")
}
s