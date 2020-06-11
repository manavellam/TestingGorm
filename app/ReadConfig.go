package app

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/TestingGorm/models"
)

//ReadConfig Reads the configuration values for the API from the JSON config file
func ReadConfig(config *models.Config) {
	configFile, err := ioutil.ReadFile("./app/config.json")
	if err != nil {
		log.Fatal("Unable to read conf file")
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}
}
