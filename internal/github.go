package internal

import (
	"fmt"
	"io"
	"net/http"
)

type CommitPair struct {
	Owner        string
	Repo         string
	BaseCommit   string
	TargetCommit string
}

func makeReqURL(cp *CommitPair) string {
	return fmt.Sprintf("https://github.com/%s/%s/compare/%s...%s.patch", cp.Owner, cp.Repo, cp.BaseCommit, cp.TargetCommit)
}

func RequestDiff(cp *CommitPair, token string) (string, error) {
	client := &http.Client{}

	url := makeReqURL(cp)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
