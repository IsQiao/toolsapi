package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

func getIP(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		http.Error(w, fmt.Sprintf("userip: %q is not IP:port", r.RemoteAddr), 500)
		return
	}

	res := GetIP{
		ip: ip,
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	})
	http.HandleFunc("/ip", getIP)
	http.HandleFunc("/ip2", getIP2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getIP2(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "<h1>static file server</h1><p><a href='./static'>folder</p></a>")

	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)

		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
		return
	}

	// This will only be defined when site is accessed via non-anonymous proxy
	// and takes precedence over RemoteAddr
	// Header.Get is case-insensitive
	forward := req.Header.Get("X-Forwarded-For")

	fmt.Fprintf(w, "<p>IP: %s</p>", ip)
	fmt.Fprintf(w, "<p>Port: %s</p>", port)
	fmt.Fprintf(w, "<p>Forwarded for: %s</p>", forward)
}
