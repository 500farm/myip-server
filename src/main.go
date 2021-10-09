package main

import (
	"fmt"
	"net"
	"net/http"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	listenAddress = kingpin.Flag(
		"listen",
		"Address to listen on.",
	).Default(":80").String()
)

func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		w.Write([]byte(fmt.Sprintf(`{"ip":"%s"}`, host)))
	})

	fmt.Println("MyIP server listening on", *listenAddress)
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		fmt.Println(err)
	}
}
