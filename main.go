package main

import (
	"brodsky/cmd"
	"brodsky/internal/log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
