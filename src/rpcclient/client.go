package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"
	"rpcshared"

)

func main() {
	client, err := rpc.DialHTTP("tcp", "0.0.0.0:5554")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	filepath := "linux.jpg"
	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("error reading file: ", err)
	}

	args := &rpcshared.Args{DataID: "test", Data: fileData}
	var reply string
	err = client.Call("SHA1.Evaluate", args, &reply)
	if err != nil {
		log.Fatal("sha error:", err)
	}
	fmt.Printf("Result: %s\n", reply)

}
