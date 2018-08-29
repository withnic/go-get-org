package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

const (
	baseURL = "https://api.github.com/orgs/%s/repos?access_token=%s"
)

type Repos []struct {
	FullName string `json:"full_name"`
}

func main() {
	app := cli.NewApp()
	app.Name = "go-get-org"
	app.Usage = "git commit messages with emoji"
	app.Action = GetRepos

	app.Run(os.Args)
}

func GetRepos(c *cli.Context) error {
	args := c.Args()
	if len(args) != 2 {
		log.Fatal("should be 2 args")
	}
	org := args[0]
	token := args[1]
	url := fmt.Sprintf(baseURL, org, token)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	repos := Repos{}
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		log.Fatal(err)
	}

	names := make([]string, 0, len(repos))
	for _, repo := range repos {
		names = append(names, repo.FullName)
	}

	for _, name := range names {
		path := fmt.Sprintf("github.com/%s/...", name)
		fmt.Printf("go get -u %s\n", path)
		out, err := exec.Command("go", "get", "-u", path).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(out)
	}

	fmt.Println("done")
	return nil
}