package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Issue struct {
	Name       string `json:"title"`
	Body       string `json:"body"`
	Complete   int
	Incomplete int
	Unknown    int
}

func parseIssueBody(body string) (int, int, int) {
	incomplete := strings.Count(body, "[ ]")
	complete := strings.Count(body, "[x]")
	unknown := 0
	if incomplete+complete == 0 {
		unknown = 5
	}
	return complete, incomplete, unknown
}

func RequestIssues(milestone int, stream string, username string, password string) ([]Issue, error) {
	issues := []Issue{}
	client := &http.Client{}
	reqUrl := fmt.Sprintf("https://api.github.com/repos/TabDigital/backlog-api/issues?milestone=%d&labels=%s", milestone, stream)
	fmt.Println("requesting %s", reqUrl)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	req.SetBasicAuth(username, password)
	response, err := client.Do(req)
	if err != nil {
    return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
    return nil, err
	}
	err = json.Unmarshal(contents, &issues)
	if err != nil {
    return nil, err
	}
	for i, issue := range issues {
		issues[i].Complete, issues[i].Incomplete, issues[i].Unknown = parseIssueBody(issue.Body)
	}
	return issues, nil
}

func GetIssues(iteration string, stream string, username string, password string) ([]Issue, error) {
	milestone, err := LookupMilestone(iteration, username, password)
	if err != nil {
    return nil, err
	}
	issues, err := RequestIssues(milestone, stream, username, password)
	if err != nil {
    return nil, err
	}
	return issues, nil
}
