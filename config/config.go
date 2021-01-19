package config

import (
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

//Configuration struct for connection
type Configuration struct {
	Port              string
	StaticVariables   string
	Connection_String string
}

//GetConfig
func GetConfig() Configuration {
	basePath, err := os.Getwd()
	configuration := Configuration{}
	err = gonfig.GetConf(basePath+"/config/config.json", &configuration)
	if err != nil {
		panic(err)
	}

	fmt.Println(configuration)

	return configuration

}
