package main

import (
	"github.com/kainhuck/dod/cmd"
	"log"
)

func main(){
	if err := cmd.Execute(); err != nil {
		log.Fatal("tidy run failed: %v", err)
	}
}
