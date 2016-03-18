package repos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type OrgRepos struct {
	Repos []OrgRepo
	Next  string
	Last  string
}

type OrgRepo struct {
	Name string
}

func GetOrgRepos(url string) OrgRepos {
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

	for _, v := range strings.Split(link, ",") {
		re := regexp.MustCompile("<(.*)>; rel=\"(.*)\"")
		parts := re.FindStringSubmatch(v)
		linkurl := parts[1]
		linktype := parts[2]

		switch linktype {
		case "last":
			orgRepos.Last = linkurl
		case "next":
			orgRepos.Next = linkurl
		}
	}

	return orgRepos
}
