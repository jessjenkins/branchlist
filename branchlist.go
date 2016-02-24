package main

import "fmt"
import "os"


func main() {
	apikey := os.Getenv("GITHUB_APIKEY")
	fmt.Printf("Key is %s\n", apikey)
}
