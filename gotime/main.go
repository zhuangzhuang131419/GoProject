package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	fmt.Printf("now: %v\n", timestamp)
	parse, err := time.Parse("2006-01-02T15:04:05.000Z", timestamp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("since: %v\n", parse)
	fmt.Printf("since: %v\n", time.Since(parse))

}
