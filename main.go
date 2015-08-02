package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func usage() {
	fmt.Printf("Usage: %s [laddr]\n", os.Args[0])
	fmt.Printf("          ")
	for _ = range os.Args[0] {
		fmt.Printf(" ")
	}
	fmt.Println("^-- default to \":8080\"")
}

func main() {
	laddr := ":8080"
	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			usage()
			os.Exit(0)
		}
		laddr = os.Args[1]
	}

	var fsh http.Handler
	if pwd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		fsh = http.FileServer(http.Dir(pwd))
	}

	acceptAllDeadline := time.Now()
	mu := new(sync.Mutex)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			fmt.Printf("Non \"GET\" request received; responding with StatusForbidden\n")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if addr, er := net.ResolveTCPAddr("tcp", r.RemoteAddr); (er == nil && addr.IP.IsLoopback()) || acceptAllDeadline.After(time.Now()) {
			fsh.ServeHTTP(w, r)
			fmt.Printf("\"GET\" Request (%s) from %s has been served.\n", r.URL, r.RemoteAddr)
			return
		}

		mu.Lock()
		defer mu.Unlock()
		fmt.Printf("\"GET\" Request (%s) from %s - %#v\n", r.URL, r.RemoteAddr, r.UserAgent())
		fmt.Println("Options:")
		fmt.Println("  [y] accept;")
		fmt.Println("  [n] reject;")
		fmt.Println("  [d] details;")
		fmt.Println("  [2] accept all requests in following 2 seconds")

	poll:
		for {
			var s string
			fmt.Printf("* Accept? ")
			fmt.Scanf("%s", &s)
			switch strings.TrimSpace(s) {
			case "y":
				fsh.ServeHTTP(w, r)
				fmt.Println("The request has been served.")
				break poll
			case "n":
				w.WriteHeader(http.StatusForbidden)
				fmt.Println("The request has been denied with StatusForbidden.")
				break poll
			case "d":
				fmt.Printf("%#v\n", r)
			case "2":
				acceptAllDeadline = time.Now().Add(2 * time.Second)
				fmt.Println("The request has been served. All requests within 2 seconds will be served automatically.")
				fsh.ServeHTTP(w, r)
				break poll
			default:
			}
		}
	})

	fmt.Printf("Listening at %s ...\n", laddr)
	if err := http.ListenAndServe(laddr, nil); err != nil {
		usage()
		panic(err)
	}
}
