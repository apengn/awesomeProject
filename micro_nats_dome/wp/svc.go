package main

import (
	"awesomeProject/micro_nats_dome/wp/debug"
	"context"
	"fmt"
	"time"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/nats"

	gonats "github.com/nats-io/go-nats"

	// nats "github.com/nats-io/nats.go"

	natsRegistry "github.com/micro/go-plugins/registry/nats"
	natsTransport "github.com/micro/go-plugins/transport/nats"
)

var (
	addrs1                     = []string{"nats://dev-7:4222", "nats://dev-8:4222", "nats://dev-9:4222"}
	NATS_STREAMING_CLUSTER_ID1 = "test-cluster"
	NATS_TOKEN1                = "NATS12345"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *debug.Request, rsp *debug.Response) error {
	rsp.Msg = "Hello 3" + req.Name
	fmt.Println(rsp.Msg)
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	natsServers := addrs1
	defaultOptions := gonats.GetDefaultOptions()

	defaultOptions.ClosedCB = func(conn *gonats.Conn) {
		fmt.Println("closed")
	}
	defaultOptions.Servers = natsServers
	defaultOptions.Token = NATS_TOKEN1
	defaultOptions.ReconnectedCB = func(conn *gonats.Conn) {
		fmt.Println("reconnected")
	}
	defaultOptions.DisconnectedCB = func(conn *gonats.Conn) {
		fmt.Println("disconnected")
	}
	defaultOptions.DiscoveredServersCB = func(conn *gonats.Conn) {
		fmt.Println("disccoveredservice")
	}

	registry := natsRegistry.NewRegistry(natsRegistry.Options(defaultOptions))
	transport := natsTransport.NewTransport(natsTransport.Options(defaultOptions))

	broker := nats.NewBroker(
		nats.Options(defaultOptions),
	)

	var err error
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(registry),
		micro.Broker(broker),
		micro.Transport(transport),
		micro.Context(ctx),
	)

	time.Sleep(2 * time.Second)
	debug.RegisterGreeterHandler(service.Server(), &Greeter{})
	err = service.Run()
	if err != nil {
		fmt.Println(err)
	}

}
