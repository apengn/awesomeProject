package main

import (
	"awesomeProject/micro_nats_streming_dome/wp/debug"
	"context"
	"encoding/json"
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
	addrs = []string{"nats://10.0.0.41:4222", "nats://10.0.0.42:4222", "nats://10.0.0.43:4222"}
	// addrs = []string{"nats://118.31.50.70:4222"}

	NATS_STREAMING_CLUSTER_ID = "test-cluster"
	// NATS_STREAMING_CLUSTER_ID = "test-cluster"
	NATS_TOKEN = "NATS12345"
)

func main() {
	// for {
	Send()
	// }
}

func TestClient() {
	for {
		Send()
	}
}

func Send() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	natsServers := addrs
	defaultOptions := nats.GetDefaultOptions()

	defaultOptions.ClosedCB = func(conn *nats.Conn) {
		fmt.Println("closed")
	}
	defaultOptions.Servers = natsServers
	defaultOptions.Token = NATS_TOKEN
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
	stanOptions.NatsURL = strings.Join(addrs, ",")
	var err error
	stanOptions.NatsConn, err = gonats.Connect(stanOptions.NatsURL, gonats.Token(NATS_TOKEN))
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
		stanBroker.ClusterID(NATS_STREAMING_CLUSTER_ID),
	)

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(registry),
		micro.Broker(broker),
		micro.Transport(transport),
		micro.Context(ctx),
	)

	c := service.Client()
	time.Sleep(1 * time.Second)

	response, err := debug.NewGreeterService("greeter", c).Hello(context.TODO(), &debug.Request{Name: "pppppp"})

	if err != nil {
		return
	}
	b, _ := json.Marshal(response)
	fmt.Println(string(b))
}
