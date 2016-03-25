package main

import "fmt"
import "os"

import "github.com/missjessjenkins/branchlist/repos"

func main() {
	apikey := os.Getenv("GITHUB_APIKEY")
	org := os.Getenv("GITHUB_ORG")
	fmt.Printf("Org is %s\n", org)
	fmt.Printf("Key is %s\n\n", apikey)

	reposurl := "https://api.github.com/orgs/%s/repos?type=all&per_page=50&access_token=%s"
	url := fmt.Sprintf(reposurl, org, apikey)

	getAllRepos(url, 0)

}

func getAllRepos(url string, z int) {
	orgRepos := repos.GetOrgReposFromURL(url)
	//fmt.Printf("orgRepos=%s\n", orgRepos)

	newz := z
	for i, repo := range orgRepos.Repos {
		fmt.Printf("%d - %s\n", z+i, repo.Name)
		newz++
	}
	fmt.Println("----")
	if orgRepos.Next != "" {
		getAllRepos(orgRepos.Next, newz)
	}
}
