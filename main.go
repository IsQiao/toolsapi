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

	log.Fatal(http.ListenAndServe(":8080", nil))
}
