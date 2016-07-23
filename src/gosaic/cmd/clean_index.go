package cmd

import (
	"gosaic/controller"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CleanIndexCmd)
}

var CleanIndexCmd = &cobra.Command{
	Use:   "clean_index",
	Short: "Remove stale index entries",
	Long:  "Remove stale index entries",
	Run: func(c *cobra.Command, args []string) {
		err := Env.Init()
		if err != nil {
			Env.Fatalf("Unable to initialize environment: %s\n", err.Error())
		}
		defer Env.Close()

		controller.CleanIndex(Env)
	},
}