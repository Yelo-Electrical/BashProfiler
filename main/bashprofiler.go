package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

const (
	bash_profile = "../bashprofilefiles/.bash_profile"
	bash_profile_deleted = "../bashprofilefiles/.bash_profile_deleted"
	bash_profile_repo = "../bashprofilefiles/.bash_profile_repo"
)

func main() {
	log.Print("Merging Bash Profiler...")

	err := Pull()
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}
}

func Pull() error {
	// read repo, bashprofile and bash_profile_deleted into array
	repo, err := GetCommands(bash_profile_repo)
	if err != nil {
		return err
	}

	// split bashprofile to bashprofil and deleted
	bashProfile, deletedNew, err := SplitDeleted()

	// add deleted from bash profile to delete
	deleted, err := GetCommands(bash_profile_deleted)
	if err != nil {
		return err
	}

	if len(deletedNew)!=0 {
		deleted = append(deleted, deletedNew...)
	}


	// repo - bash_profile_deleted
	// we want to get rid of what we have already deleted
	repoFilteredOnDeleted := AMinusB(repo, deleted)

	// repoFilteredOnDeleted - bashprofile
	// we want to see what is new from the repo file
	newBash := AMinusB(repoFilteredOnDeleted, bashProfile)
	newBash = append([]string{GetNewCommandsHeader()}, newBash...)

	// bashProfile + newBash
	// merge bashProfile and newBash
	result := append(bashProfile, newBash...)

	// write to .bash_profile
	fileOut := ""
	for _, p := range result {
		fileOut += "\r\n" + p
	}
	log.Printf("Writing to: %v", bash_profile)
	err = ioutil.WriteFile(bash_profile, []byte(fileOut), os.ModePerm)
	if err != nil {
		return err
	}

	// write to .bash_profile_delete
	fileOut = ""
	for _, p := range deleted {
		fileOut += "\r\n" + p
	}
	log.Printf("Writing to: %v", bash_profile_deleted)
	err = ioutil.WriteFile(bash_profile_deleted, []byte(fileOut), os.ModePerm)
	if err != nil {
		return err
	}

	// result - repo
	newRepo := AMinusB(bashProfile, repo)
	repo = append(repo, newRepo...)

	// write to .bash_profile_repo //merge
	fileOut = ""
	for _, p := range repo {
		fileOut += "\r\n" + p
	}
	log.Printf("Writing to: %v", bash_profile_repo)
	err = ioutil.WriteFile(bash_profile_repo, []byte(fileOut), os.ModePerm)
	if err != nil {
		return err
	}

	log.Printf("Everything pulled and rewritten to working files. Now to copy .bash_profile")
	return nil
}

func GetCommands (file string) ([]string, error) {
	log.Printf("Reading: %v", file)
	var commands []string
	commandsRaw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	requests := strings.Split(string(commandsRaw), "\r\n")
	bashCommand := ""
	isReadingCommand := false
	for _, c := range requests {
		if strings.Contains(c, "{") && !isReadingCommand{
			// starting a multi-line command
			isReadingCommand = true
			bashCommand = "\n"+c
		} else if isReadingCommand{
			// ending a multi-line command
			bashCommand += "\n"+c
			if strings.Contains(c, "}") {
				isReadingCommand = false
				commands = append(commands, "\n"+bashCommand)
			}
		} else {
			// single line command
			commands = append(commands, "\n"+c)
		}
	}
	return commands, nil
}

func SplitDeleted () ([]string, []string, error) {
	log.Printf("Reading: %v", bash_profile)
	var commandsKeep []string
	var commandsDelete []string
	deletedFound := false
	commandsRaw, err := ioutil.ReadFile(bash_profile)
	if err != nil {
		return nil, nil, err
	}
	requests := strings.Split(string(commandsRaw), "\n")
	for i, c := range requests {
		if c=="\n" {
			continue
		}
		if strings.Contains(c, "#Deleted") {
			commandsKeep = requests[:i]
			commandsDelete = requests[i+1:]
			deletedFound = true
		}
	}
	if len(commandsDelete)!=0 {
		commandsDelete[0] = "\r\n" + commandsDelete[0]
	}
	if !deletedFound {
		return requests, nil, nil
	}

	return commandsKeep, commandsDelete, nil
}

func AMinusB(a []string, b []string) []string {
	var rs []string
	for _, ca := range a {
		aInB := false
		for _, cb := range b {
			if cb == ca {
				// command exists, move on
				aInB = true
				continue
			}
		}
		if !aInB {
			rs = append(rs, ca)
		}
	}
	return rs
}

func GetNewCommandsHeader() string {
	s := []string{
		"#New commands for sweet, sweet Bash Profiler",
		"#Oh yes baby! New commands swinging in from Bash Profiler",
		"#Bash profiler! Sweet, new commands",
		"#YES! Getting those new commands baby Bash Profiler!",
		"#Oh my gosh, yes. This is sweet. New commands! Bash Profiler!",
		"#New commands! Bash Profiler! Yes!",
		"#Nice work there Bash Profiler! YES!",
	}
	return s[rand.Intn(len(s))]
}
















