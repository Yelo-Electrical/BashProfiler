package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
)

const (
	bash_profile = "../bashprofilefiles/.bash_profile"
	bash_profile_deleted = "../bashprofilefiles/.bash_profile_deleted"
	bash_profile_repo = "../bashprofilefiles/.bash_profile_repo"
)

type BashProfiler struct {}

func main() {
	log.Print("Merging Bash Profiler...")
	bp := &BashProfiler{}
	err := bp.Pull()
	if err != nil {
		log.Fatalf("Err: %v", err.Error())
	}
	log.Printf("Go code finished executing.")
}

func (bp *BashProfiler) Pull() error {
	repo, bashProfile, deleted, err := bp.GetBashProfileFiles()
	if err != nil {
		return err
	}

	// repo - bash_profile_deleted
	// we want to get rid of what we have already deleted
	repoFilteredOnDeleted := bp.AMinusB(repo, deleted)

	// repoFilteredOnDeleted - bashprofile
	// we want to see what is new from the repo file
	newBash := bp.AMinusB(repoFilteredOnDeleted, bashProfile)
	newBash = append([]string{bp.GetNewCommandsHeader()}, newBash...)

	// bashProfile + newBash
	// merge bashProfile and newBash
	bashProfile = append(bashProfile, newBash...)

	// result - repo
	newRepo := bp.AMinusB(bashProfile, repo)
	repo = append(repo, newRepo...)

	err = bp.WriteBashFiles(repo, bashProfile, deleted)
	if err != nil {
		return err
	}

	return nil
}

func (bp *BashProfiler) SplitDeleted () ([]string, []string, error) {
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

func (bp *BashProfiler) WriteBashFiles(repo []string, bashProfile []string, deleted []string) error {
	// write to .bash_profile
	fileOut := ""
	for _, p := range bashProfile {
		fileOut += "\r\n" + p
	}
	log.Printf("Writing to: %v", bash_profile)
	err := ioutil.WriteFile(bash_profile, []byte(fileOut), os.ModePerm)
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
	return nil
}

func (bp *BashProfiler) AMinusB(a []string, b []string) []string {
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

func (bp *BashProfiler) GetNewCommandsHeader() string {
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

func (bp *BashProfiler) GetBashProfileFiles() ([]string, []string, []string, error) {
	// read repo, bashprofile and bash_profile_deleted into array
	repo, err := bp.getCommands(bash_profile_repo)
	if err != nil {
		return nil, nil, nil, err
	}

	// split bashprofile to bashprofile and deleted (as there is a deleted subsection under bash_profile
	bashProfile, deletedNew, err := bp.SplitDeleted()
	if err != nil {
		return nil, nil, nil, err
	}

	// add deleted from bash profile to delete
	// using another error variable for this as we can have nil bash_profile_deleted
	deleted, errD := bp.getCommands(bash_profile_deleted)
	if errD != nil && deleted != nil {
		return nil, nil, nil, errD
	}

	if len(deletedNew)!=0 {
		deleted = append(deleted, deletedNew...)
	}

	sort.Strings(repo)
	sort.Strings(bashProfile)
	sort.Strings(deleted)
	return repo, bashProfile, deleted, err
}

func (bp *BashProfiler) getCommands (file string) ([]string, error) {
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















