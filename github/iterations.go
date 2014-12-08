package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Milestone struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
}

func LookupMilestone(iteration string, username string, password string) (int, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/TabDigital/backlog-api/milestones", nil)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	req.SetBasicAuth(username, password)
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	var milestones []Milestone
	err = json.Unmarshal(contents, &milestones)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s", milestones[0].Title)
	return milestones[0].Number, nil
}
