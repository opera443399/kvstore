package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	backends "github.com/opera443399/kvstore/backends"
)

//BackendsConfig kvstore config
type BackendsConfig = backends.Config

// A Config structure is used to configure app.
type Config struct {
	BackendsConfig
	PrintVersion bool
}

var config Config

func processEnv() {
	nodes := os.Getenv("CICD_BACKEND_NODES")
	if len(nodes) > 0 && len(config.BackendNodes) == 0 {
		config.BackendNodes = []string{nodes}
	}

	cakeys := os.Getenv("CICD_CLIENT_CAKEYS")
	if len(cakeys) > 0 && config.ClientCaKeys == "" {
		config.ClientCaKeys = cakeys
	}

	cert := os.Getenv("CICD_CLIENT_CERT")
	if len(cert) > 0 && config.ClientCert == "" {
		config.ClientCert = cert
	}

	key := os.Getenv("CICD_CLIENT_KEY")
	if len(key) > 0 && config.ClientKey == "" {
		config.ClientKey = key
	}
}

// initConfig initializes the configuration
func initConfig() error {
	// Update config from environment variables.
	processEnv()

	if len(config.BackendNodes) == 0 {
		switch config.Backend {
		case "etcdv3":
			config.BackendNodes = []string{"127.0.0.1:2379"}

		}
	}
	// Initialize the storage client
	log.Print("KVStore backend set to: " + config.Backend)

	return nil
}

func init() {
	flag.StringVar(&config.Backend, "backend", "etcdv3", "backend to use")
	flag.BoolVar(&config.BasicAuth, "basic-auth", false, "Use Basic Auth to authenticate (only used with -backend=consul and -backend=etcd)")
	flag.StringVar(&config.ClientCaKeys, "client-ca-keys", "", "client ca keys")
	flag.StringVar(&config.ClientCert, "client-cert", "", "the client cert")
	flag.StringVar(&config.ClientKey, "client-key", "", "the client key")
	flag.Var(&config.BackendNodes, "node", "list of backend nodes")
	flag.StringVar(&config.Username, "username", "", "the username to authenticate as (only used with vault and etcd backends)")
	flag.StringVar(&config.Password, "password", "", "the password to authenticate with (only used with vault and etcd backends)")
	flag.BoolVar(&config.PrintVersion, "version", false, "print version and exit")

	flag.Parse()
	if config.PrintVersion {
		fmt.Printf("kvstore %s (Git SHA: %s, Go Version: %s)\n", Version, GitSHA, runtime.Version())
		os.Exit(0)
	}
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}
}

//KVRead read from kvstore
func KVRead(key string) (map[string]string, error) {
	if key[:1] != "/" {
		appEnv := os.Getenv("DEPLOY_ENV")
		if len(appEnv) == 0 {
			appEnv = AppEnv
		}
		key = fmt.Sprintf("/%s/%s/%s", AppName, appEnv, key)
	}
	storeClient, err := backends.New(config.BackendsConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := storeClient.GetValues(key)
	if err != nil {
		return nil, err
	}
	//log.Printf("Got the following map from kvstore: %v", result)
	return result, nil
}
