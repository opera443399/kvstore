package backends

import (
	"errors"

	"github.com/opera443399/kvstore/backends/etcdv3"
)

// The StoreClient interface is implemented by objects that can retrieve
// key/value pairs from a backend store.
type StoreClient interface {
	GetValues(key string) (map[string]string, error)
}

// New is used to create a storage client based on our configuration.
func New(config Config) (StoreClient, error) {
	backendNodes := config.BackendNodes
	switch config.Backend {
	case "etcdv3":
		return etcdv3.NewEtcdClient(backendNodes, config.ClientCert, config.ClientKey, config.ClientCaKeys, config.BasicAuth, config.Username, config.Password)

	}
	return nil, errors.New("Invalid backend")
}