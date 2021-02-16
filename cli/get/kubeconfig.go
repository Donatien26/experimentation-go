package get

import (
	"encoding/json"
	"fmt"
	"net/http"
	"onydev/utils"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/viper"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

//OnboardingResponse ...
type OnboardingResponse struct {
	ApiserverURL string
	Token        string
	Namespace    string
	User         string
	ClusterName  string
	Onboarded    string
}

//ExecuteGetKubeConfig ...
func ExecuteGetKubeConfig(group string, write bool, destination string) {
	var resp *http.Response
	var err error
	client := utils.GetAuthClient()
	if group == "" {
		resp, err = client.Get(viper.GetString("onboardingURL") + "/api/cluster")
	} else {
		resp, err = client.Get(viper.GetString("onboardingURL") + "/api/cluster/credentials/" + group)
	}

	if err != nil {
		panic(err)
	}

	onboardingResponse := new(OnboardingResponse)
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

	if write {
		if destination != "" {
			pathOptions.GlobalFile = filepath.Join(filepath.Clean(destination), "kubeconfig")
		} else {
			fmt.Println(color.GreenString("kubeconfig set to follow context: "), onboardingResponse.ClusterName)
		}
		if err := clientcmd.ModifyConfig(pathOptions, *config, true); err != nil {
			fmt.Println("Unexpected error:" + err.Error())
		}

	}
	f, _ := clientcmd.Write(*config)
	fmt.Printf("%s", f)
}
