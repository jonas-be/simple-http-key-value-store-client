package main

import (
	"flag"
	"fmt"
	"simple-http-key-value-store-client/internal/client"
)

func main() {
	methodFlag := flag.String("m", "", "method to execute")
	keyFlag := flag.String("key", "", "key to access")
	valueFlag := flag.String("value", "", "value to set")

	flag.Parse()

	client := client.Client{BaseURL: "http://localhost:8080"}
	fmt.Println(client.Request(*methodFlag, *keyFlag, *valueFlag))
}
