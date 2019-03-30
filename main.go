package main

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	embeddedFile := packr.New("embedded", "./templates")
	fmt.Println(embeddedFile)
}
