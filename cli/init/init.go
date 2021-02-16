package init

import (
	"fmt"
	"log"
	"onydev/utils"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	oauth2ns "github.com/nmrshll/oauth2-noserver"
	"github.com/spf13/viper"
)

//Execute ...
func Execute(clientID string, realm string, keycloakURL string, onboardingURL string, updateFile bool) {
	create := createConfigFile()
	if create || updateFile {
		viper.Set("auth.conf.clientID", clientID)
		viper.Set("auth.conf.keycloakURL", keycloakURL)
		viper.Set("auth.conf.realm", realm)
		viper.Set("onboardingURL", onboardingURL)
		viper.WriteConfig()

		conf := utils.GenerateOauthConfigFromParams(clientID, realm, keycloakURL)

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
