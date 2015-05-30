package main

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type tokenSource struct {
	token *oauth2.Token
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	return t.token, nil
}

func createIssue(title, body, owner, repo, token string) {
	ts := &tokenSource{
		&oauth2.Token{AccessToken: token},
	}

	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	issue := github.IssueRequest{
		Title: &title,
		Body:  &body,
	}

	client.Issues.Create(owner, repo, &issue)
}
