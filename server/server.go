package main

import (
	"fmt"
	"log"
	"net/http"
)

var serverlist string

// execution start
func main() {
	_init()
}

func _init() {
	// handlers
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/list", handleList)

	// webserver init
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	addr := r.Header["Addr"]
	fmt.Fprintf(w, "Registering, %v", addr)
	updateServerlist(addr)
}

func handleList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", serverlist)
}

func updateServerlist(addr string) {
	serverlist = serverlist + addr
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
