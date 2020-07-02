package main

import (
	pkg "github.com/Yelo-Electrical/BashProfiler/pkg"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	log.Print("Merging Bash Profiler...")
	log.Print(os.Args[1])
	log.Print(GetHeading())
	bp := &pkg.BashProfiler{}
	err := bp.Pull()
	if err != nil {
		log.Fatalf("Err: %v", err.Error())
	}
	log.Printf("Go code finished executing.")
}

func GetHeading() string {
	// get body and append new message
	b, err := ioutil.ReadFile("headers.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	rand.Seed(int64(time.Now().Minute()))
	Headers := strings.Split(string(b), "")
	return  Headers[rand.Intn(len(Headers))]
}


