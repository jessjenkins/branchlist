package main

import (
	"fmt"
	"sort"

	"github.com/jessjenkins/branchlist/config"
	"github.com/jessjenkins/branchlist/repos"
)

var apikey string

func main() {

	config.Setup()

	cr := make(chan repoBranchList)

	// For each repo on github, fire off a routine to collect
	// all the branches for that repo
	c := make(chan repos.OrgRepo)
	go repos.GetOrgRepos(config.Org, c)
	numrepos := 0
	for repo := range c {
		go collectBranchesForRepo(repo, cr)
		numrepos++
	}

	// When each fired off routine above finishes, collect up
	// the results and add to a map
	branchmap := make(map[string][]string)
	for i := 0; i < numrepos; i++ {
		rbl := <-cr
		//fmt.Printf("%d - %s\n", i, rbl.Repo)
		for _, br := range rbl.Branches {
			//fmt.Printf("  %d - %s\n", x, br)
			branchmap[br] = append(branchmap[br], rbl.Repo)
		}
	}

	// Sort and print out list of repos per branchname
	branches := make([]string, 0, len(branchmap))
	for branch := range branchmap {
		branches = append(branches, branch)
	}
	sort.Strings(branches)

	for _, branch := range branches {
		fmt.Printf("%s: %v\n", branch, branchmap[branch])
	}
}

// =============================================================
type repoBranchList struct {
	Repo     string
	Branches []string
}

func collectBranchesForRepo(repo repos.OrgRepo, repoChan chan repoBranchList) {
	rbl := repoBranchList{repo.Name, nil}

	c := make(chan string)
	go repos.GetRepoBranches(repo.FullName, c)

	for branch := range c {
		rbl.Branches = append(rbl.Branches, branch)
	}

	repoChan <- rbl
}
