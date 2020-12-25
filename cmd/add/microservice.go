package add

// TODO:
// * Add skaffold configuration for the microservice
// * Add default values for the microservice
// * Provide git url to init submodule

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/keitaroinc/enabler/cmd/util"
	"github.com/spf13/cobra"
)

var templateValues struct {
	Name string
	Port int
}

func writeToFile(dirPath string, tmpl *template.Template) error {
	filePath := filepath.Join(dirPath, tmpl.Name())
	f, err := os.Create(filePath)
	err = tmpl.Execute(f, templateValues)
	f.Sync()
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}

var microserviceCmd = &cobra.Command{
	Use:   "microservice",
	Short: "Add microservice",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		log := util.NewLogger("INFO", nil)
		// Create directories
		log.Info("Before creating dirs")
		dirPath := fmt.Sprintf("infrastructure/%s/templates", templateValues.Name)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			os.MkdirAll(dirPath, 0755)
		}

		tmpl, err := template.New("service.yaml").Delims("[[", "]]").ParseFiles("templates/Chart/templates/service.yaml")
		err = writeToFile(dirPath, tmpl)
		if err != nil {
			log.Fatal(err)
		}

		tmpl, err = template.New("configmap.yaml").Delims("[[", "]]").ParseFiles("templates/Chart/templates/configmap.yaml")
		err = writeToFile(dirPath, tmpl)
		if err != nil {
			log.Fatal(err)
		}

		tmpl, err = template.New("deployment.yaml").Delims("[[", "]]").ParseFiles("templates/Chart/templates/deployment.yaml")
		err = writeToFile(dirPath, tmpl)
		if err != nil {
			log.Fatal(err)
		}

		tmpl, err = template.New("Chart.yaml").Delims("[[", "]]").ParseFiles("templates/Chart/Chart.yaml")
		err = writeToFile(dirPath, tmpl)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// namesoaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pf := microserviceCmd.PersistentFlags()
	pf.StringVarP(&templateValues.Name, "name", "n", "", "The name of the microservice that you want to add")
	pf.IntVarP(&templateValues.Port, "port", "p", 80, "Port on which the process will listen to")
	cobra.MarkFlagRequired(pf, "name")
}
