// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func ProcessData(data *IssuesSearchResult) map[string][]*Issue {
	hash := make(map[string][]*Issue, 3)
	hash["不到一月"] = make([]*Issue, 0)
	hash["不到一年"] = make([]*Issue, 0)
	hash["超过一年"] = make([]*Issue, 0)

	for _, issue := range data.Items {
		duration := time.Since(issue.CreatedAt)
		if duration.Hours() <= 24*30 {
			hash["不到一月"] = append(hash["不到一月"], issue)
		} else if duration.Hours() < 24*365 { // 假设每年365天
			hash["不到一年"] = append(hash["不到一年"], issue)
		} else {
			hash["超过一年"] = append(hash["超过一年"], issue)
		}
	}
	return hash
}
