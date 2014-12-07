package github

import (
  "strings"
)

type Issue struct {
  Name string
  Body string
  Complete int
  Incomplete int
  Unknown int
}

func parseIssueBody(body string) (int, int, int) {
  incomplete := strings.Count(body, "[ ]")
  complete := strings.Count(body, "[x]")
  unknown := 0
  if (incomplete + complete == 0) {
    unknown = 5
  }
  return complete, incomplete, unknown
}

func GetIssues (iteration string, stream string) []Issue {
  issues := []Issue{}
  issues = append(issues, Issue{Name: "fix bla bla", Body: "fdsfsdf [x] [ ]", Complete: 1, Incomplete: 1, Unknown: 0})
  issues = append(issues, Issue{Name: "make bla bla", Body: "fdgdfgsf", Complete: 0, Incomplete: 0, Unknown: 5})
  return issues
}

