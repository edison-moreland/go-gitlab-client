package cmd

import (
	"testing"

	"github.com/edison-moreland/go-gitlab-client/v3/test"
)

func TestRootCmd(t *testing.T) {
	test.RunCommandTestCases(t, "users", []*test.CommandTestCase{
		{
			[]string{},
			nil,
			//configs["default"],
			"help",
			false,
			nil,
		},
		{
			[]string{"help"},
			nil,
			//configs["default"],
			"help",
			false,
			nil,
		},
	})
}
