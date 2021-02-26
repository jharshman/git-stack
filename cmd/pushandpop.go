package cmd

import (
	"errors"
	"fmt"
	"git-stack/branchstack"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "git-stack",
		Short: "Keep track of WIP git branches in a stack.",
	}
	pushCmd = &cobra.Command{
		Use:    "push",
		Short:  "Push a branch name onto the stack.",
		PreRun: loadFile,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("missing branch")
			}
			branchstack.PushBranch(args[0])
			return nil
		},
		PostRun: writeFile,
		SilenceUsage: true,
	}
	popCmd = &cobra.Command{
		Use:    "pop",
		Short:  "Pop a branch name from the stack.",
		PreRun: loadFile,
		Run: func(cmd *cobra.Command, args []string) {
			item := branchstack.PopBranch()
			fmt.Fprint(os.Stdout, item)
		},
		PostRun: writeFile,
		SilenceUsage: true,
	}
)

func init() {
	rootCmd.AddCommand(pushCmd, popCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func loadFile(cmd *cobra.Command, args []string) {
	branchstack.MustRead()
}

func writeFile(cmd *cobra.Command, args []string) {
	branchstack.MustWrite()
}