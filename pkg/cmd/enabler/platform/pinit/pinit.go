package pinit

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"os"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	var repoPath, subModules string
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize platform",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// init repo
			r, err := git.PlainOpen(repoPath)
			if err != nil {
				log.Fatal(err)
			}
			// initcmd worktree
			w, err := r.Worktree()
			if err != nil {
				log.Fatal(err)
			}
			// update submodules
			var subModulesList git.Submodules
			if subModules == "all" {
				subModulesList, err = w.Submodules()
				if err != nil {
					log.Error(err)
				}
				for _, sub := range subModulesList {
					log.Infof("Fetching latest changes for (%s)", sub.Config().Name)
					r, err := sub.Repository()
					if err != nil {
						log.Error(err)
						continue
					}
					w, err := r.Worktree()
					if err != nil {
						log.Error(err)
						continue
					}
					_ = w.Pull(&git.PullOptions{RemoteName: "origin"})
				}
			}
		},
	}

	cwd, _ := os.Getwd()
	cmd.Flags().StringVarP(&repoPath, "repo-path", "", cwd, "repository location")
	cmd.Flags().StringVarP(&subModules, "submodules", "", "all", "submodules to update")
	return cmd
}