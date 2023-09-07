// * Package `utils` implements a collection of utility functions for use in other packages.
package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

// * flag for specifying the address to connect to
var host = flag.String("host", "localhost", "the host to connect to")
var port = flag.String("port", "50051", "the port to connect to")

// * GetAddr returns the address to connect to
func GetAddr() string {
	flag.Parse()
	addr := *host + ":" + *port
	return addr
}

// * GetPort returns the port to connect to
func GetPort() string {
	flag.Parse()
	return *port
}

// * transform any struct/interface to json
func MakeJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
	}
	return string(b)
}

// * wrapper for logging errors without exit / termination
func Log(err error, msg_optional ...string) {
	if err != nil {
		msg := fmt.Sprintf("ERR: %v", err)
		if len(msg_optional) > 0 {
			msg += "\n - " + msg_optional[0]
		}
		log.Println(msg)
	}
}

// * wrapper for logging errors with exit / termination
func LogF(err error, msg_optional ...string) {
	if err != nil {
		msg := fmt.Sprintf("FTL: %v", err)
		if len(msg_optional) > 0 {
			msg += "\n - " + msg_optional[0]
		}
		log.Fatal(msg)
	}
}

// * wrapper for logging a struct value to stdout (includes timestamp)
func LogJSON(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		Log(err)
	}
	log.Println(string(b))
}

// * wrapper for printing a struct to stout (does not include timestamp)
func Pjson(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		Log(err)
	}
	fmt.Println(string(b))
}

// * wrapper for printing a payload to stout (does not include timestamp)
func LogPayload(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		Log(err)
	}
	log.Printf("Payload: %s", string(b))
}
