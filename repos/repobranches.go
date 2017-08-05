package repos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/jessjenkins/branchlist/config"
)

type RepoBranches struct {
	Branches []RepoBranch
	Next     string
}

type RepoBranch struct {
	Name string
}

func GetRepoBranchesFromURL(url string) RepoBranches {
	repoBranches := RepoBranches{}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	jsonblob, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonblob, &repoBranches.Branches)
	if err != nil {
		fmt.Println("error:", err)
	}

	link := res.Header.Get("Link")
	re := regexp.MustCompile("<([^>]*)>; rel=\"next\"")
	parts := re.FindStringSubmatch(link)
	if len(parts) > 1 {
		repoBranches.Next = parts[1]
	}

	return repoBranches
}

func GetRepoBranches(repo string, c chan string) {
	branchesurl := "https://api.github.com/repos/%s/branches?per_page=100&access_token=%s"
	url := fmt.Sprintf(branchesurl, repo, config.ApiKey)

	for url != "" {
		repoBranches := GetRepoBranchesFromURL(url)
		for _, branch := range repoBranches.Branches {
			c <- branch.Name
		}
		url = repoBranches.Next
	}
	close(c)

}
