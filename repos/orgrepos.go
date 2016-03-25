package repos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/missjessjenkins/branchlist/config"
)

type OrgRepos struct {
	Repos []OrgRepo
	Next  string
}

type OrgRepo struct {
	Name     string
	FullName string `json:"full_name"`
}

func GetOrgReposFromURL(url string) OrgRepos {
	orgRepos := OrgRepos{}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	jsonblob, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonblob, &orgRepos.Repos)
	if err != nil {
		fmt.Println("error:", err)
	}

	link := res.Header.Get("Link")
	re := regexp.MustCompile("<([^>]*)>; rel=\"next\"")
	parts := re.FindStringSubmatch(link)
	if len(parts) > 1 {
		orgRepos.Next = parts[1]
	}

	return orgRepos
}

func GetOrgRepos(org string, c chan OrgRepo) {
	reposurl := "https://api.github.com/orgs/%s/repos?type=all&per_page=50&access_token=%s"
	url := fmt.Sprintf(reposurl, org, config.ApiKey)

	for url != "" {
		orgRepos := GetOrgReposFromURL(url)
		for _, repo := range orgRepos.Repos {
			c <- repo
		}
		url = orgRepos.Next
	}
	close(c)

}
