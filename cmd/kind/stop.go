package kind

import "github.com/spf13/cobra"

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop kind cluster",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

