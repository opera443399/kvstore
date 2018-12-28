package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	backends "github.com/opera443399/kvstore/backends"
)

//kvGetValue get key from kvstore
func kvGetValue(key string) (map[string]string, error) {
	log.Printf("[kvstore] %s: get [%s]\n", config.BackendNodes, key)
	storeClient, err := backends.New(config.BackendsConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := storeClient.GetValues(key)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func main() {
	flag.Parse()
	if config.PrintVersion {
		fmt.Printf("demo %s (Git SHA: %s, Go Version: %s)\n", Version, GitSHA, runtime.Version())
		os.Exit(0)
	}
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	data, err := kvGetValue("/kvstore/demo/hello")
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(data) > 0 {
		log.Printf("Client get %d groups data from kvstore:", len(data))
		for k, v := range data {
			log.Printf("\n[-] \tkv.Key: %s \n[-] \tkv.Value: %v\n", k, v)
		}
	} else {
		log.Print("Query path not found in kvstore")
	}

}
