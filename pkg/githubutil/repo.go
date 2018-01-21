package githubutil

import (
	"fmt"
	"strings"

	"sourcegraph.com/sourcegraph/sourcegraph/pkg/api"
)

// SplitRepoURI splits a string like "github.com/alice/myrepo" to "alice" and
// "myrepo".
func SplitRepoURI(input api.RepoURI) (owner, repo string, err error) {
	uri := string(input)

	// HACK: treat sourcegraph.com/... as github.com/...
	if strings.HasPrefix(uri, "sourcegraph.com/") {
		uri = strings.Replace(uri, "sourcegraph.com/", "github.com/", 1)
	}

	if strings.HasPrefix(uri, "github.com/") {
		uri = strings.TrimPrefix(uri, "github.com/")
	} else {
		return "", "", fmt.Errorf("not a GitHub repository URI: %q", uri)
	}

	parts := strings.Split(uri, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid GitHub repository owner/repo string: %q", uri)
	}

	owner, repo = parts[0], parts[1]
	if owner == "" || repo == "" {
		return "", "", fmt.Errorf("invalid GitHub owner or repo in %q", uri)
	}

	return owner, repo, nil
}
