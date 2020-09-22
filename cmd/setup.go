/*
Copyright © 2020 Keitaro Inc dev@keitaro.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// create dirs
		path, err := os.Getwd()
		if err != nil {
			log.Println("error msg", err)
		}
		binPath := filepath.Join(path, "bin")
		if _, err := os.Stat(binPath); os.IsNotExist(err) {
			os.Mkdir(binPath, 0755)
		}

		// kubectl
		kubectl := viper.GetStringMapString("kubectl")
		url := fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/v%s/bin/%s/amd64/kubectl", kubectl["version"], runtime.GOOS)
		err = Download("bin/kubectl", url)
		if err != nil {
			fmt.Println("SOME ERROR")
		} else {
			fmt.Println("Kubetl downloaded successfully")
		}
		err = os.Chmod("bin/kubectl", 0755)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("kubectl %s %s✓", kubectl["version"], string(GREEN)))

		helm := viper.GetStringMapString("helm")
		istio := viper.GetStringMapString("istio")
		kind := viper.GetStringMapString("kind")
		skaffold := viper.GetStringMapString("skaffold")

		fmt.Println("Versions", kubectl["version"], helm["version"], istio["version"], kind["version"], skaffold["version"])

		//kind := viper.Get("kind")
		fmt.Println(kind)
		fmt.Println(kind["version"])
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Download(path string, url string) error {
	fmt.Println(fmt.Sprintf("Downloading %s ...", url))
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}