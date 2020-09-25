package kind

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete kind cluster",
	Long:  ``,
	Run: func(inCmd *cobra.Command, args []string) {

	},
}
