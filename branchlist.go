package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/jessjenkins/branchlist/config"
	"github.com/jessjenkins/branchlist/repos"
)

var apikey string

func main() {

	config.Setup()

	cr := make(chan bool)

	// For each repo on github, fire off a routine to collect
	// all the branches for that repo
	c := make(chan repos.OrgRepo)
	go repos.GetOrgRepos(config.Org, c)
	numrepos := 0
	for repo := range c {
		go cloneRepo(repo, cr)
		numrepos++
	}

	// When each fired off routine above finishes, collect up
	// the results and add to a map

	for i := 0; i < numrepos; i++ {
		<-cr
	}

}

// =============================================================
type repoBranchList struct {
	Repo     string
	Branches []string
}

func cloneRepo(repo repos.OrgRepo, repoChan chan bool) {

	repolocation := "/home/jess/dev/" + repo.FullName
	//fmt.Println(repo.FullName)
	if _, err := os.Stat(repolocation); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", repo.GitURL, repolocation)
		giterr := cmd.Run()
		if giterr != nil {
			log.Fatal(giterr)
		}
		fmt.Printf("%s cloned\n", repo.FullName)
	} else {
		cmd := exec.Command("git", "-C", repolocation, "fetch")
		giterr := cmd.Run()
		if giterr != nil {
			log.Fatal(giterr)
		}
		fmt.Printf("%s updated\n", repo.FullName)
	}
	repoChan <- true
}
