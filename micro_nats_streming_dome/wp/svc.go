package main

import (
	"awesomeProject/micro_nats_streming_dome/wp/debug"
	"context"
	"fmt"
	"strings"
	"time"

	nats "github.com/nats-io/go-nats"
	gonats "github.com/nats-io/nats.go"

	"github.com/astaxie/beego/logs"
	micro "github.com/micro/go-micro"
	stanBroker "github.com/micro/go-plugins/broker/stan"
	natsRegistry "github.com/micro/go-plugins/registry/nats"
	natsTransport "github.com/micro/go-plugins/transport/nats"
	stan "github.com/nats-io/go-nats-streaming"
)

// NATS_STREAMING_CLUSTER_ID: test-cluster
// NATS_STREAMING_CONNECTION_TIMEOUT: "2"
// NATS_STREAMING_MAX_RECONNECT: "60"
// NATS_STREAMING_RECONNECT_WAIT: "4"
// NATS_STREAMING_URLS: nats://192.168.1.191:4222,nats://192.168.1.188:4222,nats://192.168.1.190:4222
// NATS_URL: nats://192.168.1.191:4222,nats://192.168.1.188:4222,nats://192.168.1.190:4222
var (
	addrs1 = []string{"nats://10.0.0.41:4222", "nats://10.0.0.42:4222", "nats://10.0.0.43:4222"}
	// addrs = []string{"nats://118.31.50.70:4222"}

	NATS_STREAMING_CLUSTER_ID1 = "test-cluster"
	// NATS_STREAMING_CLUSTER_ID = "test-cluster"
	NATS_TOKEN1 = "NATS12345"
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
	defaultOptions := nats.GetDefaultOptions()

	defaultOptions.ClosedCB = func(conn *nats.Conn) {
		fmt.Println("closed")
	}
	defaultOptions.Servers = natsServers
	defaultOptions.Token = NATS_TOKEN1
	defaultOptions.ReconnectedCB = func(conn *nats.Conn) {
		fmt.Println("reconnected")
	}
	defaultOptions.DisconnectedCB = func(conn *nats.Conn) {
		fmt.Println("disconnected")
	}
	defaultOptions.DiscoveredServersCB = func(conn *nats.Conn) {
		fmt.Println("disccoveredservice")
	}

	registry := natsRegistry.NewRegistry(natsRegistry.Options(defaultOptions))
	transport := natsTransport.NewTransport(natsTransport.Options(defaultOptions))

	stanOptions := stan.DefaultOptions
	stanOptions.NatsURL = strings.Join(addrs1, ",")
	var err error
	stanOptions.NatsConn, err = gonats.Connect(stanOptions.NatsURL, gonats.Token(NATS_TOKEN1))
	if err != nil {
		logs.Info("nats.Connect err: %q", err)
	}
	stanOptions.ConnectionLostCB = func(conn stan.Conn, e error) {
		defer conn.Close()
		if e != nil {
			logs.Error("go-stan close! Reason: %q", e)
		}
		logs.Info("go-stan close!")
	}

	broker := stanBroker.NewBroker(
		stanBroker.Options(stanOptions),
		stanBroker.ClusterID(NATS_STREAMING_CLUSTER_ID1),
	)

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(registry),
		micro.Broker(broker),
		micro.Transport(transport),
		micro.Context(ctx),
	)
	// service.Init()
	s := service.Server()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			err = debug.RegisterGreeterHandler(s, &Greeter{})
		}

	}()
	// micro.RegisterSubscriber("greeter",s,&Greeter{})
	err = service.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func TestSvc() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	natsServers := addrs1
	// options := nats.GetDefaultOptions()
	// options.Servers = natsServers
	// options.Token = NATS_TOKEN
	// options.ReconnectedCB = func(_ *nats.Conn) {
	// 	logs.Info("go nats reconnect!")
	// }
	// options.ClosedCB = func(nc *nats.Conn) {
	// 	if err := nc.LastError(); err != nil {
	// 		logs.Error("go nats close! err:", err)
	// 		return
	// 	}
	// 	logs.Info("go nats close!")
	// }
	// options.DisconnectedCB = func(_ *nats.Conn) {
	// 	logs.Info("go nats disconnect!")
	// }

	defaultOptions := nats.GetDefaultOptions()

	defaultOptions.ClosedCB = func(conn *nats.Conn) {
		fmt.Println("closed")
	}
	defaultOptions.Servers = natsServers
	defaultOptions.Token = NATS_TOKEN1
	defaultOptions.ReconnectedCB = func(conn *nats.Conn) {
		fmt.Println("reconnected")
	}
	defaultOptions.DisconnectedCB = func(conn *nats.Conn) {
		fmt.Println("disconnected")
	}
	defaultOptions.DiscoveredServersCB = func(conn *nats.Conn) {
		fmt.Println("disccoveredservice")
	}

	registry := natsRegistry.NewRegistry(natsRegistry.Options(defaultOptions))
	transport := natsTransport.NewTransport(natsTransport.Options(defaultOptions))

	stanOptions := stan.DefaultOptions
	stanOptions.NatsURL = strings.Join(addrs1, ",")
	var err error
	stanOptions.NatsConn, err = gonats.Connect(stanOptions.NatsURL, gonats.Token(NATS_TOKEN1))
	if err != nil {
		logs.Info("nats.Connect err: %q", err)
	}
	stanOptions.ConnectionLostCB = func(conn stan.Conn, e error) {
		defer conn.Close()
		if e != nil {
			logs.Error("go-stan close! Reason: %q", e)
		}
		logs.Info("go-stan close!")
	}

	broker := stanBroker.NewBroker(
		stanBroker.Options(stanOptions),
		stanBroker.ClusterID(NATS_STREAMING_CLUSTER_ID1),
	)

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(registry),
		micro.Broker(broker),
		micro.Transport(transport),
		micro.Context(ctx),
	)

	s := service.Server()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			err = debug.RegisterGreeterHandler(s, &Greeter{})
		}

	}()

	err = service.Run()
	if err != nil {
		fmt.Println(err)
	}

}
