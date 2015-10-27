package config

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateGithub(name string, token string, exclude string) map[string]*Repo {
	repos := make(map[string]*Repo)

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	opt := &github.RepositoryListByOrgOptions{
		Type:        "all",
		ListOptions: github.ListOptions{PerPage: 100},
	}

	for {
		newRepos, resp, err := client.Repositories.ListByOrg(name, opt)
		check(err)

		for _, newRepo := range newRepos {
			url := fmt.Sprintf("git@github.com:%s/%s.git", name, *newRepo.Name)
			repos[*newRepo.Name] = &Repo{
				Url: url,
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}

	return repos
}
