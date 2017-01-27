package rpcshared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"time"
)

type SHA1 struct {
	NumberRequests int
	RequestHistory []float64
}

type Args struct {
	Data   []byte
	DataID string
}

func (t *SHA1) Evaluate(args *Args, reply *string) error {
	pathToTool := "/usr/bin/sha1sum"
	fmt.Println("Path to sha1sum: ", pathToTool)

	//Setup the shell command to launch sha1sum
	cmd := exec.Command(pathToTool, "-")
	cmd.Stdin = bytes.NewReader(args.Data)

	//Capture STDOUT
	var out bytes.Buffer
	cmd.Stdout = &out

	// Let's measure execution time:
	startTime := time.Now()

	// Actually run the command:
	err := cmd.Run()
	fmt.Println("[-] Output: ", out.String())

	// Capture duration
	executionTime := time.Since(startTime).Seconds() //use seconds as opposed to nanoseconds, returns float64 which is required with stats package
	t.NumberRequests += 1
	t.RequestHistory = append(t.RequestHistory, executionTime)

	//Post process the output
	jsonMapping := make(map[string]string)
	jsonKey := "sha1sum"
	jsonMapping[jsonKey] = out.String()

	// Dump everything into JSON in preperation for Elasticsearch upload
	jsonString, err := json.Marshal(jsonMapping)
	if err != nil {
		log.Println(err)
	}
	// Print raw json
	// fmt.Println(string(jsonString))

	//We want to return the JSON in addition to STDOUT
	*reply = string(jsonString)
	if err != nil {
		log.Println(err)
	}

	return nil
}
