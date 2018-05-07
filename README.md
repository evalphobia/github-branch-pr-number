github-branch-pr-number
----

[![GoDoc][1]][2] [![License: MIT][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Codecov Coverage][11]][12] [![Go Report Card][13]][14] [![Code Climate][19]][20] [![BCH compliance][21]][22]

[1]: https://godoc.org/github.com/evalphobia/github-branch-pr-number?status.svg
[2]: https://godoc.org/github.com/evalphobia/github-branch-pr-number
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/github-branch-pr-number.svg
[6]: https://github.com/evalphobia/github-branch-pr-number/releases/latest
[7]: https://travis-ci.org/evalphobia/github-branch-pr-number.svg?branch=master
[8]: https://travis-ci.org/evalphobia/github-branch-pr-number
[9]: https://coveralls.io/repos/evalphobia/github-branch-pr-number/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/github-branch-pr-number?branch=master
[11]: https://codecov.io/github/evalphobia/github-branch-pr-number/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/github-branch-pr-number?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/github-branch-pr-number
[14]: https://goreportcard.com/report/github.com/evalphobia/github-branch-pr-number
[15]: https://img.shields.io/github/downloads/evalphobia/github-branch-pr-number/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/github-branch-pr-number/releases
[17]: https://img.shields.io/github/stars/evalphobia/github-branch-pr-number.svg
[18]: https://github.com/evalphobia/github-branch-pr-number/stargazers
[19]: https://codeclimate.com/github/evalphobia/github-branch-pr-number/badges/gpa.svg
[20]: https://codeclimate.com/github/evalphobia/github-branch-pr-number
[21]: https://bettercodehub.com/edge/badge/evalphobia/github-branch-pr-number?branch=master
[22]: https://bettercodehub.com/

`github-branch-pr-number` gets GitHub PullRequst number from branch name through GitHub API.

# Installation

Install github-branch-pr-number using `go get` command:

```bash
$ go get github.com/evalphobia/github-branch-pr-number
```

# Usage

```bash
$ github-branch-pr-number
[ERROR] 2017/05/19 14:21:26 GitHub API token is missing.

usage:
$ github-branch-pr-number -token="<GitHub API Token>" -project="<GitHub project name>" -branch="<branch name>"
ex) github-branch-pr-number -token="this_is_secret" -project="evalphobia/github-branch-pr-number" -branch="develop"

# if https://github.com/evalphobia/github-branch-pr-number/pull/34 is exists and still open, and this Ref branch is same as current HEAD.
$ github-branch-pr-number -token="this_is_secret" -project="evalphobia/github-branch-pr-number" -branch=`git rev-parse --abbrev-ref HEAD`

34
```


# But What For?

For Pull Request review comment by bot.

[reviewdog](https://github.com/haya14busa/reviewdog) needs PR number(CI_PULL_REQUEST env variable).
But some CI or your original CI might not conatain PR number info.

If you cannot get PR number but want to send review comment to GitHub by Bot/CI/etc, `github-branch-pr-number` might help.


# Parameters

## flag option

| name | required | description | example |
| ------- | ------- | ------- | ------- |
| `-token`  | yes | GitHub API Token to perform API request | get this from here: https://github.com/settings/tokens |
| `-branch`  | yes | Git branch name for the Pull Request (Ref) | `master`, `develop`, `$(git rev-parse --abbrev-ref HEAD)` |
| `-project` | * | GitHub project name | `evalphobia/github-branch-pr-number`, `google/go-github`, `github/hubot` |
| `-owner`  | * | GitHub owner name (project is: `evalphobia/github-branch-pr-number`, then owner is: `evalphobia`) | `evalphobia`, `google`, `github` |
| `-repo`  | * | GitHub repository name (project is: `evalphobia/github-branch-pr-number`, then repository is: `github-branch-pr-number`) | `github-branch-pr-number`, `go-github`, `hubot` |
| `-size`  | no | How many latest updated pull requests are searched by this tool (default: 50) | `10`, `30`, `100` |

- `*` At least, either [`-project`] or [`-owner` and `-repo`] is required.


## env var

| name | description |
| ------- | ------- |
| `GITHUB_API_TOKEN` | Used as GitHub API Token if `-token` is not set |
| `REVIEWDOG_GITHUB_API_TOKEN` | (same as above) |
