package cmd

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/v3/test"
)

func TestRmCmd(t *testing.T) {
	test.RunCommandTestCases(t, "users", []*test.CommandTestCase{
		{
			[]string{"rm"},
			nil,
			//configs["default"],
			"rm_help",
			false,
			nil,
		},
		{
			[]string{"rm", "--help"},
			nil,
			//configs["default"],
			"rm_help",
			false,
			nil,
		},
	})
}
