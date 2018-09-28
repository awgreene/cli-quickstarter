package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type configElement struct {
	Foo string
	Bar []string
}

type config struct {
	Items []configElement
}

func contains(sArr []string, s string) bool {
	for _, element := range sArr {
		if element == s {
			return true
		}
	}
	return false
}

func main() {
	// Declare and Parse Flags
	debugPtr := flag.Bool("debug", false, "Run with debugging enabled")
	flag.Parse()

	// Get command line argument
	someArg := flag.Arg(0)
	if someArg == "" {
		log.Fatal("App must be provided someArg")
	}

	if *debugPtr {
		log.Printf("Running with Debug enabled")
	}

	var config config
	config.getConf("config.yaml")

	for _, element := range config.Items {
		updatedString := strings.Replace(element.Bar[2], "{{SOME_STRING}}", someArg, -1)
		if *debugPtr {
			log.Printf("Original string: %s", element.Bar[2])
			log.Printf("Updated  string: %s", updatedString)
		}
	}
}

func (c *config) getConf(configFilePath string) *config {
	log.Printf("Getting Config Information from file \"%s\"", configFilePath)
	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		os.Exit(1)
	}

	return c
}
