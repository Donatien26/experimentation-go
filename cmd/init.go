/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"onydev/utils"
	"os"
	"path/filepath"

	oauth2ns "github.com/nmrshll/oauth2-noserver"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientID, _ := cmd.Flags().GetString("clientID")
		authURL, _ := cmd.Flags().GetString("authURL")
		tokenURL, _ := cmd.Flags().GetString("tokenURL")
		execute(clientID, authURL, tokenURL)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("clientID", "", "onboarding", "clientID")
	initCmd.Flags().StringP("authURL", "", "https://keycloak.dev.insee.io/auth/realms/dev/protocol/openid-connect/auth", "auth url")
	initCmd.Flags().StringP("tokenURL", "", "https://keycloak.dev.insee.io/auth/realms/dev/protocol/openid-connect/token", "tokenURL")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func execute(clientID string, authURL string, tokenURL string) {
	create := createConfigFile()
	if create {
		viper.Set("auth.conf.clientID", clientID)
		viper.Set("auth.conf.authURL", authURL)
		viper.Set("auth.conf.tokenURL", tokenURL)
		viper.WriteConfig()

		conf := utils.GenerateOauthConfig()

		client, err := oauth2ns.AuthenticateUser(conf)
		if err != nil {
			log.Fatal(err)
		}
		utils.SaveToken(client.Token)
	}
}

func createConfigFile() bool {
	home, _ := homedir.Dir()
	path := filepath.Join(home, ".onydev")
	filePath := filepath.Join(home, ".onydev", "config.yaml")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(color.GreenString("A config file will be created at " + filePath))
		os.MkdirAll(path, os.ModePerm)
		var file, _ = os.Create(filePath)
		defer file.Close()
		fmt.Println(color.GreenString("==> done creating file", filePath))
		return true
	}
	fmt.Println(color.RedString("A file already exist at " + filePath))
	return false
}
