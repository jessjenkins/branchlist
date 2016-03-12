package main

import "fmt"
import "os"

import "net/http"
import "io/ioutil"
import "log"

func main() {
	apikey := os.Getenv("GITHUB_APIKEY")
	org := os.Getenv("GITHUB_ORG")
	fmt.Printf("Org is %s\n", org)
	fmt.Printf("Key is %s\n", apikey)

	reposurl := "https://api.github.com/orgs/%s/repos?type=all&per-page=100&access_token=%s"
	url := fmt.Sprintf(reposurl, org, apikey)
	fmt.Printf("Url=%s\n", url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	json, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", json)
}
