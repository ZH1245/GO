package main

import (
	"fmt"
	p "pretty-printer/printer"
)

func main() {
	p.AddBorder("WELCOME", 10)
	message:= `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

	cleaned :=p.CleanUpMessage(message)
	fmt.Println(cleaned)
}
