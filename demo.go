package main

import (
	"log"
)

//AppName in path: /kvstore/ns-dev
var AppName = "kvstore"
//AppEnv in path: /kvstore/ns-dev
var AppEnv = "ns-dev"

func main(){
	//read from path: /kvstore/ns-dev/demo
	data, err := KVRead("demo/token")
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