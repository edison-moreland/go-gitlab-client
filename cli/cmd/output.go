package cmd

import (
	out "github.com/edison-moreland/go-gitlab-client/v3/cli/output"
	"github.com/edison-moreland/go-gitlab-client/v3/gitlab"
)

func printMeta(meta *gitlab.ResponseMeta, withPagination bool) {
	if verbose {
		out.Meta(meta, withPagination)
	}
}
