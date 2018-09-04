package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

const (
	perpage = 100
)

const (
	timeout = 30 * time.Second
)

func main() {
	app := cli.NewApp()
	app.Name = "go-get-org"
	app.Usage = "go get all repositories of the organization"
	app.Action = Run

	app.Run(os.Args)
}

func Run(c *cli.Context) error {
	args := c.Args()
	if len(args) != 2 {
		log.Fatal("args must be organization and token")
	}
	org := args[0]
	token := args[1]

	repos, err := getRepos(org, token)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := goGetRepos(repos); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}

func goGetRepos(repos []*github.Repository) error {
	out := make(chan string)
	rest := make([]string, 0, len(repos))
	for _, repo := range repos {
		go func(repo *github.Repository) error {
			name := filepath.Join("github.com", *repo.FullName, "...")
			o, err := exec.Command("go", "get", "-u", name).Output()
			if err != nil {
				return err
			}
			out <- string(o)
			return nil
		}(repo)
		select {
		case <-out:
			fmt.Printf("Installing %s SUCCEEDED\n", *repo.FullName)
		case <-time.After(timeout):
			fmt.Printf("Installing %s timeout\n", *repo.FullName)
			rest = append(rest, *repo.FullName)
		}
	}

	fmt.Printf("\ninstalled repositories: %d\n", len(repos)-len(rest))
	fmt.Printf("not installed repositories: %d\n", len(rest))
	if len(rest) == 0 {
		return nil
	}

	fmt.Printf("\nthe following repositories not installed\n")
	for _, r := range rest {
		fmt.Println(r)
	}
	return nil
}

func getRepos(org, token string) ([]*github.Repository, error) {
	client, err := NewClient(token)
	if err != nil {
		return nil, err
	}

	var repos []*github.Repository
	page := 1
	for {
		rs, _, err := client.Repositories.ListByOrg(
			context.Background(),
			org,
			&github.RepositoryListByOrgOptions{
				ListOptions: github.ListOptions{
					Page:    page,
					PerPage: perpage,
				},
			})
		if err != nil {
			return nil, err
		}
		if len(rs) == 0 {
			break
		}
		for _, r := range rs {
			repos = append(repos, r)
		}
		page++
	}

	fmt.Printf("installing the following repositories (%d)\n", len(repos))
	for _, repo := range repos {
		fmt.Println(*repo.FullName)
	}
	return repos, nil
}

func NewClient(token string) (*github.Client, error) {
	if len(token) == 0 {
		return nil, errors.New("token is missing")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return github.NewClient(tc), nil
}
