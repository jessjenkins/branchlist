package repos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type OrgRepos struct {
	Repos []OrgRepo
	Next  string
}

type OrgRepo struct {
	Name string
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
