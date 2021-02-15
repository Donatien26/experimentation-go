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
	"time"

	"github.com/fatih/color"
	oauth2ns "github.com/nmrshll/oauth2-noserver"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := utils.GenerateOauthConfig()
		if time.Now().Before(utils.GetLastToken().Expiry) {
			client, err := oauth2ns.AuthenticateUser(conf)
			if err != nil {
				log.Fatal(err)
			}
			// use client.Get / client.Post for further requests, the token will automatically be there
			utils.SaveToken(client.Token)
		}
		fmt.Println(color.GreenString("accessToken: ") + viper.GetString("auth.accessToken"))
	},
}

func init() {
	getCmd.AddCommand(tokenCmd)
}
