package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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
		if request.Method == http.MethodGet {
			fmt.Println(request.Header)
			fmt.Fprint(writer, "response writer: method"+request.Method)
			return
		}
		backendURL, err := url.Parse("http://registry.wise2c.com:65002/")
		if err != nil {
			fmt.Println(err)
		}
		reverseProxy := httputil.NewSingleHostReverseProxy(backendURL)
		reverseProxy.ServeHTTP(writer, request)
	})
	http.ListenAndServeTLS(":65001", pemPath, keyPath, nil)
}
