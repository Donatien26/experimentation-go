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
	"encoding/json"
	"fmt"
	"onydev/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// kubeconfigCmd represents the kubeconfig command
var kubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "get kubeconfig from kube onboarding",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getKubeConfig()
	},
}

func init() {
	getCmd.AddCommand(kubeconfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kubeconfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kubeconfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type onboardingResponse struct {
	ApiserverURL string
	Token        string
	Namespace    string
	User         string
	ClusterName  string
	Onboarded    string
}

func getKubeConfig() {
	client := utils.GetAuthClient()
	resp, err := client.Get("https://dev.insee.io/api/cluster")

	if err != nil {
		panic(err)
	}

	onboardingResponse := new(onboardingResponse)
	err = json.NewDecoder(resp.Body).Decode(onboardingResponse)

	config := clientcmdapi.NewConfig()
	config.Clusters[onboardingResponse.ClusterName] = &clientcmdapi.Cluster{
		Server: onboardingResponse.ApiserverURL,
	}
	config.AuthInfos[onboardingResponse.ClusterName] = &clientcmdapi.AuthInfo{
		Token: onboardingResponse.Token,
	}
	config.Contexts[onboardingResponse.ClusterName] = &clientcmdapi.Context{
		Cluster:   onboardingResponse.ClusterName,
		AuthInfo:  onboardingResponse.ClusterName,
		Namespace: onboardingResponse.Namespace,
	}
	config.CurrentContext = onboardingResponse.ClusterName
	pathOptions := clientcmd.NewDefaultPathOptions()
	if err := clientcmd.ModifyConfig(pathOptions, *config, true); err != nil {
		fmt.Println("Unexpected error:" + err.Error())
	}
	f, _ := clientcmd.Write(*config)
	fmt.Printf("%s", f)
	fmt.Println("")
	fmt.Println(color.GreenString("kubeconfig set to follow context: "), onboardingResponse.ClusterName)
}
