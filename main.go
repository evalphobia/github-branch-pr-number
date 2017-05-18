package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	// under score variables are used as flag value.
	_verbose bool
	_token   string
	_branch  string
	_project string
	_owner   string
	_repo    string
	_size    int
)

var usage = `
usage:
$ github-branch-pr-number -token="<GitHub API Token>" -project="<GitHub project name>" -branch="<branch name>"
ex) github-branch-pr-number -token="this_is_secret" -project="evalphobia/github-branch-pr-number" -branch="develop"
`

func init() {
	flag.BoolVar(&_verbose, "v", _verbose, "output verbose log")
	flag.StringVar(&_token, "token", _token, "set GitHub API Token (or set GITHUB_API_TOKEN env var)")
	flag.StringVar(&_project, "project", _project, "GitHub project name (e.g. 'evalphobia/github-branch-pr-number')")
	flag.StringVar(&_owner, "owner", _owner, "GitHub owner name (e.g. 'evalphobia')")
	flag.StringVar(&_repo, "repo", _repo, "GitHub repository name (e.g. 'github-branch-pr-number')")
	flag.StringVar(&_branch, "branch", _branch, "branch name (e.g. 'master', 'develop', 'feature/fix-bug')")
	flag.IntVar(&_size, "size", 50, "how many pull requests to search latest update")
}

func main() {
	parseFlag()

	err := validateFlag()
	if err != nil {
		exitWithError(err)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: _token},
	)
	tc := oauth2.NewClient(ctx, ts)
	cli := github.NewClient(tc)
	opt := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{PerPage: _size},
		Sort:        "updated",
		Direction:   "desc",
	}

	list, _, err := cli.PullRequests.List(ctx, _owner, _repo, opt)
	if err != nil {
		exitWithError(err)
	}

	for _, pr := range list {
		if pr.Head == nil {
			continue
		}
		if pr.Head.Ref == nil {
			continue
		}
		if *pr.Head.Ref == _branch {
			n := pr.GetNumber()
			if n != 0 {
				fmt.Printf("%d", n)
				os.Exit(0)
			}
		}
	}
	os.Exit(1)
}

// parseFlag parses command line flag options.
func parseFlag() {
	flag.Parse()

	if _verbose  {
		enableInfoLog()
		loggingInfo("enabled info log")
	}
	if _project != "" {
		s := strings.Split(_project, "/")
		if len(s) == 2 {
			_owner = s[0]
			_repo = s[1]
		}
	}

	if _token == "" {
		if v := os.Getenv("GITHUB_API_TOKEN"); v != "" {
			_token = v
		} else if v := os.Getenv("REVIEWDOG_GITHUB_API_TOKEN"); v != ""  {
			_token = v
		}
	}
}

// validateFlag validates flag values.
func validateFlag() (err error) {
	if _token == "" {
		return fmt.Errorf("GitHub API token is missing.\n%s", usage)
	}
	if _owner == "" {
		return fmt.Errorf("GitHub owner is missing.\n%s", usage)
	}
	if _repo == "" {
		return fmt.Errorf("GitHub repository is missing.\n%s", usage)
	}
	if _branch == "" {
		return fmt.Errorf("branch name is missing.\n%s", usage)
	}
	return nil
}

// exitWithError outputs error log and exits with error code.
func exitWithError(err error) {
	loggingError(err.Error())
	os.Exit(2)
}
