package main

import "fmt"
import "os"
import "net/http"

func

func main() {
	apikey := os.Getenv("GITHUB_APIKEY")
	fmt.Printf("Key is %s\n", apikey)
}
