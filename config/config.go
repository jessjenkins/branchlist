package config

import "os"

var ApiKey, Org string

func Setup() {
	ApiKey = os.Getenv("GITHUB_APIKEY")
	Org = os.Getenv("GITHUB_ORG")
	//fmt.Printf("Org is %s\n", Org)
	//fmt.Printf("Key is %s\n\n", ApiKey)
}
