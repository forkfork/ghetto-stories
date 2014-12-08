package main

import (
	"fmt"
	"github.com/forkfork/ghetto-stories/github"
	"net/http"
	"strings"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path[1:], "/")
	if len(urlParts) != 2 {
		errMsg := "503 Bad Request (should be /iteration/stream)"
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	iteration, stream := urlParts[0], urlParts[1]
	issues, err := github.GetIssues(iteration, stream, "dummyusername", "dummypassword")
  if (err != nil) {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
  }
	for _, issue := range issues {
		fmt.Fprintf(w, "%s done %d todo %d unknown %d\n", issue.Name, issue.Complete, issue.Incomplete, issue.Unknown)
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
