package main

import (
  "fmt"
  "strings"
  "net/http"
  "github.com/forkfork/ghettostories/github"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
  urlParts := strings.Split(r.URL.Path[1:], "/")
  if (len(urlParts) != 2) {
    errMsg := "503 Bad Request (should be /iteration/stream)"
    http.Error(w, errMsg, http.StatusBadRequest)
    return
  }
  iteration, stream := urlParts[0], urlParts[1]
  issues := github.GetIssues(iteration, stream)
  for _, issue := range issues {
    //complete, incomplete, unknown := parseIssueBody(issueBody)
    fmt.Fprintf(w, "%s done %d todo %d unknown %d\n", issue.Name, issue.Complete, issue.Incomplete, issue.Unknown)
  }
}


func main() {
  // get request parameters
  // request issues descriptions from github => returns list of names, blurbs
  // parse blurbs to get progress

  http.HandleFunc("/", handleRequest)
  http.ListenAndServe(":8080", nil)
}
