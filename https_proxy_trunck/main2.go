package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

var okHeader = []byte("HTTP/1.1 200 OK\r\n\r\n")

func handleTunneling(w http.ResponseWriter, r *http.Request) {
	dest_conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	client_conn, _, err := hijacker.Hijack()

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	// client_conn.Write(okHeader)
	go transfer(dest_conn, client_conn)
	go transfer(client_conn, dest_conn)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

type Proxy struct {
	secondHandle http.Handler
}

func (proxy *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect {
		handleTunneling(w, r)
	} else if r.URL.Scheme == "" {
		proxy.secondHandle.ServeHTTP(w, r)
	} else {
		handleHTTP(w, r)
	}
}

func main() {
	var pemPath string
	flag.StringVar(&pemPath, "pem", "https_proxy_trunck/registry.wise2c.com.crt", "path to pem file")
	var keyPath string
	flag.StringVar(&keyPath, "key", "https_proxy_trunck/registry.wise2c.com.key", "path to key file")
	var proto string
	flag.StringVar(&proto, "proto", "http", "Proxy protocol (http or https)")
	flag.Parse()
	if proto != "http" && proto != "https" {
		log.Fatal("Protocol must be either http or https")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/222/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "xxxxxxxx")
	}))
	mux.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ffffffff")
	}))
	server := http.Server{Addr: ":65001"}
	proxy := &Proxy{}

	proxy.secondHandle = mux

	server.Handler = proxy
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
