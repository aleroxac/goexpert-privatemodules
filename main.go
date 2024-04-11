package main

import (
	"fmt"

	"github.com/aleroxac/goexpert-utils/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
