package main

import (
	"log"
	"net/http"
	"os"

	"github.com/husaria-dev/party/server/connection"
	"github.com/husaria-dev/party/server/inmem"
)

var (
	port         = os.Getenv("PORT")
	roomIdLength = 10
)

func main() {
	rs := inmem.NewRoomService()
	connHandler := connection.NewHandler(rs, roomIdLength)

	if err := http.ListenAndServe(":"+port, connHandler); err != nil {
		log.Fatal(err)
	}
}
