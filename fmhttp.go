package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type testMessage struct {
	Date    string `json:"date"`
	IP      string `json:"ip"`
	Release string `json:"rel"`
	Count   int    `json:"cnt"`
}

var cnt = 0

func displayTest(w http.ResponseWriter, r *http.Request) {
	data := testMessage{
		time.Now().Format(time.RFC850),
		getOutboundIP(),
		"fm1.0",
		cnt,
	}
	cnt = cnt + 1
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(b))
}

// Get preferred outbound ip of this machine
func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx]
}

func main() {
	http.HandleFunc("/", displayTest)
	http.ListenAndServe(":5555", nil)
}
