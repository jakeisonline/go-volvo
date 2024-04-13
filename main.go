/**
 * Author: Jake Holman
 * File: main.go
 */

package main

import (
	"fmt"

	"github.com/jakeisonline/go-volvo/api"
)

func main() {
	client := api.NewClient("15b6c230efdb408ba8bd260a4757e71a")
	fmt.Print(client)
}
