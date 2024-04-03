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
	resp, err := api.Call("vehicles", "")
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(resp)
	}
}
