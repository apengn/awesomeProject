package main

import (
	"awesomeProject/micro_nats_dome/wp/debug"
	"context"
	"encoding/json"
	"fmt"
	"time"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/nats"

	nats2 "github.com/nats-io/go-nats"

	natsRegistry "github.com/micro/go-plugins/registry/nats"
	natsTransport "github.com/micro/go-plugins/transport/nats"
)

var (
	addrs                     = []string{"nats://dev-7:4222", "nats://dev-8:4222", "nats://dev-9:4222"}
	NATS_STREAMING_CLUSTER_ID = "test-cluster"
	NATS_TOKEN                = "NATS12345"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	natsServers := addrs
	defaultOptions := nats2.GetDefaultOptions()

	defaultOptions.ClosedCB = func(conn *nats2.Conn) {
		fmt.Println("closed")
	}
	defaultOptions.Servers = natsServers
	defaultOptions.Token = NATS_TOKEN

	defaultOptions.ReconnectedCB = func(conn *nats2.Conn) {
		fmt.Println("reconnected")
	}
	defaultOptions.DisconnectedCB = func(conn *nats2.Conn) {
		fmt.Println("disconnected")
	}
	defaultOptions.DiscoveredServersCB = func(conn *nats2.Conn) {
		fmt.Println("disccoveredservice")
	}

	// var err error
	registry := natsRegistry.NewRegistry(natsRegistry.Options(defaultOptions))
	transport := natsTransport.NewTransport(natsTransport.Options(defaultOptions))
	broker := nats.NewBroker(nats.Options(defaultOptions))
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(registry),
		micro.Broker(broker),
		micro.Transport(transport),
		micro.Context(ctx),
	)

	c := service.Client()

	time.Sleep(1 * time.Second)

	response, err := debug.NewGreeterService("greeter", c).Hello(context.TODO(), &debug.Request{Name: "xxxxxx"})

	if err != nil {
		return
	}
	b, _ := json.Marshal(response)
	fmt.Println(string(b))
}
