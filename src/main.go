package main

import (
	"brodsky/cmd"
	"brodsky/pkg/log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
