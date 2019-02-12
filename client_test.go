package main

import (
	"os"
	"os/signal"
	"testing"
	"net/url"
	"github.com/gorilla/websocket"
	"log"
	"time"
	"encoding/base64"
	"net/http"
)

func TestClient(t *testing.T) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8996", Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte("admin:admin"))
	requestHeader := make(http.Header)
	requestHeader.Add("Authorization", auth)
	//c, _, err := dialer.Dial(u.String(), requestHeader)
	c, _, err := websocket.DefaultDialer.Dial(u.String(), requestHeader)

	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			s := `{"jsonrpc":"1.0","method":"uptime","params":[],"id":1}`
			err := c.WriteMessage(websocket.TextMessage, []byte(s))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
