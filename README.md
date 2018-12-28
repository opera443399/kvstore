# kvstore

## POC

* setup etcd

```bash
ETCDCTL_API=3 /usr/local/bin/etcdctl put /kvstore/demo/hello demo

```

* run

```bash
$ ./bin/demo -node "127.0.0.1:2379"
2018/12/28 11:35:52 KVStore backend set to: etcdv3
2018/12/28 11:35:52 [kvstore] [127.0.0.1:2379]: get [/kvstore/demo/hello]
2018/12/28 11:35:52 Client get 1 groups data from kvstore:
2018/12/28 11:35:52
[-]     kv.Key: /kvstore/demo/hello
[-]     kv.Value: demo

```