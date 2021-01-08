package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

var wg sync.WaitGroup

type Pagador struct {
	Name string `json:"name"`
}

func main() {

	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	nc.Subscribe("pagado.create", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
	})

	nc.Subscribe("pagado.create.request", func(msg *nats.Msg) {
		p := Pagador{}
		err := json.Unmarshal(msg.Data, &p)
		if err != nil {
			panic(err)
		}
		fmt.Println("Chergou aqui " + p.Name)
		info := fmt.Sprintf("%v %v", p.Name, "seja bem vindo")
		// adicionar no banco
		nc.Publish(msg.Reply, []byte(info))
	})

	wg.Wait()
}
