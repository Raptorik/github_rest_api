package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}
}
func main() {
	a := http.Client{Timeout: 10 * time.Second}
	GitHubClientIDToken, exists := os.LookupEnv("CLIENT_ID_TOKEN")
	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}
	req, err := http.NewRequest(`GET`, `https://api.github.com/user`, nil)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	req.Header.Add(`Accept`, `application/json`)
	req.Header.Add(`Authorization`, `Bearer `+GitHubClientIDToken)

	//req, err := http.NewRequest(`GET`, `https://api.github.com/users/Raptorik/repos`, nil)
	//if err != nil {
	//fmt.Printf("Error: %s//n", err)
	//return
	//}

	resp, err := a.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Printf("Body: %s\\n", body)
}
