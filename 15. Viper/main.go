package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./config/SysRequirement.json")

	// Searches for config file in given paths and reads it
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	fmt.Println(viper.GetFloat64("ram.minimumfreeSpace.size"))
	fmt.Println(viper.GetFloat64("processor.maxUsages.percentage"))

	// To write config to a predefined file (file existing already)
	viper.SetConfigName("writeFile")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")

	// File path can also be set by
	// viper.SetConfigFile("./config/writeFile.json")
	viper.WriteConfig()

	// alternatively can also write to given filepath. Creates the file if it doesn't exist
	viper.SafeWriteConfigAs("./config/anotherWriteFile.json")

	// Read values in struct
	viper.SetConfigFile("./config/env.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	prod := viper.Sub("prod")

	// Unmarshal into struct. Struct fields should match keys from config (case in-sensitive)
	type config struct {
		Host    string
		Port    int
		enabled bool
	}

	var C config

	err := prod.Unmarshal(&C)
	fmt.Println(err)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	fmt.Println(C.Host)
}
