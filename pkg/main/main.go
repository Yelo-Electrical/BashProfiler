package main

import (
	pkg "github.com/Yelo-Electrical/BashProfiler/pkg"
	"log"
)

func main() {
	log.Print("Merging Bash Profiler...")
	bp := &pkg.BashProfiler{}
	err := bp.Pull()
	if err != nil {
		log.Fatalf("Err: %v", err.Error())
	}
	log.Printf("Go code finished executing.")
}

