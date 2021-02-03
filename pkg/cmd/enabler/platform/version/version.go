package version

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
		Use:   "version",
		Short: "Show platform version",
		Long: `Check versions of microservices in git submodules
    You can provide a comma separated list of submodules
    or you can use 'all' for all submodules`,
		Run: func(cmd *cobra.Command, args []string) {
			// initcmd repo
			r, err := git.PlainOpen(repoPath)
			if err != nil {
				log.Fatal(err)
			}
			// fetch repo HEAD
			ref, err := r.Head()
			if err != nil {
				log.Fatal(err)
			}
			// fetch repo config
			c, _ := r.Config()
			if origin, ok := c.Remotes["origin"]; ok {
				log.Infof("Checking platform version for repository (%s@%s)", origin.URLs[0], ref.Hash().String())
			}

			// fetch worktree
			w, err := r.Worktree()
			if err != nil {
				log.Fatal(err)
			}

			var subModulesList git.Submodules
			if subModules == "all" {
				subModulesList, err = w.Submodules()
				if err != nil {
					log.Error(err)
				}
				for _, sub := range subModulesList {
					r, err := sub.Repository()
					if err != nil {
						log.Error(err)
						continue
					}
					ref, err := r.Head()
					if err != nil {
						log.Error(err)
						continue
					}
					log.Infof("%s@%s", sub.Config().URL, ref.Hash().String())
				}
			} else {
				log.Fatal("Not implemented!")
			}
		},
	}
	cwd, _ := os.Getwd()
	cmd.Flags().StringVarP(&repoPath, "repo-path", "", cwd, "repository location")
	cmd.Flags().StringVarP(&subModules, "submodules", "", "all", "submodules to update")
	return cmd
}