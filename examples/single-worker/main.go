package main

import (
	"context"
	"fmt"
	"goku"
)

type Email struct {
	email   string
	message string
}

func (e Email) Process(context.Context) error {
	fmt.Printf("Sending %s to %s", e.message, e.email)
	return nil
}

func main() {
	q, err := goku.NewQueue()
	if err != nil {
		return
	}

	q.RegisterJob("example", Email{})
}
