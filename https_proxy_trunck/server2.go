package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var pemPath string
	flag.StringVar(&pemPath, "pem", "https_proxy_trunck/registry.wise2c.com.crt", "path to pem file")
	var keyPath string
	flag.StringVar(&keyPath, "key", "https_proxy_trunck/registry.wise2c.com.key", "path to key file")
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

	http.ListenAndServe(":65002", nil)
	// http.ListenAndServeTLS(":65002", pemPath, keyPath, nil)
}
