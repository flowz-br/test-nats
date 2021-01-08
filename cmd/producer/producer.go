package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		panic(err)
	}

	// nc.Publish("pagado.create", []byte("{asdasdasdasd}"))

	resp, err := nc.Request("pagado.create.request", []byte("{\"name\":\"romani\"}"), 3000*time.Millisecond)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp.Data))

}
