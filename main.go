package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
)

func main() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(fmt.Sprintf("%s/.config/porecast", homedir))
	viper.SetConfigType("toml")

	if len(os.Args) < 2 {
		fmt.Println("Usage: porecast <COMMAND>\nCommands: init, run")
		os.Exit(0)
	}

	if os.Args[1] == "init" {
		if fileExists(fmt.Sprintf("%s/.config/porecast/config", homedir)) {
			fmt.Printf("Configuration file already exists!")
			os.Exit(0)
		}
		fmt.Printf("Generating configuration file...\n")
		viper.SetDefault("api", "Grab your API key from https://openweathermap.org")
		viper.SetDefault("unit", "Choose between following: metric (Celsius), standard (Kelvin) or imperial (Fahrenheit)")
		viper.SetDefault("longitude", "Grab coordinates from a site like https://latlong.net")
		viper.SetDefault("latitude", "Grab coordinates from a site like https://latlong.net")
		err := os.Mkdir(fmt.Sprintf("%s/.config/porecast", homedir), 0755)
		if err != nil {
			panic(fmt.Errorf("could not create .config/porecast directory: %v", err))
		}
		_, err = os.Create(fmt.Sprintf("%s/.config/porecast/config", homedir))
		if err != nil {
			panic(fmt.Errorf("could not create .config/porecast/config: %v", err))
		}
		err = viper.WriteConfig()
		if err != nil {
			panic(fmt.Errorf("error writing configuration file: %s", err))
		}
		fmt.Printf("Configuration template generated, make sure to edit it before running!\n")
		os.Exit(0)
	}

	if os.Args[1] == "run" {
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error whilst reading configuration file: %s \n", err))
		}
		response, _ := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", viper.Get("latitude"), viper.Get("longitude"), viper.Get("api")))
		body, _ := io.ReadAll(response.Body)
		var m Message
		_ = json.Unmarshal(body, &m)
		fmt.Printf("%d", int(m.Main.Temp))
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(fmt.Sprintf("Error checking if file exists: %s\n", err))
}
