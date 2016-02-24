package main

import "fmt"
import "os"


func main() {
	apikey := os.Getenv("GITHUB_APIKEY")
	org := os.Getenv("GITHUB_ORG")
	fmt.Printf("Org is %s\n", org)
	fmt.Printf("Key is %s\n", apikey)
	reposurl := "https://api.github.com/orgs/%s/repos?type=all&per-page=100&access_token=%s"
	url := fmt.Sprintf(reposurl, org, apikey)
	fmt.Printf("Url=%s\n", url)
}
