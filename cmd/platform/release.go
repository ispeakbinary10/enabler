package platform

import (
	"github.com/keitaroinc/enabler/cmd/util"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Show releasees",
	Long:  `Release platform by tagging platform repo and
    tagging all individual components (git submodules)
    using their respective SHA that the submodules point at`,
	Run: func(cmd *cobra.Command, args []string) {
		log := util.NewLogger("INFO", nil)
		log.Fatal("Not implemented!")
	},
}
