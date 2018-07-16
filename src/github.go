package src

import (
	"os"

	"github.com/google/go-github/github"
	"github.com/gregjones/httpcache"
)

var (
	GitHub *github.Client
)

func InitGithub() {
	transport := github.BasicAuthTransport{
		Username:  os.Getenv("GITHUB_USERNAME"),
		Password:  os.Getenv("GITHUB_PASSWORD"),
		OTP:       os.Getenv("GITHUB_OTP"),
		Transport: httpcache.NewMemoryCacheTransport(),
	}
	GitHub = github.NewClient(transport.Client())
}
