package main

import (
	"fmt"
	"log"

	"github.com/goswagger-2471/pkg/example/client"
)

func main() {
	resp, err := client.Default.Operations.All(operations.AllParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}
