package pkg

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	bash_profile = "../bashprofilefiles/.bash_profile"
	bash_profile_deleted = "../bashprofilefiles/.bash_profile_deleted"
	bash_profile_repo = "../bashprofilefiles/.bash_profile_repo"
)

type BashProfiler struct {}


func (bp *BashProfiler) Pull() error {
	// get raw data from file
	repoRaw, err := bp.getStringFromFile(bash_profile_repo)
	if err != nil {
		return err
	}
	bashProfileRaw, err := bp.getStringFromFile(bash_profile)
	if err != nil {
		return err
	}
	deletedRaw, err := bp.getStringFromFile(bash_profile_deleted)
	if err != nil {
		return err
	}

	// use raw data to create an array full of bash commands to compare
	repo, bashProfile, deleted, err := bp.getBashProfileFiles(repoRaw, bashProfileRaw, deletedRaw)
	if err != nil {
		return err
	}

	// repo - bash_profile_deleted
	// we want to get rid of what we have already deleted
	repoFilteredOnDeleted := bp.aMinusB(repo, deleted)

	// repoFilteredOnDeleted - bashprofile
	// we want to see what is new from the repo file
	newBash := bp.aMinusB(repoFilteredOnDeleted, bashProfile)

	// bashProfile + newBash
	// merge bashProfile and newBash
	bashProfile = bp.makeUnique(bashProfile)
	sort.Strings(bashProfile)
	if newBash!=nil {
		comment := bp.getNewCommandsHeader()
		bashProfile = append(bashProfile, comment)
	}
	bashProfile = append(bashProfile, newBash...)
	bashProfile = append(bashProfile, "\n#Deleted")

	// bashprofile - repo
	newRepo := bp.aMinusB(bashProfile, repo)
	// clean garbage lines we dont want in the repo
	newRepo = bp.aMinusB(newRepo, bp.getGarbage())
	repo = append(repo, newRepo...)

	// make all commands unique in a set
	repo = bp.makeUnique(repo)
	deleted = bp.makeUnique(deleted)

	err = bp.writeBashFiles(repo, bashProfile, deleted)
	if err != nil {
		return err
	}

	return nil
}

func (bp *BashProfiler) makeUnique(d []string) []string {
	check := make(map[string]int)
	res := make([]string,0)
	for _, val := range d {
		check[val] = 1
	}

	for letter, _ := range check {
		res = append(res,letter)
	}

	return res
}

func (bp *BashProfiler) splitDeleted(bashProfileRaw string) ([]string, []string, error) {
	var commandsKeep []string
	var commandsDelete []string
	deletedFound := false
	requests := strings.Split(bashProfileRaw, "\n")
	for i, c := range requests {
		if strings.Contains(c, "#Deleted") {
			commandsKeep = requests[:i]
			commandsDelete = requests[i+1:]
			deletedFound = true
		}
	}
	if !deletedFound {
		return requests, nil, nil
	}

	return commandsKeep, commandsDelete, nil
}

func (bp *BashProfiler) aMinusB(a []string, b []string) []string {
	var rs []string
	for _, ca := range a {
		aInB := false
		for _, cb := range b {
			// comparing commands, removing white spaces to avoid confusions
			if strings.TrimSpace(cb) == strings.TrimSpace(ca) {
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

func (bp *BashProfiler) SaveHeaderName() {
	// get body and append new message
	b, err := ioutil.ReadFile("headers.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	b = []byte(string(b) + "\n#" + os.Args[1])

	// write the whole body at once
	if len(os.Args) == 2 {
		err = ioutil.WriteFile("headers.txt", b, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func (bp *BashProfiler) getNewCommandsHeader() string {
		bp.SaveHeaderName()

		// get body and append new message
		b, err := ioutil.ReadFile("headers.txt")
		if err != nil {
			log.Fatalf(err.Error())
		}

		rand.Seed(int64(time.Now().Minute()))
		Headers := strings.Split(string(b), "\n")
		return  Headers[rand.Intn(len(Headers))]
}

func (bp *BashProfiler) getBashProfileFiles(
	repoRaw string, bashProfileRaw string, deletedRaw string) ([]string, []string, []string, error) {
	// read repo, bashprofile and bash_profile_deleted into array
	repo, err := bp.getCommands(strings.Split(repoRaw, "\n"))
	if err != nil {
		return nil, nil, nil, err
	}

	// split bashprofile to bashprofile and deleted (as there is a deleted subsection under bash_profile
	bashProfileArray, deletedNewArray, err := bp.splitDeleted(bashProfileRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	bashProfile, err := bp.getCommands(bashProfileArray)
	if err != nil {
		return nil, nil, nil, err
	}
	deletedNew, err := bp.getCommands(deletedNewArray)
	if err != nil {
		return nil, nil, nil, err
	}

	// add deleted from bash_profile to bash_profile_deleted
	// using another error variable for this as we can have nil bash_profile_deleted
	deleted, errD := bp.getCommands(strings.Split(deletedRaw, "\n"))
	if errD != nil && deleted != nil {
		return nil, nil, nil, errD
	}

	if len(deletedNew)!=0 {
		deleted = append(deleted, deletedNew...)
	}

	return repo, bashProfile, deleted, err
}

func (bp *BashProfiler) getStringFromFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (bp *BashProfiler) getCommands (requests []string) ([]string, error) {
	var commands []string
	bashCommand := ""
	isReadingCommand := false
	for _, c := range requests {
		if c == "" {
			continue
		}
		c = strings.TrimSpace(c)
		if strings.Contains(c, "{") && !isReadingCommand{
			// starting a multi-line command
			isReadingCommand = true
			bashCommand = c
		} else if isReadingCommand{
			// ending a multi-line command
			if strings.Contains(c, "}") {
				if !strings.Contains(c, "$") {
					bashCommand += "\n" + c
					isReadingCommand = false
					commands = append(commands, bashCommand)
					continue					
				}
			}
			bashCommand += "\n\t"+c
		} else {
			// single line command
			commands = append(commands, c)
		}
	}
	return commands, nil
}

func (bp *BashProfiler) getGarbage() []string {
	return []string {
		"\n",
		"#Deleted",
		"#New commands for sweet, sweet Bash Profiler",
		"#Oh yes baby! New commands swinging in from Bash Profiler",
		"#Bash profiler! Sweet, new commands",
		"#YES! Getting those new commands baby Bash Profiler!",
		"#Oh my gosh, yes. This is sweet. New commands! Bash Profiler!",
		"#New commands! Bash Profiler! Yes!",
		"#Nice work there Bash Profiler! YES!",
	}
}

func (bp *BashProfiler) writeBashFiles(repo []string, bashProfile []string, deleted []string) error {
	// write to .bash_profile
	sort.Strings(repo)
	sort.Strings(deleted)

	fileOut := ""
	for _, p := range bashProfile {
		fileOut +=  p + "\n"
	}
	err := ioutil.WriteFile(bash_profile, []byte(fileOut), os.ModePerm)
	if err != nil {
		return err
	}

	// write to .bash_profile_delete
	fileOut = ""
	for _, p := range deleted {
		fileOut += p + "\n"
	}
	err = ioutil.WriteFile(bash_profile_deleted, []byte(fileOut), os.ModePerm)
	if err != nil {
		return err
	}

	// write to .bash_profile_repo //merge
	fileOut = ""
	for _, p := range repo {
		fileOut += p + "\n"
	}
	err = ioutil.WriteFile(bash_profile_repo, []byte(fileOut), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
