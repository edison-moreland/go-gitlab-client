package cmd

import (
	"github.com/edison-moreland/go-gitlab-client/v3/gitlab"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	RootCmd.AddCommand(ciInfoCmd)
}

var ciInfoCmd = &cobra.Command{
	Use:     "ci-info",
	Aliases: []string{"ci"},
	Short:   "Print information about CI environment",
	Run: func(cmd *cobra.Command, args []string) {
		i, err := gitlab.GetCiInfo()
		if err != nil {
			panic(err)
		}

		i.RenderJson(os.Stdout)
	},
}
