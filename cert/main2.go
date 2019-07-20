package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var pemPath string
	flag.StringVar(&pemPath, "pem", "nginx.crt", "path to pem file")
	var keyPath string
	flag.StringVar(&keyPath, "key", "nginx.key", "path to key file")
	var proto string
	flag.StringVar(&proto, "proto", "https", "Proxy protocol (http or https)")
	flag.Parse()
	if proto != "http" && proto != "https" {
		log.Fatal("Protocol must be either http or https")
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "post==========")
		fmt.Println(request.Host)
		fmt.Println("=====" + request.URL.Scheme)
	})

	http.ListenAndServeTLS(":999", pemPath, keyPath, nil)
}
