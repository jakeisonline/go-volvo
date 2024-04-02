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
	foo := api.GetEndpoint("vehicle", "awdawd")
	fmt.Println(foo.Url)
}
