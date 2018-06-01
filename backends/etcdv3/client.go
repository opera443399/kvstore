package etcdv3

import (
	"log"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	"context"

	"github.com/coreos/etcd/clientv3"
)

// Client is a wrapper around the etcd client
type Client struct {
	client *clientv3.Client
}

// NewEtcdClient returns an *etcdv3.Client with a connection to named machines.
func NewEtcdClient(machines []string, cert, key, caCert string, basicAuth bool, username string, password string) (*Client, error) {
	cfg := clientv3.Config{
		Endpoints:            machines,
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    10 * time.Second,
		DialKeepAliveTimeout: 3 * time.Second,
	}

	if basicAuth {
		cfg.Username = username
		cfg.Password = password
	}

	tlsEnabled := false
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
	}

	if caCert != "" {
		certBytes, err := ioutil.ReadFile(caCert)
		if err != nil {
			return &Client{}, err
		}

		caCertPool := x509.NewCertPool()
		ok := caCertPool.AppendCertsFromPEM(certBytes)

		if ok {
			tlsConfig.RootCAs = caCertPool
		}
		tlsEnabled = true
	}

	if cert != "" && key != "" {
		tlsCert, err := tls.LoadX509KeyPair(cert, key)
		if err != nil {
			return &Client{}, err
		}
		tlsConfig.Certificates = []tls.Certificate{tlsCert}
		tlsEnabled = true
	}

	if tlsEnabled {
		cfg.TLS = tlsConfig
	}
	
	client, err := clientv3.New(cfg)
	if err != nil {
		return &Client{}, err
	}
	
	return &Client{client}, nil
}

var requestTimeout = time.Duration(3) * time.Second

// GetValues query etcd for prefix
func (c *Client) GetValues(prefix string) (map[string]string, error) {
	vars := make(map[string]string)
	defer c.client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := c.client.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range resp.Kvs {
		//log.Printf("Query etcd for prefix:\n[-] \tkv.Key: %s \n[-] \tkv.Value: %v\n", kv.Key, string(kv.Value))
		vars[string(kv.Key)] = string(kv.Value)
	}
	return vars, nil
}
