package main

import (
	"log"
	"net/http"
	"os"

	"github.com/husaria-dev/party/server"
	"github.com/husaria-dev/party/server/handler"
	"github.com/husaria-dev/party/server/inmem"
)

var (
	port         = os.Getenv("PORT")
	roomIdLength = 10
)

func main() {
	rs := inmem.NewRoomService()

	connPool := server.NewConnectionPool()
	connHandler := handler.New(rs, connPool)

	mux := http.NewServeMux()

	mux.Handle("/api/", http.StripPrefix("/api", connHandler))
	mux.Handle("/", http.FileServer(http.Dir("./testui")))

	addr := ":8080"
	if port != "" {
		addr = ":" + port
	}

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
