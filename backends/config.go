package backends

import (
	"fmt"
)


// Nodes is a custom flag Var representing a list of etcd nodes.
type Nodes []string

// String returns the string representation of a node var.
func (n *Nodes) String() string {
	return fmt.Sprintf("%s", *n)
}

// Set appends the node to the etcd node list.
func (n *Nodes) Set(node string) error {
	*n = append(*n, node)
	return nil
}

//Config kvstore client config
type Config struct {
	Backend      string     `toml:"backend"`
	BasicAuth    bool       `toml:"basic_auth"`
	ClientCaKeys string     `toml:"client_cakeys"`
	ClientCert   string     `toml:"client_cert"`
	ClientKey    string     `toml:"client_key"`
	BackendNodes Nodes      `toml:"nodes"`
	Username     string     `toml:"username"`
	Password     string     `toml:"password"`
}