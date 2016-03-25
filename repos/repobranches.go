package repos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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
