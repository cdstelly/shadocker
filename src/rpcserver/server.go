package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"rpcshared"
	"time"

	// Log results to our webserver
	"rpclogger"
)

var (
	MyName     string
	BrokerHost string
	MyType     string
)

func init() {
	MyName = Generate(2, "-")
	BrokerHost = os.Getenv("BROKERHOST")
	if len(BrokerHost) == 0 {
		BrokerHost = "trex1:5050"
	}
	MyType = "SHA1"
}

func startServer() {
	on := new(rpcshared.SHA1)
	rpc.Register(on)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":5554")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	go http.Serve(l, nil)
//	go PeriodicUpdate(on)
}

// Send the whole request history periodically
// TODO: Decay the RequestHistory buffer. This struct will eventually get huge..
func PeriodicUpdate(myRPCInstance *rpcshared.SHA1) {
	for {
		time.Sleep(time.Millisecond * 5000)
		rpclogger.SubmitReport(BrokerHost, MyName, MyType, myRPCInstance.RequestHistory)
	}
}

//Start the server, listen forever.
func main() {
	startServer()
	fmt.Println("[*] Server started. \tMy name:", MyName, "\tBrokerHost: ", BrokerHost)

	meta := make(chan int)
	x := <-meta /// wait for a while, and listen
	fmt.Println(x)
}
